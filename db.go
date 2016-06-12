package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/smtc/glog"
)

var (
	dbs *sql.DB
)

func init() {
	deferinit.AddInit(sqlConntion, sqlClose, 20)
}

/**
数据库连接
创建人:邵炜
创建时间:2016年3月7日11:24:48
*/
func sqlConntion() {

	var (
		err error
	)

	dbuser := config.GetStringMust("dbuser")
	dbhost := config.GetStringMust("dbhost")
	dbport := config.GetIntDefault("dbport", 3306)
	dbpass := config.GetStringMust("dbpass")
	dbname := config.GetStringMust("dbname")

	dbclause := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&allowAllFiles=true", dbuser, dbpass, dbhost, dbport, dbname)

	dbs, err = sql.Open("mysql", dbclause)

	if err != nil {
		glog.Error("mysql can't connection %s \n", err.Error())
		return
	}

	err = dbs.Ping()

	if err != nil {
		glog.Error("mysql can't ping , err: %s \n", err.Error())
		return
	}

	glog.Info("mysql is open \n")
}

/**
数据库关闭
创建人:邵炜
创建时间:2016年3月7日11:26:23
*/
func sqlClose() {
	err := dbs.Close()

	if err != nil {
		glog.Error("mysql can't close ,err: %s \n", err.Error())
	}
}

/**
查询方法
创建人:邵炜
创建时间:2015年12月29日17:26:41
输入参数: sqlstr 要执行的sql语句 param执行SQL的语句参数化传递
输出参数: 查询返回条数  错误对象输出
*/
func sqlSelect(sqlStr string, param ...interface{}) (*sql.Rows, error) {

	var (
		row *sql.Rows
		err error
	)

	err = dbs.Ping()

	if err != nil {
		glog.Error("mysql can't ping %s \n", err.Error())
		sqlClose()
		sqlConntion()
	}

	if param == nil {
		row, err = dbs.Query(sqlStr)
	} else {
		row, err = dbs.Query(sqlStr, param...)
	}

	if err != nil {
		glog.Error("mysql query can't select sql: %s err: %s \n", sqlStr, err.Error())
		return nil, err
	}

	return row, nil
}

/**
增删改查方法
创建人:邵炜
创建时间:2015年12月29日17:33:06
输入参数: sqlstr 要执行的sql语句  param执行SQL的语句参数化传递
输出参数: 执行结果对象  错误对象输出
*/
func sqlExec(sqlStr string, param ...interface{}) (sql.Result, error) {
	var (
		exec sql.Result
		err  error
	)

	err = dbs.Ping()

	if err != nil {
		glog.Error("mysql can't ping %s \n", err.Error())
		sqlClose()
		sqlConntion()
	}

	if param == nil {
		exec, err = dbs.Exec(sqlStr)
	} else {
		exec, err = dbs.Exec(sqlStr, param...)
	}

	if err != nil {
		glog.Error("mysql exec can't carried out sql: %s err: %s \n", sqlStr, err.Error())
		return nil, err
	}

	return exec, nil
}
