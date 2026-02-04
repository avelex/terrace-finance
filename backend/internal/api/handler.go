package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avelex/terrace-finance/backend/internal/models/enum"

	"github.com/labstack/echo/v4"
)

type StrategyManager interface {
	SendAllFunds(ctx context.Context, srcDomain, dstDomain uint32) (string, error)
	SupplyAllFundsToAaveV3(ctx context.Context, domain uint32) (string, error)
	WithdrawAllFundsFromAaveV3(ctx context.Context, domain uint32) (string, error)
}

type Handler struct {
	manager StrategyManager
}

func NewHandler(manager StrategyManager) *Handler {
	return &Handler{
		manager: manager,
	}
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/send/:src/:dst", h.send)
	g.GET("/strategy/supply/:domain", h.supply)
	g.GET("/strategy/withdraw/:domain", h.withdraw)
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

	txHash, err := h.manager.SendAllFunds(c.Request().Context(), srcDomain, dstDomain)
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

	txHash, err := h.manager.SupplyAllFundsToAaveV3(c.Request().Context(), domain)
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

	txHash, err := h.manager.WithdrawAllFundsFromAaveV3(c.Request().Context(), domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"txHash": txHash,
	})
}

func parseDomain(domain string) (uint32, error) {
	d, err := strconv.ParseUint(domain, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid domain: %w", err)
	}

	if _, ok := enum.DomainNetworkMapping[uint32(d)]; !ok {
		return 0, fmt.Errorf("unknown domain: %d", d)
	}

	return uint32(d), nil
}
