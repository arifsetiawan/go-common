package middleware

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/arifsetiawan/go-common/query"

	"github.com/labstack/echo"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile("\\[(.*?)\\]")
}

// ListQuery middleware get jsonapi defined paging, sort and filter query
func ListQuery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := c.QueryParams()

		pageNumberExist := false
		pageOffsetExist := false

		listParams := new(query.ListParams)
		listParams.Filter = make(map[string]string)

		for k := range queryParams {
			//fmt.Printf("key[%s] value[%s]\n", k, v)
			if k == "page[number]" {
				listParams.PageNumber, _ = strconv.Atoi(c.QueryParam("page[number]"))
				pageNumberExist = true
			}
			if k == "page[size]" {
				listParams.PageSize, _ = strconv.Atoi(c.QueryParam("page[size]"))
			}
			if k == "page[offset]" {
				listParams.PageOffset, _ = strconv.Atoi(c.QueryParam("page[offset]"))
				pageOffsetExist = true
			}
			if k == "page[limit]" {
				listParams.PageLimit, _ = strconv.Atoi(c.QueryParam("page[limit]"))
			}
			if strings.Contains(k, "filter") {
				match := re.FindStringSubmatch(k)
				if len(match) == 2 {
					listParams.Filter[match[1]] = c.QueryParam(k)
				}
			}
			if strings.Contains(k, "sort") {
				listParams.Sort = strings.Split(c.QueryParam(k), ",")
			}
		}

		// if we have page-size we convert it into offset-limit
		if pageNumberExist && !pageOffsetExist {
			// boundary check
			if listParams.PageNumber < 1 {
				listParams.PageNumber = 1
			}

			listParams.PageOffset = (listParams.PageNumber - 1) * listParams.PageSize
			listParams.PageLimit = listParams.PageSize
		}

		// if we have offset-limit we convert it into page-size
		if !pageNumberExist && pageOffsetExist {
			listParams.PageNumber = listParams.PageOffset/listParams.PageLimit + 1

			// boundary check
			if listParams.PageNumber < 1 {
				listParams.PageNumber = 1
			}

			listParams.PageSize = listParams.PageLimit
		}

		// default
		if listParams.PageSize == 0 {
			listParams.PageSize = 10
		}

		if listParams.PageLimit == 0 {
			listParams.PageLimit = 10
		}

		c.Set("listParams", listParams)
		c.Set("gridParams", query.NewGridParamsFromListParams(listParams))
		return next(c)
	}
}
