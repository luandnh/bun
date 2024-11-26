/*
TODO:
	- Commands:
		- init - Create migration+locks tables [--cmd to add cmd/ folder structure]
		- migrate - Apply database migrations
		- rollback - Rollback the last migration group
		- create - Create template SQL migration filex [--go | --sql | --transactional]
		- unlock - Unlock locks table
	- provide NewCommand() *cli.Command intead of the cli.App, so that buncli could be embeded in the existing CLIs
	- configure logging and verbosity
	- (experimental, low prio) add FromPlugin() to read config from plugin and use from cmd/bundb.
*/
package buncli

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

// bunApp is the root-level bundb app that all other commands attach to.
var bunApp = &cli.App{
	Name:  "bundb",
	Usage: "Database migration tool for uptrace/bun",
}

// New creates a new CLI application for managing bun migrations.
func New(c *Config) *App {
	bunApp.Commands = cli.Commands{
		CmdAuto(c),
	}
	return &App{
		App: bunApp,
	}
}

type Config struct {
	DB           *bun.DB
	AutoMigrator *migrate.AutoMigrator
}

// Run calls cli.App.Run and returns its error.
func Run(args []string, c *Config) error {
	return New(c).Run(args)
}

// RunCtx calls cli.App.RunContexta and returns its error.
func RunContext(ctx context.Context, args []string, c *Config) error {
	return New(c).RunContext(ctx, args)
}

// App is a wrapper around cli.App that extends it with bun-specific features.
type App struct {
	*cli.App
}
