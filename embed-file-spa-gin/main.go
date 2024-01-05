package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "5050", "port server")
	flag.Parse()
	r := gin.Default()

	r.LoadHTMLFiles(filePathJoin("/web/dist/index.html"))

	r.Static("/assets", filePathJoin("/web/dist/assets"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":" + port)
}

func filePathJoin(filename string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(cwd, filename)
}
