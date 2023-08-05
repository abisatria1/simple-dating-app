package model

type PaginationQueryParams struct {
	Page    int64 `query:"page"`
	PerPage int64 `query:"per_page"`
}
