/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PaluMacil/dan2/applogger"
	"github.com/PaluMacil/dan2/config"
	"github.com/PaluMacil/dan2/database"
	"github.com/PaluMacil/dan2/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the routes for the dan2 web application",
	Long: `Serves the routes for the dan2 web application.

Detailed configuration is not yet available.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			fx.Provide(database.NewEntClient, config.NewAppEnv, applogger.NewLogger),
			fx.Invoke(server.Serve),
		)

		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
