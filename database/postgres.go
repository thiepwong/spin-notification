package database

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/spin-notification/common"
)

type Database struct {
	Db *pg.DB
}

func NewDb(cfg *common.Db) (*Database, error) {
	db := pg.Connect(&pg.Options{Addr: cfg.Host + ":" + cfg.Port, User: cfg.Username, Password: cfg.Password, Database: cfg.DbName})
	//	db.Close()
	return &Database{Db: db}, nil
}

func (db *Database) Request(plate string) (*common.VehicleHolder, error) {
	var result common.VehicleHolder
	var var1 string
	var var2 string
	_, e := db.Db.QueryOne(pg.Scan(&var1, &var2), `select * from get_activated_member(?)`, plate)
	if e != nil {
		return nil, e
	}
	result.Mobile = var1
	result.VehiclePlate = var2
	return &result, e
}
