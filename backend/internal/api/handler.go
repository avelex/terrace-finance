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
)

type StrategyManager interface {
	SendAllFunds(ctx context.Context, srcDomain, dstDomain enum.CircleDomain) (string, error)
	SupplyAllFundsToAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
	WithdrawAllFundsFromAaveV3(ctx context.Context, domain enum.CircleDomain) (string, error)
}

type WalletManager interface {
	ProtocolBalances(ctx context.Context, address common.Address) (*response.ProtocolBalances, error)
	UnifyUSDC(ctx context.Context, address common.Address, networks []uint32) (*response.PermitPayload, error)
	SaveUnifyPermits(ctx context.Context, address common.Address, networks map[uint32]string) error
	SaveBurnIntents(ctx context.Context, address common.Address, intents []models.BurnIntent) error
}

type Handler struct {
	strategyManager StrategyManager
	walletManager   WalletManager
}

func NewHandler(sm StrategyManager, wm WalletManager) *Handler {
	return &Handler{
		strategyManager: sm,
		walletManager:   wm,
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

	txHash, err := h.strategyManager.SendAllFunds(c.Request().Context(), srcDomain, dstDomain)
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

	txHash, err := h.strategyManager.SupplyAllFundsToAaveV3(c.Request().Context(), domain)
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

	txHash, err := h.strategyManager.WithdrawAllFundsFromAaveV3(c.Request().Context(), domain)
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

	balances, err := h.walletManager.ProtocolBalances(c.Request().Context(), address)
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

	permits, err := h.walletManager.UnifyUSDC(c.Request().Context(), address, req.Networks)
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

	var req request.SignedPermit
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := h.walletManager.SaveUnifyPermits(c.Request().Context(), address, req.Networks); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *Handler) walletDepositAndStake(c echo.Context) error {
	address, err := parseAddress(c.Param("address"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid address")
	}

	var req request.BurnIntents
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := h.walletManager.SaveBurnIntents(c.Request().Context(), address, req.Intents); err != nil {
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
