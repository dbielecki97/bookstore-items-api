package es

import (
	"context"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClient = &defaultEsClient{}
)

type esClient interface {
	setClient(*elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type defaultEsClient struct {
	c *elastic.Client
}

func Init() {
	l := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
		elastic.SetErrorLog(l),
		elastic.SetInfoLog(l),
	)
	if err != nil {
		logger.Fatal("could not create elasticsearch client", err)
	}

	pingService := client.Ping("http://127.0.0.1:9200")
	_, _, err = pingService.Do(context.Background())
	if err != nil {
		logger.Fatal("could not ping elasticsearch", err)
	}

	Client.setClient(client)
}

func (c *defaultEsClient) setClient(client *elastic.Client) {
	c.c = client
}

func (c defaultEsClient) Index(i interface{}) (*elastic.IndexResponse, error) {
	return c.c.Index().BodyJson(i).Do(context.Background())
}
