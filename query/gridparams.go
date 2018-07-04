package query

import "strings"

// GridQueryBuilder is
type GridQueryBuilder interface {
	FilterClause(f GridFilter) string
	FullQuery(g GridParams, preQuery string) string
	FilterQuery(g GridParams, preQuery string) string
	SortQuery(g GridParams) string
	SortPagingQuery(g GridParams) string
}

// GridParams is
type GridParams struct {
	Page      int         `json:"page"`
	Skip      int         `json:"skip"`
	Take      int         `json:"take"`
	PageSize  int         `json:"pageSize"`
	Sort      []GridSort  `json:"sort,omitempty"`
	Group     interface{} `json:"group,omitempty"`
	HasSort   bool        `json:"hasSort,omitempty"`
	HasFilter bool        `json:"hasFilter,omitempty"`
	Filter    struct {
		Logic   string       `json:"logic,omitempty"`
		Filters []GridFilter `json:"filters,omitempty"`
	} `json:"filter,omitempty"`
}

// GridSort is
type GridSort struct {
	Field string `json:"field"`
	Dir   string `json:"dir"`
}

// GridFilter is
type GridFilter struct {
	Field    string      `json:"field,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`

	// Filter can be nested
	HasSubFilter bool         `json:"hasSubFilter,omitempty"`
	Logic        string       `json:"logic,omitempty"`
	Filters      []GridFilter `json:"filters,omitempty"`
}

// ComparisonOperator is
type ComparisonOperator struct {
	Operator       string
	WildcardBefore bool
	WildcardAfter  bool
	Unary          bool
}

/*
// ListParams is
type ListParams struct {
	PageOffset int               `url:"page[offset]"`
	PageLimit  int               `url:"page[limit]"`
	PageNumber int               `url:"page[number]"`
	PageSize   int               `url:"page[size]"`
	Sort       []string          `url:"page[sort]"`
	Filter     map[string]string `url:"page[filter]"`
}
*/

// NewGridParamsFromListParams is
func NewGridParamsFromListParams(l *ListParams) *GridParams {
	g := &GridParams{}

	// paging
	g.Page = l.PageNumber
	g.PageSize = l.PageSize
	g.Skip = l.PageOffset
	g.Take = l.PageNumber

	// sort
	for _, v := range l.Sort {
		gs := &GridSort{}
		if strings.HasPrefix(v, "-") {
			gs.Dir = "desc"
			gs.Field = strings.TrimPrefix(v, "-")
		} else {
			gs.Dir = "asc"
			gs.Field = v
		}
		g.Sort = append(g.Sort, *gs)
	}

	if len(g.Sort) > 0 {
		g.HasSort = true
	}

	if len(l.Filter) > 0 {
		g.HasFilter = true
		g.Filter.Logic = "and"
	}

	// filter
	for k, v := range l.Filter {
		gf := &GridFilter{}

		vs := strings.Split(v, ",")
		if len(vs) == 1 {
			gf.Field = k
			gf.Operator = "eq"
			gf.Value = v
		} else {
			gf.HasSubFilter = true
			gf.Logic = "or"

			gfe := &GridFilter{}
			for _, ve := range vs {
				gfe.Field = k
				gfe.Operator = "eq"
				gfe.Value = ve
			}

			gf.Filters = append(gf.Filters, *gfe)
		}

		g.Filter.Filters = append(g.Filter.Filters, *gf)
	}

	return g
}
