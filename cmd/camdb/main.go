package main

import (
	"flag"
	"fmt"

	"github.com/dynamicgo/config"
	"github.com/dynamicgo/slf4go"
	"github.com/go-xorm/xorm"
	"github.com/inwecrypto/camdb"
	_ "github.com/lib/pq"
)

var logger = slf4go.Get("camdb")
var configpath = flag.String("conf", "./camdb.json", "camdb config file")

func main() {

	flag.Parse()

	conf, err := config.NewFromFile(*configpath)

	if err != nil {
		logger.ErrorF("load eth indexer config err , %s", err)
		return
	}

	username := conf.GetString("camdb.username", "xxx")
	password := conf.GetString("camdb.password", "xxx")
	port := conf.GetString("camdb.port", "6543")
	host := conf.GetString("camdb.host", "localhost")
	scheme := conf.GetString("camdb.schema", "postgres")

	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("user=%v password=%v host=%v dbname=%v port=%v sslmode=disable", username, password, host, scheme, port))

	if err != nil {
		logger.ErrorF("create postgres orm engine err , %s", err)
		return
	}

	tables := []interface{}{
		new(camdb.Tx),
		new(camdb.Block),
		new(camdb.UTXO),
		new(camdb.Order),
		new(camdb.Wallet),
	}

	if err := engine.Sync2(tables...); err != nil {
		logger.ErrorF("sync table schema error , %s", err)
		return
	}

}
