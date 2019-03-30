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
	//	defer db.Close()
	return &Database{Db: db}, nil
}

func (db *Database) Request(plate string) *common.VehicleHolder {
	result := &common.VehicleHolder{}
	db.Db.QueryOne(result, `select * from get_activated_member(?)`, plate)
	return result
}
