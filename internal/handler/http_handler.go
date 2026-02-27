package handler

import (
	"net/http"
	"strconv"

	"api-pokemon-meta-go/internal/domain"
	"api-pokemon-meta-go/internal/service"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	service *service.PokemonService
}

func NewPokemonHandler(service *service.PokemonService) *PokemonHandler {
	return &PokemonHandler{service: service}
}

// GetMetaSnapshot godoc
// @Summary Obtener Snapshot del Meta
// @Tags Meta
// @Produce json
// @Param tier query string false "Tier" default(ou)
// @Param gen query string false "Generation" default(gen9)
// @Param limit query int false "Limit" default(10)
// @Success 200 {array} domain.MetaSnapshotItem
// @Failure 500 {object} domain.ErrorResponse
// @Router /meta/snapshot [get]
func (h *PokemonHandler) GetMetaSnapshot(c *gin.Context) {
	tier := c.DefaultQuery("tier", "ou")
	gen := c.DefaultQuery("gen", "gen9")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	result, err := h.service.GetMetaSnapshot(tier, gen, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if result == nil {
		result = []domain.MetaSnapshotItem{}
	}

	c.JSON(http.StatusOK, result)
}

// GetSpeedCreepers godoc
// @Summary Obtener Anti-Meta Speed Creepers
// @Tags Anti-Meta
// @Produce json
// @Param tier query string false "Tier" default(ou)
// @Param gen query string false "Generation" default(gen9)
// @Success 200 {array} domain.SpeedCreeperItem
// @Failure 500 {object} domain.ErrorResponse
// @Router /anti-meta/speed-creepers [get]
func (h *PokemonHandler) GetSpeedCreepers(c *gin.Context) {
	tier := c.DefaultQuery("tier", "ou")
	gen := c.DefaultQuery("gen", "gen9")

	result, err := h.service.GetSpeedCreepers(tier, gen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if result == nil {
		result = []domain.SpeedCreeperItem{}
	}

	c.JSON(http.StatusOK, result)
}

// GetWallbreakers godoc
// @Summary Obtener Anti-Meta Wallbreakers
// @Tags Anti-Meta
// @Produce json
// @Param tier query string false "Tier" default(ou)
// @Param gen query string false "Generation" default(gen9)
// @Param move_type query string false "Move Type" default(ice)
// @Success 200 {array} domain.WallbreakerItem
// @Failure 500 {object} domain.ErrorResponse
// @Router /anti-meta/wallbreakers [get]
func (h *PokemonHandler) GetWallbreakers(c *gin.Context) {
	tier := c.DefaultQuery("tier", "ou")
	gen := c.DefaultQuery("gen", "gen9")
	moveType := c.DefaultQuery("move_type", "ice")

	result, err := h.service.GetWallbreakers(tier, gen, moveType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if result == nil {
		result = []domain.WallbreakerItem{}
	}

	c.JSON(http.StatusOK, result)
}

// GetPokemonTrend godoc
// @Summary Obtener Tendencia de Uso
// @Tags Analytics
// @Produce json
// @Param name query string true "Pokemon Name"
// @Param tier query string false "Tier" default(ou)
// @Param gen query string false "Generation" default(gen9)
// @Success 200 {object} domain.PokemonTrendResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /analytics/trend [get]
func (h *PokemonHandler) GetPokemonTrend(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "El nombre del Pokémon es requerido"})
		return
	}
	tier := c.DefaultQuery("tier", "ou")
	gen := c.DefaultQuery("gen", "gen9")

	result, err := h.service.GetPokemonTrend(name, tier, gen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	if result != nil && result.History == nil {
		result.History = []domain.TrendPoint{}
	}

	c.JSON(http.StatusOK, result)
}
