/**
@Time : 2019/11/22 14:49
@Author : hechen
@File : mysql
@Software: GoLand
*/
package mysql

import (
	"common/helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Config struct {
	Host          string
	User          string
	Password      string
	Port          string
	DbName        string
	Prefix        string
	SingularTable bool
}

type DB struct {
	Db    *gorm.DB
	Error error
}

func (c Config) New() (db DB) {
	connectStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	connect := fmt.Sprintf(connectStr, c.User, c.Password, c.Host, c.Port, c.DbName)
	db.Db, db.Error = gorm.Open("mysql", connect)
	if db.Error != nil {
		return db
	}
	db.Db.SingularTable(c.SingularTable)
	db.Db.DB().SetConnMaxLifetime(time.Hour * 4)
	db.Db.DB().SetMaxIdleConns(10)
	db.Db.DB().SetMaxIdleConns(100)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if strings.HasPrefix(defaultTableName, c.Prefix) {
			return defaultTableName
		}
		return c.Prefix + defaultTableName
	}

	return db
}

/**
批量插入
*/
func (db DB) BatchInsert(tableName string, field []string, data [][]string) {
	insert := "insert into " + tableName + "(" + strings.Join(field, ",") + ") values "
	fieldLen := len(field)
	for i := 0; i < fieldLen; i++ {
		insert += "("
		insert += strings.Repeat("?,", len(field))
		insert = helper.SubStrByEnd(insert, 0, -1)
		insert += "),"
	}
	insert = helper.SubStrByEnd(insert, 0, -1)
	dataLen := len(data)
	values := []interface{}{}
	for i := 0; i < dataLen; i++ {
		if len(data[i]) == fieldLen {
			values = append(values, data[i])
		}
	}
	if len(values) > 0 {
		db.Db.Exec(insert, values...)
	}
}
