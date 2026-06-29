package cmdmigrate

import (
	"context"
	"fmt"

	"github.com/gosom/google-maps-scraper/migrations"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Run database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "database-url",
				Usage:    "PostgreSQL connection string",
				Sources:  cli.EnvVars("DATABASE_URL"),
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			dsn := cmd.String("database-url")
			fmt.Println("Running database migrations...")
			n, err := migrations.RunWithDSN(dsn)
			if err != nil {
				return fmt.Errorf("migrations failed: %w", err)
			}
			fmt.Printf("Applied %d migration(s)\n", n)
			return nil
		},
	}
}
