package database

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
)

const dbFileName = "llmos-ui.db"

func RegisterDBClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", dbFileName))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err = client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, nil
}
