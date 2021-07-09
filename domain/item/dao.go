package item

import (
	"errors"
	"github.com/dbielecki97/bookstore-items-api/client/es"
	"github.com/dbielecki97/bookstore-utils-go/errs"
)

const (
	indexItems = "items"
	TypeItems  = "item"
)

func (d *Dto) Save() errs.RestErr {
	result, err := es.Client.Index(indexItems, TypeItems, d)
	if err != nil {
		return errs.NewInternalServerErr("could not save in elastic", errors.New("database error"))
	}

	d.ID = result.Id
	return nil
}
