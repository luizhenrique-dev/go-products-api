package main

import "github.com/luizhenrique-dev/go-products-api/configs"

func main() {
	config := configs.NewConfig()
	println(config.GetDBConnectionString())
	
}
