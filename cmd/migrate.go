/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/PaluMacil/dan2/database"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "automigrates the database",
	Long:  `migrates the database--currently uses automigrate and assumes the dev database`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// TODO: use a central config; conn str host and port inside compose would be: postgresql://dev:dev@db:5400/dan2
		//       outside would be postgresql://dev:dev@localhost:5400/dan2
		dbHost := os.Getenv("DAN2_TEST_DB_HOST")
		if dbHost == "" {
			dbHost = "localhost"
		}
		connectionString := fmt.Sprintf("postgresql://dev:dev@%s:5400/dan2", dbHost)
		client, err := database.Open(connectionString)
		if err != nil {
			log.Fatalf("failed creating db client: %v", err)
		}
		defer client.Close()
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
