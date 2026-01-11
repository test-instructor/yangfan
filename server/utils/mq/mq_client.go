package mq

import (
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/test-instructor/yangfan/server/v2/config"
	"go.uber.org/zap"
)

type MQClient struct {
	cfg      config.MQ
	conn     *amqp.Connection
	channel  *amqp.Channel
	connLock sync.RWMutex
	isClosed bool
	logger   *zap.Logger
	done     chan bool
}

func NewMQClient(cfg config.MQ, logger *zap.Logger) (*MQClient, error) {
	client := &MQClient{
		cfg:    cfg,
		logger: logger,
		done:   make(chan bool),
	}

	if err := client.connect(); err != nil {
		return nil, err
	}

	go client.handleReconnect()

	return client, nil
}

func (c *MQClient) connect() error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.isClosed {
		return nil
	}

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		c.cfg.Username,
		c.cfg.Password,
		c.cfg.Host,
		c.cfg.Port,
		c.cfg.VirtualHost,
	)

	config := amqp.Config{
		Heartbeat: time.Duration(c.cfg.Heartbeat) * time.Second,
	}

	conn, err := amqp.DialConfig(dsn, config)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	c.conn = conn
	c.channel = ch

	c.logger.Info("MQ Connected", zap.String("host", c.cfg.Host))
	return nil
}

func (c *MQClient) handleReconnect() {
	for {
		if c.isClosed {
			return
		}

		c.connLock.RLock()
		conn := c.conn
		c.connLock.RUnlock()

		if conn == nil || conn.IsClosed() {
			c.logger.Warn("MQ Connection lost, reconnecting...")
			// Retry logic
			for i := 0; i < c.cfg.RetryCount; i++ {
				if err := c.connect(); err == nil {
					c.logger.Info("MQ Reconnected")
					break
				}
				c.logger.Error("MQ Reconnect failed", zap.Int("attempt", i+1))
				time.Sleep(time.Duration(c.cfg.Timeout) * time.Second)
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func (c *MQClient) Close() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.isClosed = true
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	close(c.done)
}

func (c *MQClient) CreateExchange(name, kind string, durable, autoDelete bool) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if c.channel == nil {
		return fmt.Errorf("channel is nil")
	}
	return c.channel.ExchangeDeclare(
		name,
		kind,
		durable,
		autoDelete,
		false, // internal
		false, // noWait
		nil,   // args
	)
}

func (c *MQClient) CreateQueue(name string, durable, autoDelete bool) (amqp.Queue, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if c.channel == nil {
		return amqp.Queue{}, fmt.Errorf("channel is nil")
	}
	return c.channel.QueueDeclare(
		name,
		durable,
		autoDelete,
		false, // exclusive
		false, // noWait
		nil,   // args
	)
}

func (c *MQClient) BindQueue(queueName, routingKey, exchange string) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if c.channel == nil {
		return fmt.Errorf("channel is nil")
	}
	return c.channel.QueueBind(
		queueName,
		routingKey,
		exchange,
		false,
		nil,
	)
}

func (c *MQClient) Publish(exchange, routingKey string, body []byte) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if c.channel == nil {
		return fmt.Errorf("channel is nil")
	}

	return c.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}

func (c *MQClient) Consume(queue, consumer string) (<-chan amqp.Delivery, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if c.channel == nil {
		return nil, fmt.Errorf("channel is nil")
	}

	if err := c.channel.Qos(
		c.cfg.PrefetchCount, // prefetch count
		0,                   // prefetch size
		false,               // global
	); err != nil {
		return nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	return c.channel.Consume(
		queue,
		consumer,
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}
