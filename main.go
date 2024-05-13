package main

import (
	"wh-hard01.kol.wb.ru/wh-tech/wh-tech-back/wh_tech_gitlab_api/internal/service"
	"wh-hard01.kol.wb.ru/wh_core/gocore_rest_auto_api"
)

func main() {
	runner := rest_auto_api.NewService()
	runner.DisableJWTTokenAuth()
	gitApi := service.NewGitLabApiService()
	runner.RegisterConfigurableEntity(gitApi.Init)
	runner.RegisterCustomHandler("GetUser", gitApi.GetUser)
	runner.RegisterCustomHandler("ListUsers", gitApi.GetUsers)
	runner.RegisterCustomHandler("Registration", gitApi.Registration)
	runner.RegisterCustomHandler("BlockUser", gitApi.BlockUser)
	runner.RegisterCustomHandler("CheckValidToken", gitApi.CheckValidToken)
	runner.Run()
}
