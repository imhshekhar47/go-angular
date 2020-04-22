package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-angular/middleware"
)

func main() {
	server := gin.Default()

	server.Use(middleware.BasicAuth())
	server.Static("/static", "./static")

	server.NoRoute(func(ctx *gin.Context) {
		dir, file := path.Split(ctx.Request.RequestURI)
		ext := filepath.Ext(file)

		if file == "" || ext == "" {
			ctx.File("./webapp/dist/webapp/index.html")
		} else {
			ctx.File("./webapp/dist/webapp/" + path.Join(dir, file))
		}

	})

	apiRouts := server.Group("/api")
	{
		apiRouts.GET("/health", func(ctx *gin.Context) {
			ctx.Writer.Header().Set("host", os.Getenv("hostname"))
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})
	}

	const port = 8080

	if err := server.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(fmt.Sprintf("Fatal Error: %d", err.Error))
	} else {
		log.Printf("Server stared on port %d\n", port)
	}

}
