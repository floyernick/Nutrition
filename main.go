package main

import (
	"flag"

	"Nutrition/config"
	"Nutrition/controller"
	"Nutrition/presenter"
	"Nutrition/storage"
	"Nutrition/tools/logger"
)

func main() {

	environmentName := *flag.String("c", "default.yaml", "Environment config name")

	flag.Parse()

	config_, err := config.LoadConfig(environmentName)

	if err != nil {
		logger.Error(err.Error())
	}

	storage_, err := storage.Init(config_.Database)

	if err != nil {
		logger.Error(err.Error())
	}

	controller_ := controller.Init(storage_)

	err = presenter.Init(config_.Server, controller_)

	if err != nil {
		logger.Error(err.Error())
	}

}
