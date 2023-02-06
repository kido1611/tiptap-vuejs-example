package main

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type image struct {
	Name string
	Url  string
}

func createDirectoryIfNotExist() {
	if _, err := os.Stat("files"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("files", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/files", func(c *gin.Context) {
		createDirectoryIfNotExist()

		var imageList []image

		files, err := ioutil.ReadDir("files")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			newImage := image{
				Name: file.Name(),
				Url:  "http://localhost:8080/files/" + file.Name(),
			}
			imageList = append(imageList, newImage)
		}

		c.JSON(http.StatusOK, imageList)
	})

	r.POST("/files", func(c *gin.Context) {
		createDirectoryIfNotExist()

		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}

		err = c.SaveUploadedFile(file, "files/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}

		fileExtension := filepath.Ext("files/" + file.Filename)

		newName := uuid.New().String() + fileExtension
		err = os.Rename("files/"+file.Filename, "files/"+newName)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, image{
			Name: file.Filename,
			Url:  "http://localhost:8080/files/" + newName,
		})
	})

	r.GET("/files/:file", func(c *gin.Context) {
		name := c.Param("file")

		createDirectoryIfNotExist()

		file, err := ioutil.ReadFile("files/" + name)
		if err != nil {
			log.Fatal(err)
		}

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		_, err = c.Writer.Write(file)
		if err != nil {
			return
		}
	})

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
