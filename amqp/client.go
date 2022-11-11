package amqp

import (
	"github.com/locngodn/gas-common/util"
	"github.com/streadway/amqp"
)

type amqpClient struct {
	amqpPublisherManager
	amqpSubscriptionManager
	connection *amqp.Connection
	logger     util.Logger
}

func (cli *amqpClient) Init(connection *amqp.Connection) error {
	cli.connection = connection

	cli.logger.Infof("initiate publisher manager")
	if err := cli.amqpPublisherManager.Init(connection, cli.timeout, cli.maxRetries, cli.logger); err != nil {
		cli.logger.Warnf("Fail initiate publisher manager %v", err)
		return err
	}

	cli.logger.Infof("initiate subscription manager")
	if err := cli.amqpSubscriptionManager.Init(connection, cli.ttl, cli.logger); err != nil {
		cli.logger.Warnf("Fail initiate subscription manager %v", err)
		return err
	}
	return nil
}

func (cli *amqpClient) Close() error {
	cli.logger.Infof("try close subscription manager")
	if err := cli.amqpSubscriptionManager.Close(); err != nil {
		cli.logger.Errorf("Failed to close subscription manager: %v\n", err)
	}

	cli.logger.Infof("try close publisher manager")
	if err := cli.amqpPublisherManager.Close(); err != nil {
		cli.logger.Errorf("Failed to close publisher manager: %v\n", err)
	}

	cli.logger.Infof("try close connection")
	return cli.connection.Close()
}
