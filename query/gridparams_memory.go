package query

// MemoryQueryBuilder is
type MemoryQueryBuilder struct {
}

// FilterClause is
func (p MemoryQueryBuilder) FilterClause(f GridFilter) (query string) {
	return
}

// FullQuery is
func (p MemoryQueryBuilder) FullQuery(g GridParams, preQuery string) (query string) {
	return
}

// FilterQuery is
func (p MemoryQueryBuilder) FilterQuery(g GridParams, preQuery string) (query string) {
	return
}

// SortQuery is
func (p MemoryQueryBuilder) SortQuery(g GridParams) (query string) {
	return
}

// SortPagingQuery is
func (p MemoryQueryBuilder) SortPagingQuery(g GridParams) (query string) {
	return
}
