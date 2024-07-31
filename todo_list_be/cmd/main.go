package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"todo_list/config"
	"todo_list/database"
	"todo_list/logger"
	"todo_list/server"
)

func main() {
	// 1. Load config
	config.Load()

	// 2. Init logger
	_, err := logger.InitLogger(config.GlobalCfg.Logger)
	if err != nil {
		log.Fatalf("===== Init logger failed: %+v", err.Error())
	}

	// 3. Init db
	err = database.GetDBInstance().Open(config.GlobalCfg.Database.MySQLConfig)
	if err != nil {
		logger.Fatalf("===== Open db failed: %+v", err.Error())
	}

	// 4. Init server
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name: "server",
				Action: func(ctx *cli.Context) error {
					srv := server.NewServer()
					return srv.Serve(config.GlobalCfg.Server.HTTP)
				},
			},
		},
	}

	// 4. Run server
	if err = app.Run(os.Args); err != nil {
		logger.Fatalf("===== Run server failed: %+v", err.Error())
	}
}
