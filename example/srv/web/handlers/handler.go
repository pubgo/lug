package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/lava/service"
	"go.uber.org/zap"
)

func New(l *zap.Logger) service.WebHandler {
	return &Handler{L: l}
}

type Handler struct {
	L *zap.Logger
}

func (t *Handler) Router(r fiber.Router) {
	r.Get("/hello", t.Get)
}

func (t *Handler) Get(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"hello": "ok"})
}
