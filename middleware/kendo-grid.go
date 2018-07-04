package middleware

import (
	"github.com/arifsetiawan/go-common/query"
	"github.com/labstack/echo"
)

// KendoGrid middleware transform Kendo grid POST body into paging, sort and filter query
func KendoGrid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		gridParams := new(query.GridParams)
		err := c.Bind(gridParams)
		if err != nil {
			return err
		}

		if len(gridParams.Sort) > 0 {
			gridParams.HasSort = true
		}

		if len(gridParams.Filter.Filters) > 0 {
			gridParams.HasFilter = true

			for i, v := range gridParams.Filter.Filters {
				if len(v.Filters) > 0 {
					gridParams.Filter.Filters[i].HasSubFilter = true
				}
			}

		}

		// default
		if gridParams.PageSize == 0 {
			gridParams.PageSize = 10
		}

		if gridParams.Page == 0 {
			gridParams.Page = 1
		}

		if gridParams.Take == 0 {
			gridParams.Take = 10
		}

		c.Set("gridParams", gridParams)
		return next(c)
	}
}
