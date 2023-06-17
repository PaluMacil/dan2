package database

import (
	"context"
	"fmt"
	"github.com/PaluMacil/dan2/ent"
)

// Migrate creates the schema a then closes the connection for this client
func Migrate(entClient *ent.Client) error {
	ctx := context.Background()
	if err := entClient.Schema.Create(ctx); err != nil {
		entClient.Close()
		return fmt.Errorf("failed creating schema: %v", err)
	}
	return entClient.Close()
}
