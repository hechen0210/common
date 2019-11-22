/**
@Time : 2019/11/22 14:49
@Author : hechen
@File : mysql
@Software: GoLand
*/
package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
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

type DB *gorm.DB

func (c Config) New() (db DB, err error) {
	connectStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	connect := fmt.Sprintf(connectStr, c.User, c.Password, c.Host, c.Port, c.DbName)
	db, err = gorm.Open("mysql", connect)
	if err != nil {
		return db, err
	}
	fmt.Printf("%+v",db)
	//db.SingularTable(c.SingularTable)
	//db.DB().SetConnMaxLifetime(time.Hour * 4)
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxIdleConns(100)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if strings.HasPrefix(defaultTableName, c.Prefix) {
			return defaultTableName
		}
		return c.Prefix + defaultTableName
	}
	return db, err
}

//func (db DB) BatchInsert(table string, field []string, data [][]string) {
//
//}
