package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avelex/terrace-finance/backend/internal/models"
	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/models/request"
	"github.com/avelex/terrace-finance/backend/internal/models/response"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type StrategyService interface {
	SupplyAllFundsToAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
	WithdrawAllFundsFromAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
}

type BridgeService interface {
	BridgeFunds(ctx context.Context, srcDomain, dstDomain enum.CircleDomain) (string, error)
	GetBridgeOperations(ctx context.Context, limit, page int) (response.PagableBridgeOps, error)
	GetVaultsInfo(ctx context.Context) ([]response.VaultInfo, error)
}

type WalletService interface {
	ProtocolBalances(ctx context.Context, user common.Address) (*response.ProtocolBalances, error)
	UserUnifyDeposits(ctx context.Context, user common.Address) ([]models.UserUnifiedPermits, error)
	UnifyUSDC(ctx context.Context, user common.Address, domains []enum.CircleDomain) (*response.PermitPayload, error)
	SaveUnifyPermitsSignatures(ctx context.Context, user common.Address, signedPermits request.SignedPermits) error
	InitDepositAndStake(ctx context.Context, user common.Address, deposit request.DepositAndStake) error
	UserDeposits(ctx context.Context, user common.Address) ([]models.UserDeposit, error)
}

type Handler struct {
	strategyService StrategyService
	walletService   WalletService
	bridgeService   BridgeService
}

func NewHandler(ss StrategyService, ws WalletService, bs BridgeService) *Handler {
	return &Handler{
		strategyService: ss,
		walletService:   ws,
		bridgeService:   bs,
	}
}

func (h *Handler) Register(g *echo.Group) {
	bridgeGroup := g.Group("/bridge")
	{
		bridgeGroup.GET("/send/:src/:dst", h.bridgeSend)
		bridgeGroup.GET("/operations", h.showBridgeOperations)
		bridgeGroup.GET("/vaults", h.showVaultsInfo)
	}

	strategyGroup := g.Group("/strategy")
	{
		strategyGroup.GET("/supply/:domain", h.supply)
		strategyGroup.GET("/withdraw/:domain", h.withdraw)
	}

	walletGroup := g.Group("/wallet")
	{
		walletGroup.GET("/:address/balances", h.walletBalances)
		walletGroup.POST("/:address/unify", h.walletUnify)
		walletGroup.GET("/:address/unify/list", h.walletUnifyDeposits)
		walletGroup.POST("/:address/unify/permits", h.walletUnifyPermits)
		walletGroup.POST("/:address/deposit_stake", h.walletDepositAndStake)
		walletGroup.GET("/:address/deposit_stake/list", h.walletDeposits)
	}

}

func (h *Handler) bridgeSend(c echo.Context) error {
	srcDomain, err := parseDomain(c.Param("src"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid src domain")
	}

	dstDomain, err := parseDomain(c.Param("dst"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid dst domain")
	}

	txHash, err := h.bridgeService.BridgeFunds(c.Request().Context(), srcDomain, dstDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"txHash": txHash,
	})
}

func (h *Handler) supply(c echo.Context) error {
	domain, err := parseDomain(c.Param("domain"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid domain")
	}

	txHash, err := h.strategyService.SupplyAllFundsToAaveV3(c.Request().Context(), domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"txHash": txHash,
	})
}

func (h *Handler) withdraw(c echo.Context) error {
	domain, err := parseDomain(c.Param("domain"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid domain")
	}

	txHash, err := h.strategyService.WithdrawAllFundsFromAaveV3(c.Request().Context(), domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"txHash": txHash,
	})
}

func (h *Handler) walletBalances(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	balances, err := h.walletService.ProtocolBalances(c.Request().Context(), address)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, balances)
}

func (h *Handler) walletUnify(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		log.Error().Err(err).Msg("invalid address")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	var req request.UnifyBalances
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("invalid request")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	permits, err := h.walletService.UnifyUSDC(c.Request().Context(), address, req.Domains)
	if err != nil {
		log.Error().Err(err).Msg("failed to unify USDC")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, permits)
}

func (h *Handler) walletUnifyPermits(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	var req request.SignedPermits
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := h.walletService.SaveUnifyPermitsSignatures(c.Request().Context(), address, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *Handler) walletDepositAndStake(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	var req request.DepositAndStake
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := h.walletService.InitDepositAndStake(c.Request().Context(), address, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *Handler) walletUnifyDeposits(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	permits, err := h.walletService.UserUnifyDeposits(c.Request().Context(), address)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, permits)
}

func (h *Handler) walletDeposits(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	deposits, err := h.walletService.UserDeposits(c.Request().Context(), address)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, deposits)
}

func (h *Handler) showBridgeOperations(c echo.Context) error {
	limit := 10
	if l, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		limit = l
	}

	page := 0
	if p, err := strconv.Atoi(c.QueryParam("page")); err == nil {
		page = p
	}

	ops, err := h.bridgeService.GetBridgeOperations(c.Request().Context(), limit, page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ops)
}

func (h *Handler) showVaultsInfo(c echo.Context) error {
	vaultsInfo, err := h.bridgeService.GetVaultsInfo(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, vaultsInfo)
}

func parseAddress(addr string) (common.Address, error) {
	if !common.IsHexAddress(addr) {
		return common.Address{}, fmt.Errorf("invalid address: %s", addr)
	}
	return common.HexToAddress(addr), nil
}

func parseDomain(domain string) (enum.CircleDomain, error) {
	d, err := strconv.ParseUint(domain, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid domain: %w", err)
	}

	cd := enum.CircleDomain(d)

	if _, ok := enum.DomainNetworkMapping[cd]; !ok {
		return 0, fmt.Errorf("unknown domain: %d", d)
	}

	return cd, nil
}
