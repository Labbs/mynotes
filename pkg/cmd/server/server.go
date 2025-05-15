package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labbs/mynotion/internal/database"
	"github.com/labbs/mynotion/pkg/config"
	"github.com/labbs/mynotion/pkg/flags"
	htserver "github.com/labbs/mynotion/pkg/httpserver"
	logger "github.com/labbs/mynotion/pkg/logger"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewInstance() *cli.Command {
	serverFlags := getFlags()

	return &cli.Command{
		Name:   "server",
		Usage:  "Start the server",
		Flags:  serverFlags,
		Before: altsrc.InitInputSourceWithContext(serverFlags, altsrc.NewJSONSourceFromFlagFunc("config")),
		Action: runServer,
	}
}

func getFlags() (list []cli.Flag) {
	// Add the flags to the list
	// and return the list
	list = append(list, flags.GenericFlags()...)
	list = append(list, flags.ServerFlags()...)
	list = append(list, flags.DatabaseFlags()...)
	list = append(list, flags.LoggerFlags()...)
	list = append(list, flags.SessionFlags()...)
	return
}

func runServer(c *cli.Context) error {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	l := logger.NewLogger(config.Logger.Level, config.Logger.Pretty, c.App.Version)

	db := database.NewGorm(l, config.Database.Dialect, config.Database.DSN)

	// Start the HTTP server
	var httpServer htserver.Config
	httpServer.Port = config.Server.Port
	httpServer.HttpLogs = config.Server.HttpLogs
	httpServer.Logger = l
	httpServer.Stop = stopChan
	httpServer.Db = db

	httpServer.NewServer()

	// Wait for the stop signal
	<-stopChan

	if err := httpServer.Shutdown(); err != nil {
		httpServer.Logger.Error().Err(err).Msg("failed to shutdown server")
		return err
	}

	httpServer.Logger.Info().Msg("Server stopped")

	return nil
}
