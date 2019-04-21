package main

import (
	"dataCenter/utils/dbTools"
	"dataCenter/api/controllers"
)

func main() {
	defer dbTools.Eloquent.Close()
	router := controllers.InitRouter()
	router.Run(":8000")
}