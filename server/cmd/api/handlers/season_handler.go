package handlers

import (
	"net/http"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeasonHandler struct {
	service *services.SeasonService
}

func NewSeasonHandler(service *services.SeasonService) *SeasonHandler {
	return &SeasonHandler{service: service}
}

func (h *SeasonHandler) CreateSeason(c *gin.Context) {
	type CreateSeasonRequest struct {
		Year     string `json:"year" binding:"required"`
		InitDate string `json:"init" binding:"required"`
		EndDate  string `json:"end" binding:"required"`
	}

	var req CreateSeasonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leagueSlug := c.Param("leagueSlug")
	season, err := h.service.CreateSeason(&models.Season{
		Year:     req.Year,
		InitDate: req.InitDate,
		EndDate:  req.EndDate,
	}, leagueSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": season})
}

func (h *SeasonHandler) GetSeasonsByLeague(c *gin.Context) {
	leagueSlug := c.Param("leagueSlug")
	seasons, err := h.service.GetSeasonsByLeagueSlug(leagueSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": seasons})
}

func (h *SeasonHandler) GetSeasonById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID format"})
		return
	}

	season, err := h.service.GetSeasonById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Season not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": season})
}
