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
	unibase.InitConfig("LoginServer", true, "20130322")
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
	query_string := fmt.Sprintf(`select charid,accountid,accountname,charname form %s where charname like '%\%%'`, tableName)
	rows, err := db_login.Query(query_string)
	if err != nil {
		logging.Error("select channel_accounts err:%s", err.Error())
		return
	}
	defer rows.Close()
	var index uint64 = 0
	for rows.Next() {
		var charid uint64
		var accountid uint64
		var accountname string
		var charname string
		if err := rows.Scan(&charid, &accountid, &accountname, &charname); err != nil {
			logging.Error("db_login err:%s", err.Error())
			continue
		}
		logging.Info("准备处理第%d个玩家玩家信息，处理之前数据 charid %d, accountid %d, accountname %s, charname %s", index, charid, accountid, accountname, charname)
		index = index + 1
		/*
			query_string := fmt.Sprintf("replace into %s(zoneid,accountid,accountname,charname) values(%d,%d,'%s','%d')", tableName, 100, id, plataccount, id)
			_, err := db_zone.Exec(query_string)
			if err != nil {
				logging.Error("insert error %d, %s %s", index, query_string, err.Error())
			} else {
				logging.Info("process %d ok", index)
				index += 1
			}
		*/
	}
}
