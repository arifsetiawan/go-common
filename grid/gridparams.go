package grid

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
	Field    string `json:"field,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`

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
