package main

import (
	"database/sql"
	"fmt"
	"strings"

	"git.code4.in/mobilegameserver/config"
	"git.code4.in/mobilegameserver/logging"
	"git.code4.in/mobilegameserver/unibase"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db_login *sql.DB
	db_zone  *sql.DB
)

func main() {
	if unibase.InitConfig("LoginServer", true) == false {
		return
	}
	mysqlurl := config.GetConfigStr("mysql")
	tableName := config.GetConfigStr("charnames")
	logging.Info("connect mysql %s", mysqlurl)
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	db, err := sql.Open("mysql", mysqlurl)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	db_login = db

	mysqlurl = config.GetConfigStr("mysql_zone")
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	logging.Info("connect mysql %s", mysqlurl)
	db2, err := sql.Open("mysql", mysqlurl)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	db_zone = db2
	rows, err := db_login.Query(`select id, plataccount from channel_accounts`)
	if err != nil {
		logging.Error("select channel_accounts err:%s", err.Error())
		return
	}
	defer rows.Close()
	var index uint64 = 0
	for rows.Next() {
		var plataccount string
		var id uint64
		if err := rows.Scan(&id, &plataccount); err != nil {
			logging.Error("db_login err:%s", err.Error())
			continue
		}
		query_string := fmt.Sprintf("replace into %s(zoneid,accountid,accountname,charname) values(%d,%d,'%s','%d')", tableName, 100, id, plataccount, id)
		_, err := db_zone.Exec(query_string)
		if err != nil {
			logging.Error("insert error %d, %s %s", index, query_string, err.Error())
		} else {
			logging.Info("process %d ok", index)
			index += 1
		}
	}
}
