/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/PaluMacil/dan2/applogger"
	"github.com/PaluMacil/dan2/config"
	"github.com/PaluMacil/dan2/database"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "automigrates the database",
	Long:  `migrates the database--currently uses automigrate and assumes the dev database`,
	Run: func(cmd *cobra.Command, args []string) {
		appEnv := config.NewAppEnv()
		logger := applogger.NewLogger(appEnv)
		logger.Info("running migrate command")
		entClient, err := database.NewEntClient(appEnv)
		if err != nil {
			logger.Error("creating new ent client", slog.String("error", err.Error()))
			os.Exit(1)
		}
		err = database.Migrate(entClient)
		if err != nil {
			logger.Error("migrating database", slog.String("error", err.Error()))
			os.Exit(1)
		}
		logger.Info("migrate command complete")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
