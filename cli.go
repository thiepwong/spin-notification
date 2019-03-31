package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"

	"github.com/thiepwong/spin-notification/common"
	"github.com/thiepwong/spin-notification/database"
)

type conf struct {
	Time int64 `yaml:"time"`
}

func main() {
	common.DrawLogo()
	cfg, db := initCheck()
	color.Green("SPIN COMMUNICATE GATEWAY STARTED")

	mq := database.CreateClient(cfg.Db.Mqtt)

	if token := mq.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	mq.Client.Subscribe(cfg.Db.Mqtt.DbName, 0x00, func(client mqtt.Client, msg mqtt.Message) {
		if msg != nil {
			kq := db.Request(string(msg.Payload()))
			color.Yellow("Testing: %s ", kq.Mobile)
			fmt.Println(kq.Mobile)
		}
	})
	mq.Connect()

	for {

	}

}

func initCheck() (*common.Config, *database.Database) {
	var c common.Config
	//var c types.Config
	color.Blue("Config checking...")
	time.Sleep(1 * time.Second)
	c.GetConf()
	if c.Db.Pg.Host == "" || c.Db.Pg.Port == "" || c.Db.Pg.Username == "" {
		color.Red("PostgreSQL database config load failed!")
		os.Exit(1)
	}
	color.Green("PostgreSQL database config pass!")

	if c.Db.Mqtt.Host == "" || c.Db.Mqtt.Port == "" || c.Db.Mqtt.DbName == "" {
		color.Red("MQTT broker config load failed!")
		os.Exit(2)
	}
	color.Green("MQTT broker config pass!")

	time.Sleep(2 * time.Second)
	color.Blue("Connecting database...")
	db, e := database.NewDb(&c.Db.Pg)
	if e != nil {
		color.Red("Database connect failed")
		color.Red(e.Error())
		os.Exit(0)
	}
	return &c, db
}
