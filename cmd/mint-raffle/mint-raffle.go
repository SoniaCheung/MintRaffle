package main

import (
	"fmt"
	"soniacheung/mint-raffle/cmd/mint-raffle/app"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

func main() {
	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", "root:Root1234@tcp(127.0.0.1:3306)/MintRaffle?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	defer engine.Close()

	fmt.Println("MySQL DB connection established")

	routers, err := app.NewRouters()
	if err != nil {
		logrus.Fatalf("Create routers: %v", err)
	}

	routers.Run(":8080")
}
