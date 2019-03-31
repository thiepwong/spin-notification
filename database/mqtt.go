package database

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/thiepwong/spin-notification/common"
)

//MQTT struct
type MQTT struct {
	Client mqtt.Client
	Config common.Db
}

var handleMsg mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

//CreateClient func
func CreateClient(cf common.Db) MQTT {
	opts := mqtt.NewClientOptions().AddBroker("ws://" + cf.Host + ":" + cf.Port).SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(handleMsg)
	opts.SetPingTimeout(1 * time.Second)
	return MQTT{mqtt.NewClient(opts), cf}
}

//Connect func
func (c *MQTT) Connect() {
	if token := c.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Client.Subscribe(c.Config.DbName, 0, handleMsg); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}
