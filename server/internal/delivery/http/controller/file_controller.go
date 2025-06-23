// Package controller contains all used controllers
package controller

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kido1611/image-server/internal/model"
	"github.com/kido1611/image-server/internal/utils"
)

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
}

func (c *FileController) ListFiles(ctx *fiber.Ctx) error {
	utils.CheckAndCreateDirectoryIfNotExists("files")

	files, err := os.ReadDir("files")
	if err != nil {
		return fiber.ErrInternalServerError
	}

	imageExtensions := []string{".jpg", ".jpeg", ".png"}

	var images []model.Image

	if err != nil {
		return fiber.ErrInternalServerError
	}

	for _, file := range files {
		fileExtension := strings.ToLower(filepath.Ext("files/" + file.Name()))
		if slices.Contains(imageExtensions, fileExtension) {
			images = append(images, model.Image{
				Name: file.Name(),
				URL:  fmt.Sprintf("%s/files/%s", ctx.BaseURL(), file.Name()),
			})
		}
	}

	if len(images) == 0 {
		return ctx.JSON([]any{})
	}

	return ctx.JSON(images)
}

func (c *FileController) UploadFiles(ctx *fiber.Ctx) error {
	utils.CheckAndCreateDirectoryIfNotExists("files")

	file, err := ctx.FormFile("file")
	if err != nil {
		return fiber.ErrBadRequest
	}

	imageExtensions := []string{".jpg", ".jpeg", ".png"}

	fileExtension := strings.ToLower(filepath.Ext("files/" + file.Filename))
	// don't save file if not image
	if !slices.Contains(imageExtensions, fileExtension) {
		return ctx.SendStatus(201)
	}

	newName := fmt.Sprintf("%s%s", uuid.New().String(), fileExtension)
	err = ctx.SaveFile(file, fmt.Sprintf("files/%s", newName))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(201)
}

func (c *FileController) GetFile(ctx *fiber.Ctx) error {
	param := ctx.Params("file")

	filePath := fmt.Sprintf("files/%s", param)
	_, err := os.Stat(filePath)
	if err != nil {
		return fiber.ErrNotFound
	}

	return ctx.SendFile(filePath)
}
