package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func entry() error {
	var config Configuration
	configPath := "/var/aid/config.json"
	if err := config.Load(configPath); err != nil {
		return err
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run("127.0.0.1:9876")
}

func main() {
	err := entry()
	if err != nil {
		println(err.Error())
		if runtime.GOOS == "windows" {
			fmt.Println("Press enter to continue...")
			_, _ = fmt.Scanln()
		}
		os.Exit(1)
	}
}
