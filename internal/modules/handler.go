package modules

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

// List All Modulus

func (h *Handler) GetModules(c *gin.Context) {

	resp, err := h.svc.GetModules(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to fetch modules list", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch modules",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
