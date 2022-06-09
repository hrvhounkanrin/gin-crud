package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/setondji357/rest-insured/pkg/common/config"
	"github.com/setondji357/rest-insured/pkg/common/db"
	"github.com/setondji357/rest-insured/pkg/insured"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	co, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config")
	}
	r := gin.Default()
	h := db.Init(co)
	insured.RegisterRoutes(r, h)
	r.Run(co.Port)
}
