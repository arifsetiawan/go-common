package grid

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

// PostgresQueryBuilder is
type PostgresQueryBuilder struct {
}

// FilterClause is
func (p PostgresQueryBuilder) FilterClause(l GridFilter) (string, []interface{}) {
	query := ""
	var params []interface{}

	if l.HasSubFilter {
		i := 0
		query += "( "
		for _, v := range l.Filters {
			subQuery, subParams := p.FilterClause(v)
			query += subQuery
			params = append(params, subParams)
			if i != len(l.Filters)-1 {
				query += strings.ToUpper(l.Logic) + " "
			}
			i++
		}
		query += ") "
		return query, params
	}

	query += l.Field + " "
	query += pgOperatorMap[l.Operator].Operator + " "
	if !pgOperatorMap[l.Operator].Unary {
		// add placeholder to where cluase
		query += "? "
		// add actual value to params
		value := ""
		if pgOperatorMap[l.Operator].WildcardBefore {
			value += "%"
		}
		if !pgOperatorMap[l.Operator].Unary {
			value += l.Value
		}
		if pgOperatorMap[l.Operator].WildcardAfter {
			value += "%"
		}
		params = append(params, value)
	}

	return query, params
}

// FullQuery is
func (p PostgresQueryBuilder) FullQuery(query *gorm.DB, l GridParams, preQuery string, preParams []interface{}) *gorm.DB {

	fmt.Println(l)

	// Build WHERE clause
	where := ""
	if len(preQuery) > 0 {
		where += preQuery + " "
		if l.HasFilter {
			where += "AND "
		}
	}

	i := 0
	for _, v := range l.Filter.Filters {
		subWhere, subParams := p.FilterClause(v)
		where += subWhere
		preParams = append(preParams, subParams...)
		fmt.Println("Where", where)
		if i != len(l.Filter.Filters)-1 {
			where += strings.ToUpper(l.Filter.Logic) + " "
			fmt.Println("l.Filter.Logic", l.Filter.Logic)
		}
		i++
	}
	query = query.Where(where, preParams...)

	// Build ORDER BY clause
	sort := ""
	for i, v := range l.Sort {
		if i > 0 {
			sort += ", "
		}

		sort += v.Field + " " + strings.ToUpper(v.Dir)

		if i == len(l.Sort)-1 {
			sort += " "
		}
	}
	query = query.Order(sort)

	query = query.Offset(l.Skip).Limit(l.PageSize)

	return query
}

// FilterQuery is
func (p PostgresQueryBuilder) FilterQuery(query *gorm.DB, l GridParams, preQuery string, preParams []interface{}) *gorm.DB {
	where := ""
	if len(preQuery) > 0 {
		where += preQuery + " "
		if l.HasFilter {
			where += "AND "
		}
	}

	i := 0
	for _, v := range l.Filter.Filters {
		subWhere, subParams := p.FilterClause(v)
		where += subWhere
		preParams = append(preParams, subParams...)
		if i != len(l.Filter.Filters)-1 {
			where += strings.ToUpper(l.Filter.Logic) + " "
		}
		i++
	}
	return query.Where(where, preParams...)
}

// SortQuery is
func (p PostgresQueryBuilder) SortQuery(query *gorm.DB, l GridParams) *gorm.DB {
	sort := ""
	for i, v := range l.Sort {
		if i > 0 {
			sort += ", "
		}

		sort += v.Field + " " + strings.ToUpper(v.Dir)

		if i == len(l.Sort)-1 {
			sort += " "
		}
	}
	return query.Order(sort)
}

// SortPagingQuery is
func (p PostgresQueryBuilder) SortPagingQuery(query *gorm.DB, l GridParams) *gorm.DB {
	sort := ""
	for i, v := range l.Sort {
		if i > 0 {
			sort += ", "
		}

		sort += v.Field + " " + strings.ToUpper(v.Dir)

		if i == len(l.Sort)-1 {
			sort += " "
		}
	}
	query = query.Order(sort)

	return query.Offset(l.Skip).Limit(l.PageSize)
}

func makeOperation(op string, field string) string {
	placeholder := ""
	if operator, ok := pgOperatorMap[op]; ok {
		op = operator.Operator
		if !operator.Unary {
			placeholder = " ?"
		}
	}
	return field + fmt.Sprintf(" %s%s", op, placeholder)
}

func makeOperand(op string, value string, array bool) string {
	if operator, ok := pgOperatorMap[op]; ok {
		if operator.Unary {
			return ""
		}
		if array {
			return fmt.Sprintf("{%s}", value)
		}
	}

	return value
}

func isArray(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}

var pgOperatorMap = map[string]ComparisonOperator{
	"eq": ComparisonOperator{
		Operator: "=",
		Unary:    false,
	},
	"neq": ComparisonOperator{
		Operator: "<>",
		Unary:    false,
	},
	"contains": ComparisonOperator{
		Operator:       "LIKE",
		WildcardBefore: true,
		WildcardAfter:  true,
		Unary:          false,
	},
	"doesnotcontain": ComparisonOperator{
		Operator:       "NOT LIKE",
		WildcardBefore: true,
		WildcardAfter:  true,
		Unary:          false,
	},
	"startswith": ComparisonOperator{
		Operator:      "LIKE",
		WildcardAfter: true,
		Unary:         false,
	},
	"endswith": ComparisonOperator{
		Operator:       "LIKE",
		Unary:          false,
		WildcardBefore: true,
	},
	"isnull": ComparisonOperator{
		Operator: "IS NULL",
		Unary:    true,
	},
	"isnotnull": ComparisonOperator{
		Operator: "IS NOT NULL",
		Unary:    true,
	},
	"isempty": ComparisonOperator{
		Operator: "<> ''",
		Unary:    true,
	},
	"isnotempty": ComparisonOperator{
		Operator: "= ''",
		Unary:    true,
	},
}
