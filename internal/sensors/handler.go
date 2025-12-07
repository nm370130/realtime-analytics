package sensors

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

// TYPE BREAKDOWN API

func (h *Handler) GetTypeBreakdown(c *gin.Context) {

	resp, err := h.svc.GetTypeBreakdown(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to fetch sensor type breakdown", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch sensor type breakdown",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
