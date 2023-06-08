package cmd

import (
	"github.com/atadzan/gin-blog/pkg/config"
	"github.com/atadzan/gin-blog/pkg/routing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in config.yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	// Added show logged line flag
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// Set configuration from config.yml file
	config.Set()

	// Initializing router
	routing.Init()

	// Get router
	router := routing.GetRouter()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.GetString("App.Name"),
		})
	})
	routing.Serve()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
