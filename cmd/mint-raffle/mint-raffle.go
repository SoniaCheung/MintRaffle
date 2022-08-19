package main

import (
	"soniacheung/mint-raffle/cmd/mint-raffle/app"

	"github.com/sirupsen/logrus"
)

func main() {
	routers, err := app.NewRouters()
	if err != nil {
		logrus.Fatalf("Create routers: %v", err)
	}

	routers.Run(":8080")
}
