package item

import (
	"encoding/json"
	"github.com/dbielecki97/bookstore-items-api/client/es"
	"github.com/dbielecki97/bookstore-items-api/domain/query"
	"github.com/dbielecki97/bookstore-utils-go/errs"
)

const (
	indexItems = "items"
)

func (d *Dto) Save() errs.RestErr {
	result, err := es.Client.Index(indexItems, d)
	if err != nil {
		return err
	}

	d.ID = result.Id
	return nil
}

func (d *Dto) Get() errs.RestErr {
	result, restErr := es.Client.Get(indexItems, d.ID)
	if restErr != nil {
		return restErr
	}

	err := json.Unmarshal(result.Source, d)
	if err != nil {
		return errs.NewInternalServerErr("could not get item from elastic", errs.NewError("database error"))
	}

	return nil
}

func (d *Dto) Search(q query.EsQuery) ([]Dto, errs.RestErr) {
	searchResult, restErr := es.Client.Search(indexItems, q.Build())
	if restErr != nil {
		return nil, restErr
	}

	results := make([]Dto, 0)
	for _, h := range searchResult.Hits.Hits {
		var res Dto
		err := json.Unmarshal(h.Source, &res)
		if err != nil {
			return nil, errs.NewInternalServerErr("could not get item from elastic", errs.NewError("database error"))
		}
		res.ID = h.Id
		results = append(results, res)
	}

	return results, nil
}

func (d *Dto) Update() errs.RestErr {
	_, restErr := es.Client.Update(indexItems, d.ID, d)
	if restErr != nil {
		return restErr
	}

	return nil
}

func (d *Dto) Delete() errs.RestErr {
	return es.Client.Delete(indexItems, d.ID)
}
