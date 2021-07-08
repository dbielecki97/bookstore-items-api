package es

import (
	"context"
	"fmt"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClient = &defaultEsClient{}
)

type esClient interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
}

type defaultEsClient struct {
	c *elastic.Client
}

func Init() {
	l := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheck(false),
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

func (c defaultEsClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	result, err := c.c.Index().
		Type(docType).
		Index(index).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in elastic in index %s", index), err)
	}
	return result, err
}
