package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsongyan/wblog/app/models"
	"flag"
	"github.com/jinzhu/gorm"
	"os"
	"io"
	"github.com/claudiu/gocron"
	"github.com/wangsongyan/wblog/app/controllers"
	"github.com/wangsongyan/wblog/config"
	"github.com/wangsongyan/wblog/config/routers"
)

func main() {
	initConfig()
	initLog()
	defer initDB().Close()

	gin.SetMode(config.GetConfiguration().Mode) //启动模式
	app := gin.Default()
	routers.SetRouter(app)
	startTasks()
	app.Run(config.GetConfiguration().Addr)
}

func startTasks() {
	//Periodic tasks
	gocron.Every(1).Day().Do(controllers.CreateXMLSitemap)
	gocron.Every(7).Days().Do(controllers.Backup)
	gocron.Start()
}
func initLog() {
	//log
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DisableConsoleColor()
}

//config
func initConfig() {
	configFilePath := flag.String("C", "config/config.yaml", "config file path")
	if err := config.LoadConfiguration(*configFilePath); err != nil {
		//log err parsing config log file
		return
	}
}

//db
func initDB() *gorm.DB {
	db, err := models.InitDB()
	if err != nil {
		// log err open databases
		return nil
	}
	return db
}
