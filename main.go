package main

import (
	"be-fullstack-vue-go/config"
	"be-fullstack-vue-go/database"
	"be-fullstack-vue-go/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	//setup router
	r := routes.SetupRouter()

	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
