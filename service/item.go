package service

import (
	"github.com/dbielecki97/bookstore-items-api/domain/item"
	"github.com/dbielecki97/bookstore-items-api/domain/query"
)
import "github.com/dbielecki97/bookstore-utils-go/errs"

var (
	ItemService itemService = &defaultItemService{}
)

type itemService interface {
	Create(item.Dto) (*item.Dto, errs.RestErr)
	Get(string) (*item.Dto, errs.RestErr)
	Search(query.EsQuery) ([]item.Dto, errs.RestErr)
	Update(item.Dto) (*item.Dto, errs.RestErr)
	Delete(string) errs.RestErr
}

type defaultItemService struct {
}

func (s defaultItemService) Create(dto item.Dto) (*item.Dto, errs.RestErr) {
	if err := dto.Save(); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Get(id string) (*item.Dto, errs.RestErr) {
	dto := item.Dto{ID: id}

	err := dto.Get()
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Search(q query.EsQuery) ([]item.Dto, errs.RestErr) {
	dao := item.Dto{}
	return dao.Search(q)
}

func (s defaultItemService) Update(dto item.Dto) (*item.Dto, errs.RestErr) {
	err := dto.Update()
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Delete(id string) errs.RestErr {
	dao := item.Dto{ID: id}

	return dao.Delete()
}
