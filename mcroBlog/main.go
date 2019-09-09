package main

import (
	"fmt"
	"go_mod-project/mcroBlog/dbclient"
	"go_mod-project/mcroBlog/service"
)

var appName = "accountService"

func main()  {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient()  {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}