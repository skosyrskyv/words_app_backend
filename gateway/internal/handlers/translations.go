package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "words-app.local/protos/gen/translations"
)

type TranslationsHandler struct {
	client pb.TranslationsClient
	logger *slog.Logger
}

func NewTranslationsHandler(addr string, logger *slog.Logger) *TranslationsHandler {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect to translations service", slog.String("addr", addr), slog.Any("error", err))
		panic(err)
	}
	return &TranslationsHandler{
		client: pb.NewTranslationsClient(conn),
		logger: logger,
	}
}

type translateRequest struct {
	Text       string `json:"text" binding:"required"`
	SourceLang string `json:"source_lang" binding:"required"`
	TargetLang string `json:"target_lang" binding:"required"`
}

func (h *TranslationsHandler) Translate(c *gin.Context) {
	var req translateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()

	resp, err := h.client.Translate(ctx, &pb.TranslateRequest{
		Text:       req.Text,
		SourceLang: req.SourceLang,
		TargetLang: req.TargetLang,
	})
	if err != nil {
		h.logger.Error("Translation service error", slog.Any("error", err))
		c.JSON(http.StatusBadGateway, gin.H{"error": "translation service unavailable"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
