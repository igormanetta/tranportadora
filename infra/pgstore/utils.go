package pgstore

import (
	"context"
	"fmt"
	"transportadora/models"
)

func Pagination(sql, filter string, pLimit, pPage int, q *Queries, ctx context.Context) (*models.PaginationResponse, error) {

	pagination := models.PaginationResponse{}
	recordcount := 0

	var limit = 10
	var page = 1
	if pLimit > 0 {
		limit = pLimit
	}
	if pPage > 0 {
		page = pPage
	}

	rows, err := q.db.Query(ctx, sql+filter)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		recordcount++
	}

	total := recordcount / limit

	remainder := recordcount % limit
	if remainder == 0 {
		pagination.TotalPage = total
	} else {
		pagination.TotalPage = total + 1
	}

	pagination.CurrentPage = page
	pagination.RecordPerPage = limit
	pagination.TotalRecord = recordcount

	return &pagination, nil
}

func MountLimitOffset(pLimit, pPage int) string {
	var limit = 10
	var page = 1
	if pLimit > 0 {
		limit = pLimit
	}
	if pPage > 0 {
		page = pPage
	}
	return fmt.Sprintf(" limit %d offset %d", limit, limit*(page-1))
}
