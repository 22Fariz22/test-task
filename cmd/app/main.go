package main

import "log"
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