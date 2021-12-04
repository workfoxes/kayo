package server

import (
	"time"

	"github.com/workfoxes/calypso/pkg/client/db"

	"github.com/workfoxes/kayo/app/session"

	"github.com/gofiber/fiber/v2"
)

func PreRequest(ctx *fiber.Ctx) error {
	ctx.Locals("request_start_time", time.Now())
	_session := &session.BaseSession{
		DB: db.DB,
	}
	ctx.Locals("session", _session)
	return nil
}

func PostRequest(ctx *fiber.Ctx) error {
	ctx.Locals("request_end_time", time.Now())
	return nil
}
