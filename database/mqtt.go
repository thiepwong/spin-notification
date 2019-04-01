package database

import (
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"
	"github.com/thiepwong/spin-notification/common"
)

//MQTT struct
type MQTT struct {
	Client mqtt.Client
	Config common.Db
	Db     *Database
}

// var handleMsg mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

// 	fmt.Printf("TOPIC: %s\n", msg.Topic())
// 	fmt.Printf("MSG: %s\n", msg.Payload())
// }

//CreateClient func
func CreateClient(cf common.Db, db *Database) MQTT {
	opts := mqtt.NewClientOptions().AddBroker("ws://" + cf.Host + ":" + cf.Port).SetClientID("gotrivial")
	//	opts.SetDefaultPublishHandler(handleMsg)
	opts.SetPingTimeout(200 * time.Millisecond)
	var mq MQTT
	mq.Client = mqtt.NewClient(opts)
	mq.Config = cf
	mq.Db = db
	return mq
}

//Connect func
func (c *MQTT) Connect(db *Database) {
	token := c.Client.Connect()

	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		//	panic(token.Error())
		os.Exit(2)
	}

	if c.Client.IsConnected() {
		color.Green("MQTT is connected to the broker!")
	}

	fmt.Println()
	color.HiGreen("Configuration checking completed!")
	fmt.Println()
	color.Green("SPIN COMMUNICATE GATEWAY STARTED")

	c.Client.Publish("parking/input", 0, false, "ping")

	token = c.Client.Subscribe("parking/input", 0x01, func(client mqtt.Client, msg mqtt.Message) {
		message := strings.Split(string(msg.Payload()), ",")
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Println(message[0])
		fmt.Printf("MSG: %s\n", msg.Payload())

		if len(message) > 1 {
			channel, e := db.Request(message[0])
			if e != nil {
				color.Red(e.Error())
			}
			common.SendMessage(channel, message[1])
		}

	})

	if token.Wait() && token.Error() != nil && token.WaitTimeout(5*time.Millisecond) {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {

		//	fmt.Println("...")

	}

}
