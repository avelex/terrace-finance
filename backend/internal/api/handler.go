package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avelex/terrace-finance/backend/internal/models/enum"
	"github.com/avelex/terrace-finance/backend/internal/models/request"
	"github.com/avelex/terrace-finance/backend/internal/models/response"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type StrategyService interface {
	SendAllFunds(ctx context.Context, srcDomain, dstDomain enum.CircleDomain) (string, error)
	SupplyAllFundsToAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
	WithdrawAllFundsFromAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
}

type WalletService interface {
	ProtocolBalances(ctx context.Context, user common.Address) (*response.ProtocolBalances, error)
	UnifyUSDC(ctx context.Context, user common.Address, domains []enum.CircleDomain) (*response.PermitPayload, error)
	SaveUnifyPermitsSignatures(ctx context.Context, user common.Address, signedPermits request.SignedPermits) error
	SaveDepositAttestationAndSignature(ctx context.Context, user common.Address, deposit request.DepositAndStake) error
}

type Handler struct {
	strategyService StrategyService
	walletService   WalletService
}

func NewHandler(ss StrategyService, ws WalletService) *Handler {
	return &Handler{
		strategyService: ss,
		walletService:   ws,
	}
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/send/:src/:dst", h.send)

	strategyGroup := g.Group("/strategy")
	{
		strategyGroup.GET("/supply/:domain", h.supply)
		strategyGroup.GET("/withdraw/:domain", h.withdraw)
	}

	walletGroup := g.Group("/wallet")
	{
		walletGroup.GET("/:address/balances", h.walletBalances)
		walletGroup.POST("/:address/unify", h.walletUnify)
		walletGroup.POST("/:address/unify/permits", h.walletUnifyPermits)
		walletGroup.POST("/:address/deposit_stake", h.walletDepositAndStake)
	}
}

func (h *Handler) send(c echo.Context) error {
	srcDomain, err := parseDomain(c.Param("src"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid src domain")
	}

	dstDomain, err := parseDomain(c.Param("dst"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid dst domain")
	}

	txHash, err := h.strategyService.SendAllFunds(c.Request().Context(), srcDomain, dstDomain)
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
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	var req request.UnifyBalances
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	permits, err := h.walletService.UnifyUSDC(c.Request().Context(), address, req.Domains)
	if err != nil {
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

	if err := h.walletService.SaveDepositAttestationAndSignature(c.Request().Context(), address, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
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
