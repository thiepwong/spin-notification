package main

import (
	"fmt"
	"net/url"

	"github.com/thiepwong/spin-notification/database"

	"github.com/thiepwong/spin-notification/common"
)

type conf struct {
	Time int64 `yaml:"time"`
}

func main() {
	var c common.Config
	//var c types.Config
	c.GetConf()

	db, e := database.NewDb(&c.Db.Pg)
	if e != nil {
		fmt.Println(e.Error())
	}

	database.Listen(&url.URL{Host: c.Db.Mqtt.Host + ":" + c.Db.Mqtt.Port}, "xinchao")
	kq := db.Request("23-0948221")

	fmt.Println(c, kq)
}
