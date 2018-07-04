package query

// ListParams is
type ListParams struct {
	PageOffset int               `url:"page[offset]"`
	PageLimit  int               `url:"page[limit]"`
	PageNumber int               `url:"page[number]"`
	PageSize   int               `url:"page[size]"`
	Sort       []string          `url:"page[sort]"`
	Filter     map[string]string `url:"page[filter]"`
}
