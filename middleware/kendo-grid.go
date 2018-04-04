package middleware

import (
	"github.com/labstack/echo"
	"github.com/arifsetiawan/go-common/grid"
)

// KendoGrid middleware transform Kendo grid POST body into paging, sort and filter query
func KendoGrid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		gridParams := new(grid.GridParams)
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

		c.Set("gridParams", gridParams)
		return next(c)
	}
}
