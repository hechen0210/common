/**
@Time : 2019/11/25 15:51
@Author : hechen
@File : mongo
@Software: GoLand
*/
package mongo

import (
	"gopkg.in/mgo.v2"
	"time"
)

type Config struct {
	Host     []string
	DataBase string
	User     string
	Password string
}

type Mongo struct {
	Client *mgo.Session
	Error  error
}

func (c Config) New() *Mongo {
	client, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:     c.Host,
		Timeout:   time.Second * 1,
		Database:  c.DataBase,
		Username:  c.User,
		Password:  c.Password,
		PoolLimit: 1024,
	})
	if err != nil {
		return &Mongo{
			Client: client,
			Error:  err,
		}
	}
	return &Mongo{
		Client: nil,
		Error:  nil,
	}
}
