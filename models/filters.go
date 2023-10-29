package models

import (
	"errors"
	"strings"
)

type Filter struct {
	Page     int
	PageSize int
	OrderBy  string
	Query    string
}

type MetaData struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	NextPage     int
	PrevPage     int
	LastPage     int
	TotalRecords int
}

// Validate validates the amount of pages and page size
func (f *Filter) Validate() error {
	if f.Page <= 0 || f.Page >= 10_000_000 {
		return errors.New("invalid page range: 1 to 10 million")
	}

	if f.PageSize <= 0 || f.PageSize < 100 {
		return errors.New("invalid page size: 1 to 100 max")
	}
	return nil
}

// addOrdering constructs the order by clause
func (f *Filter) addOrdering(q string) string {
	if f.OrderBy == "popular" {
		return strings.Replace(q, "#orderby#", "ORDER BY votes desc, p.created_at desc", 1)
	}
	return strings.Replace(q, "#orderby#", "ORDER BY p.created_at desc", 1)
}

// addWhere construct the where clause
func (f *Filter) addWhere(q string) string {
	if len(f.Query) > 0 {
		return strings.Replace(q, "#where#", "WHERE LOWER(p.title) LIKE $1", 1)
	}
	return strings.Replace(q, "#where#", "", 1)
}

// addWhere construct the where clause
func (f *Filter) addLimitOffset(q string) string {
	if len(f.Query) > 0 {
		return strings.Replace(q, "#where#", "LIMIT $2 OFFSET $3", 1)
	}
	return strings.Replace(q, "#where#", "LIMIT $1 OFFSET $2", 1)
}

// applyTemplate constructs the queries writen above
func (f *Filter) applyTemplate(q string) string {
	return f.addLimitOffset(f.addWhere(f.addOrdering(q)))
}
