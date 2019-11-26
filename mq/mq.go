/**
@Time : 2019/11/25 15:53
@Author : hechen
@File : mq
@Software: GoLand
*/
package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Exchange string
}

type MQ struct {
	Client *amqp.Channel
	Error  error
}

func (c Config) New() *MQ {
	dial := "amqp://%s:%s@%s:%s/"
	conn, err := amqp.Dial(fmt.Sprintf(dial, c.User, c.Password, c.Host, c.Port))
	if err != nil {
		return &MQ{
			Client: nil,
			Error:  err,
		}
	}
	client, err := conn.Channel()
	if err != nil {
		return &MQ{
			Client: client,
			Error:  err,
		}
	}
	err = client.ExchangeDeclare(c.Exchange, amqp.ExchangeDirect, true, false, false, false, nil)
	return &MQ{
		Client: client,
		Error:  err,
	}
}

/**
消费消息
*/
func (mq *MQ) Consume(name, exchange string) (delivery <-chan amqp.Delivery, err error) {
	queue, err := mq.Client.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		return delivery, err
	}
	err = mq.Client.QueueBind(queue.Name, queue.Name, exchange, false, nil)
	if err != nil {
		return delivery, err
	}
	delivery, err = mq.Client.Consume(queue.Name, "", false, false, false, false, nil)
	return delivery, err
}
