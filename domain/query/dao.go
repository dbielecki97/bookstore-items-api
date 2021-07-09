package query

import "github.com/olivere/elastic/v7"

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()

	eqQueries := make([]elastic.Query, 0)
	for _, eqQuery := range q.Equals {
		eqQueries = append(eqQueries, elastic.NewMatchQuery(eqQuery.Field, eqQuery.Value))
	}
	query.Must(eqQueries...)
	return query
}
