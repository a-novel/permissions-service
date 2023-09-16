package main

import (
	"context"
	"fmt"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/go-apis"
	"github.com/a-novel/permissions-service/config"
	"github.com/a-novel/permissions-service/migrations"
	"github.com/a-novel/permissions-service/pkg/dao"
	"github.com/a-novel/permissions-service/pkg/handlers"
	"github.com/a-novel/permissions-service/pkg/services"
	"io/fs"
)

func main() {
	ctx := context.Background()
	logger := config.GetLogger()
	authClient := config.GetAuthClient(logger)

	postgres, sql, err := bunovel.NewClient(ctx, bunovel.Config{
		Driver:                &bunovel.PGDriver{DSN: config.Postgres.DSN, AppName: config.App.Name},
		Migrations:            &bunovel.MigrateConfig{Files: []fs.FS{migrations.Migrations}},
		DiscardUnknownColumns: true,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("error connecting to postgres")
	}
	defer func() {
		_ = postgres.Close()
		_ = sql.Close()
	}()

	userPermissionsDAO := dao.NewUserPermissionsRepository(postgres)

	getUserScopesService := services.NewGetUserScopesService(userPermissionsDAO, authClient)

	getUserScopesHandler := handlers.NewGetUserScopesHandler(getUserScopesService)

	router := apis.GetRouter(apis.RouterConfig{
		Logger:    logger,
		ProjectID: config.Deploy.ProjectID,
		CORS:      apis.GetCORS(config.App.Frontend.URLs),
		Prod:      config.ENV == config.ProdENV,
		Health: map[string]apis.HealthChecker{
			"postgres": func() error {
				return postgres.PingContext(ctx)
			},
			"auth-client": func() error {
				return authClient.Ping(ctx)
			},
		},
	})

	router.GET("/user/scopes", getUserScopesHandler.Handle)

	if err := router.Run(fmt.Sprintf(":%d", config.API.Port)); err != nil {
		logger.Fatal().Err(err).Msg("a fatal error occurred while running the API, and the server had to shut down")
	}
}
