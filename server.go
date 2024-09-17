package main

import (
	"rsudlampung/helper"
	"rsudlampung/versions/group_01"
	"rsudlampung/versions/group_02"

	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db_01 *gorm.DB
var db_02 *gorm.DB

func main() {

	os.Setenv("TZ", "Asia/Jakarta")
	configEnv, errEnv := helper.LoadConfig("./")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}

	logFile := configEnv.LogFile
	if logFile == "on" {
		helper.SetupLogOutput()
	}

	ginModeEnv := configEnv.GinMode
	if ginModeEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	mainServer := gin.New()
	mainServer.Use(gin.Recovery())

	config := cors.DefaultConfig()
	origins := strings.Split(configEnv.AllowOrigin, ",")
	if configEnv.Env == "staging" {
		origins = append(origins, "http://localhost:3000")
	}
	config.AllowOrigins = origins

	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE"}
	config.AllowHeaders = []string{"Access-Control-Allow-Origin", "authorization", "content-type", "content-length", "user-agent", "Host", "accept"}
	log.Println("config", config.AllowOrigins)
	mainServer.Use(cors.New(config))
	mainServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "root path")
	})

	apiSistemRoutes := mainServer.Group("/sistem")
	db_01 = helper.OpenDB(configEnv.DB, configEnv.SCHEMA, "v01")
	db_02 = helper.OpenDB(configEnv.DB, configEnv.SCHEMA, "v02")

	//initiate services
	groupServer10 := group_01.NewGroupServer(apiSistemRoutes, db_01, "v01")
	groupServer10.Init()

	groupServer20 := group_02.NewGroupServer(apiSistemRoutes, db_02, "v02")
	groupServer20.Init()

	port := configEnv.Port
	mainServer.Run(":" + port)

}
