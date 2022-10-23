package dao

import (
	"database/sql"
	"github.com/RTE/web/api/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
import "github.com/go-gorp/gorp"

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", util.Env("MYSQL_USER")+":"+util.Env("MYSQL_PASSWORD")+"@tcp("+util.Env("MYSQL_HOST")+":"+util.Env("MYSQL_PORT")+")/rte_db?parseTime=true")
	//db, err := sql.Open("mysql", "appuser:hogehogekenshin@tcp(mysql:3306)/rte_db?parseTime=true")
	//log.Print(util.Env("MYSQL_USER")+":"+util.Env("MYSQL_PASSWORD")+"@tcp("+util.Env("MYSQL_HOST")+":"+util.Env("MYSQL_PORT")+")/rte_db?parseTime=true")
	//log.Print("appuser:hogehogekenshin@tcp(mysql:3306)/rte_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
}
