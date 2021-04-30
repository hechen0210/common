/**
@Time : 2019/11/22 14:49
@Author : hechen
@File : mysql
@Software: GoLand
*/
package mysql

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hechen0210/common/helper"
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
	Charset       string
}

type DB struct {
	Client *gorm.DB
	Error  error
}

func (c Config) New() *DB {
	charset := "utf8"
	if c.Charset != "" {
		charset = c.Charset
	}
	connectStr := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	connect := fmt.Sprintf(connectStr, c.User, c.Password, c.Host, c.Port, c.DbName, charset)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return &DB{
			Client: db,
			Error:  err,
		}
	}
	db.SingularTable(c.SingularTable)
	db.DB().SetConnMaxLifetime(time.Hour * 4)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxIdleConns(100)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if strings.HasPrefix(defaultTableName, c.Prefix) {
			return defaultTableName
		}
		return c.Prefix + defaultTableName
	}

	return &DB{
		Client: db,
		Error:  err,
	}
}

/**
批量插入
*/
func (db *DB) BatchInsert(tableName string, field []string, data [][]interface{}) error {
	insert := "insert into " + tableName + "(" + strings.Join(field, ",") + ") values "
	fieldLen := len(field)
	dataLen := len(data)
	for i := 0; i < len(data); i++ {
		insert += "("
		insert += strings.Repeat("?,", fieldLen)
		insert = helper.SubStrByEnd(insert, 0, -1)
		insert += "),"
	}
	insert = helper.SubStrByEnd(insert, 0, -1)
	values := []interface{}{}
	for i := 0; i < dataLen; i++ {
		if len(data[i]) == fieldLen {
			values = append(values, data[i]...)
		}
	}
	if len(values) > 0 {
		return db.Client.Exec(insert, values...).Error
	}
	return errors.New("data is empty")
}
