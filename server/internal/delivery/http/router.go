// Package http contains router/controllers/utilities
package http

import (
	"github.com/kido1611/image-server/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app            *fiber.App
	fileController *controller.FileController
}

func NewRouter(app *fiber.App) *Router {
	fileController := controller.NewFileController()

	return &Router{
		app:            app,
		fileController: fileController,
	}
}

func (r *Router) Setup() {
	fileGroup := r.app.Group("/files")

	fileGroup.Get("/", r.fileController.ListFiles)
	fileGroup.Post("/", r.fileController.UploadFiles)
	fileGroup.Get("/:file", r.fileController.GetFile)
}
