package service

import (
	item2 "github.com/dbielecki97/bookstore-items-api/src/domain/item"
	query2 "github.com/dbielecki97/bookstore-items-api/src/domain/query"
)
import "github.com/dbielecki97/bookstore-utils-go/errs"

var (
	ItemService itemService = &defaultItemService{}
)

type itemService interface {
	Create(item2.Dto) (*item2.Dto, errs.RestErr)
	Get(string) (*item2.Dto, errs.RestErr)
	Search(query2.EsQuery) ([]item2.Dto, errs.RestErr)
	Update(item2.Dto) (*item2.Dto, errs.RestErr)
	Delete(string) errs.RestErr
}

type defaultItemService struct {
}

func (s defaultItemService) Create(dto item2.Dto) (*item2.Dto, errs.RestErr) {
	if err := dto.Save(); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Get(id string) (*item2.Dto, errs.RestErr) {
	dto := item2.Dto{ID: id}

	err := dto.Get()
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Search(q query2.EsQuery) ([]item2.Dto, errs.RestErr) {
	dao := item2.Dto{}
	return dao.Search(q)
}

func (s defaultItemService) Update(dto item2.Dto) (*item2.Dto, errs.RestErr) {
	err := dto.Update()
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Delete(id string) errs.RestErr {
	dao := item2.Dto{ID: id}

	return dao.Delete()
}
