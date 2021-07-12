package es

import (
	"context"
	"fmt"
	"github.com/dbielecki97/bookstore-utils-go/errs"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"github.com/olivere/elastic/v7"
)

var (
	Client esClient = &defaultEsClient{}
)

type esClient interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, errs.RestErr)
	Get(string, string) (*elastic.GetResult, errs.RestErr)
	Search(string, elastic.Query) (*elastic.SearchResult, errs.RestErr)
	Update(string, string, interface{}) (*elastic.UpdateResponse, errs.RestErr)
	Delete(string, string) errs.RestErr
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
	logger.Info("Successfully initialized ElasticSearch client.")
}

func (c *defaultEsClient) setClient(client *elastic.Client) {
	c.c = client
}

func (c defaultEsClient) Index(index string, doc interface{}) (*elastic.IndexResponse, errs.RestErr) {
	result, err := c.c.Index().
		Index(index).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in elastic in index %s", index), err)
		return nil, errs.NewInternalServerErr("could not save in elastic", errs.NewError("database error"))
	}
	return result, nil
}

func (c *defaultEsClient) Get(index, id string) (*elastic.GetResult, errs.RestErr) {
	result, err := c.c.Get().
		Index(index).
		Id(id).
		Do(context.Background())
	if err != nil {
		if elastic.IsNotFound(err) {
			return nil, errs.NewNotFoundErr(fmt.Sprintf("could not find item in elastic with id %s", id))
		}

		logger.Error(fmt.Sprintf("error when trying to get a document from index %s", index), err)
		return nil, errs.NewInternalServerErr("could not get item from elastic", errs.NewError("database error"))
	}

	return result, nil
}

func (c *defaultEsClient) Search(index string, query elastic.Query) (*elastic.SearchResult, errs.RestErr) {
	result, err := c.c.Search(index).
		Query(query).
		Do(context.Background())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, errs.NewInternalServerErr("error when trying to search documents", errs.NewError("database error"))
	}

	return result, nil
}

func (c *defaultEsClient) Update(index string, id string, doc interface{}) (*elastic.UpdateResponse, errs.RestErr) {
	result, err := c.c.Update().
		Index(index).
		Id(id).
		Doc(doc).
		Refresh("true").
		Do(context.Background())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to update document in index %s with id %s", index, id), err)
		return nil, errs.NewInternalServerErr("error when trying to update document", errs.NewError("database error"))
	}

	return result, nil
}

func (c *defaultEsClient) Delete(index string, id string) errs.RestErr {
	_, err := c.c.Delete().Index(index).Id(id).Do(context.Background())
	if err != nil {
		if elastic.IsNotFound(err) {
			return errs.NewNotFoundErr(fmt.Sprintf("not found document at index %s with id %s", index, id))
		}

		logger.Error(fmt.Sprintf("error when trying to delete document at index %s with id %s", index, id), err)
		return errs.NewInternalServerErr("error when trying to delete document", errs.NewError("database error"))
	}

	return nil
}
