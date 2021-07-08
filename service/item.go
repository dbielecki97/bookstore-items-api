package service

import (
	"github.com/dbielecki97/bookstore-items-api/domain"
	"net/http"
)
import "github.com/dbielecki97/bookstore-utils-go/errs"

var (
	ItemService itemService = &defaultItemService{}
)

type itemService interface {
	Create(domain.Item) (*domain.Item, *errs.RestErr)
	Get(string) (*domain.Item, *errs.RestErr)
}

type defaultItemService struct {
}

func (s defaultItemService) Create(i domain.Item) (*domain.Item, *errs.RestErr) {
	return nil, errs.NewRestErr("implement me", http.StatusNotImplemented, "not_implemented", nil)
}

func (s defaultItemService) Get(id string) (*domain.Item, *errs.RestErr) {
	return nil, errs.NewRestErr("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
