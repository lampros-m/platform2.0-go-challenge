package server

import (
	"log"

	"gwi/platform2.0-go-challenge/environment"
	"gwi/platform2.0-go-challenge/internal/app/audience"
	"gwi/platform2.0-go-challenge/internal/app/authentication"
	"gwi/platform2.0-go-challenge/internal/app/charts"
	"gwi/platform2.0-go-challenge/internal/app/dashboard"
	"gwi/platform2.0-go-challenge/internal/app/insights"
	"gwi/platform2.0-go-challenge/internal/repositories/sql"
)

// Application : Struct that defines application services.
type Application struct {
	Client         *sql.Client
	Authentication *authentication.Service
	Insights       *insights.Service
	Charts         *charts.Service
	Audience       *audience.Service
	Dashboard      *dashboard.Service
}

// NewApplication : Returns an application instance.
func NewApplication(conf *environment.Config) *Application {
	client, err := sql.NewDBClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	authenticationRepo := sql.NewAuthenticationRepo(client)
	authenticationService := authentication.NewService(authenticationRepo)

	insightsRepo := sql.NewInsightsRepo(client)
	insightsService := insights.NewService(insightsRepo)

	chartsRepo := sql.NewChartsRepo(client)
	chartsService := charts.NewService(chartsRepo)

	audienceRepo := sql.NewAudienceRepo(client)
	audienceService := audience.NewService(audienceRepo)

	dashboardRepo := sql.NewDashboardRepo(client)
	dashboardService := dashboard.NewService(dashboardRepo)

	app := Application{
		Client:         client,
		Authentication: authenticationService,
		Insights:       insightsService,
		Charts:         chartsService,
		Audience:       audienceService,
		Dashboard:      dashboardService,
	}

	return &app
}
