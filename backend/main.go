package main

import (
	"backend/app/config"
	"backend/app/database"
	"backend/app/router/v1"
	"log"
)

func main() {

	database.InitDB()

	route := v1.InitRoute()

	err := route.Run(":" + config.ProjectPort)
	if err != nil {
		log.Printf("启动失败,失败原因 %s\n", err)
	}

}
