package main

import (
	"os"
	"time"

	"github.com/fatih/color"

	"github.com/thiepwong/spin-notification/common"
	"github.com/thiepwong/spin-notification/database"
)

type conf struct {
	Time int64 `yaml:"time"`
}

func main() {
	common.DrawLogo()
	cfg, data := initCheck()

	mq := database.CreateClient(cfg.Db.Mqtt, data)

	// if token := mq.Client.Connect(); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// }
	// ch, e := data.Request("23-0948221")
	// if e != nil {
	// 	fmt.Println(e.Error())
	// }
	// fmt.Println(ch)

	mq.Connect(data)

}

func initCheck() (*common.Config, *database.Database) {
	var c common.Config
	//var c types.Config
	color.Yellow("Config checking...")
	time.Sleep(1 * time.Second)
	c.GetConf()
	if c.Db.Pg.Host == "" || c.Db.Pg.Port == "" || c.Db.Pg.Username == "" {
		color.Red("-- PostgreSQL database config load failed!")
		os.Exit(1)
	}
	color.Green("-- PostgreSQL database config pass!")

	if c.Db.Mqtt.Host == "" || c.Db.Mqtt.Port == "" || c.Db.Mqtt.DbName == "" {
		color.Red("-- MQTT broker config load failed!")
		os.Exit(2)
	}
	color.Green("-- MQTT broker config pass!")

	time.Sleep(2 * time.Second)
	color.Yellow("Connecting database...")
	db, e := database.NewDb(&c.Db.Pg)
	if e != nil {
		color.Red("Database connect failed")
		color.Red(e.Error())
		os.Exit(0)
	}
	color.Green("-- Database connected!")
	return &c, db
}
