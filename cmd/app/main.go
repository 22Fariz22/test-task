package main

import (
	"log"

	"github.com/test-task/internal/app"
	"github.com/test-task/internal/config"
)

func main(){
	//Configuration
	cfg,err:= config.NewConfig()
	if err!=nil{
		log.Fatalf("Config err: %s",err)
	}

	//Run
	app:= app.NewApp(cfg)
	app.Run()
}