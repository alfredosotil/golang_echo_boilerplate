package main

import (
	application "echo-demo-project"
	"echo-demo-project/config"
	"echo-demo-project/docs"
	"fmt"
)

// @title Echo App
// @version 1.0
// @description This is a demo version of Echo app.

// @contact.name Alfredo Sotil P.
// @contact.url https://www.linkedin.com/in/alfredo-antonio-sotil-pastor-9651837b/
// @contact.email alfredosotil@gmail.com

// @securityDefinitions.apikey ApiKeyAuthNIX Solutions
// @in header
// @name Authorization

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
