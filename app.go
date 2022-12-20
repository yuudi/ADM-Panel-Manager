package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/yuudi/ADM-Panel-manager/panel"
	"github.com/yuudi/ADM-Panel-manager/routes"
)

func entry() error {
	var config panel.Configuration
	if err := config.Load(); err != nil {
		return err
	}

	panel.NewPanel(config)

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routes.RegisterRoutes(r)
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
