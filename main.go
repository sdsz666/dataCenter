package main

import (
	"myDemo/dataFactory/utils/dbTools"
	"myDemo/dataFactory/api/controllers"
)

func main() {
	defer dbTools.Eloquent.Close()
	router := controllers.InitRouter()
	router.Run(":8000")
}