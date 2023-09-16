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
	logger := config.GetInternalLogger()

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

	hasUserScopeService := services.NewHasUserScopeService(userPermissionsDAO)
	setUserPermissionsService := services.NewSetUserPermissionsService(userPermissionsDAO)

	hasUserScopeHandler := handlers.NewHasUserScopeHandler(hasUserScopeService)
	setUserPermissionsHandler := handlers.NewSetUserPermissionsHandler(setUserPermissionsService)

	router := apis.GetRouter(apis.RouterConfig{
		Logger:    logger,
		ProjectID: config.Deploy.ProjectID,
		Prod:      config.ENV == config.ProdENV,
		Health: map[string]apis.HealthChecker{
			"postgres": func() error {
				return postgres.PingContext(ctx)
			},
		},
	})

	router.POST("/user", setUserPermissionsHandler.Handle)
	router.GET("/user/scopes", hasUserScopeHandler.Handle)

	if err := router.Run(fmt.Sprintf(":%d", config.API.PortInternal)); err != nil {
		logger.Fatal().Err(err).Msg("a fatal error occurred while running the API, and the server had to shut down")
	}
}
