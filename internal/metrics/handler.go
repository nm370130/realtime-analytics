package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	svc    Service
	logger *zap.Logger
}

func NewHandler(svc Service, logger *zap.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

// SUMMARY API

func (h *Handler) GetSummary(c *gin.Context) {
	resp, err := h.svc.GetSummary(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to fetch summary metrics", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch metrics summary",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// HISTORY API

func (h *Handler) GetHistory(c *gin.Context) {
	metric := c.Query("metric")
	interval := c.Query("interval")

	// Required params
	if metric == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query param 'metric' is required",
		})
		return
	}

	if interval == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query param 'interval' is required",
		})
		return
	}

	resp, err := h.svc.GetHistory(c.Request.Context(), metric, interval)
	if err != nil {
		h.logger.Error("failed to fetch metric history", zap.String("metric", metric), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch metric history",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
