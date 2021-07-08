package service

import (
	"github.com/dbielecki97/bookstore-items-api/domain/item"
	"net/http"
)
import "github.com/dbielecki97/bookstore-utils-go/errs"

var (
	ItemService itemService = &defaultItemService{}
)

type itemService interface {
	Create(item.Dto) (*item.Dto, *errs.RestErr)
	Get(string) (*item.Dto, *errs.RestErr)
}

type defaultItemService struct {
}

func (s defaultItemService) Create(dto item.Dto) (*item.Dto, *errs.RestErr) {
	if err := dto.Save(); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s defaultItemService) Get(id string) (*item.Dto, *errs.RestErr) {
	return nil, errs.NewRestErr("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
