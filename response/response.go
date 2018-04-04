package response

import (
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/arifsetiawan/go-common/env"
)

// Response JSONAPI object
type Response struct {
	Errors []Error     `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Links  interface{} `json:"links,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Total  int         `json:"total,omitempty"`
}

// Error object
type Error struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// Links object
type Links struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

// MakeErrorResponse is function to generate error response
func MakeErrorResponse(tenant string, title string, err error, status int) *Response {
	log.WithFields(log.Fields{
		"app":    env.Getenv("APP_NAME", "none"),
		"type":   "backend",
		"tenant": tenant,
	}).Errorln(err)

	r := new(Response)
	es := make([]Error, 1)
	es[0] = Error{Status: status, Title: title, Detail: err.Error()}
	r.Errors = es
	return r
}

// JSONError is
func JSONError(c echo.Context, status int, err error) error {
	tenant := c.Get("tenant").(string)

	log.WithFields(log.Fields{
		"app":    env.Getenv("APP_NAME", "none"),
		"type":   "backend",
		"tenant": tenant,
	}).Errorln(err)

	r := new(Response)
	es := make([]Error, 1)
	es[0] = Error{Status: status, Detail: err.Error()}
	r.Errors = es
	return c.JSON(status, r)
}

// JSONErrorSimple is
func JSONErrorSimple(c echo.Context, status int, err error) error {
	
	log.WithFields(log.Fields{
		"app":    env.Getenv("APP_NAME", "none"),
		"type":   "backend",
	}).Errorln(err)

	r := struct {
		Status  string `json:"message"`
		Message int         `json:"status"`
	}{
		err.Error(),
		status,
	}

	return c.JSON(status, r)
}

// JSON is
func JSON(c echo.Context, status int, data interface{}) error {
	r := new(Response)
	r.Data = data
	return c.JSON(status, r)
}

// JSONGrid is
func JSONGrid(c echo.Context, status int, data interface{}, length int, count int) error {
	r := struct {
		Data  interface{} `json:"data"`
		Total int         `json:"total"`
	}{
		make([]interface{}, 0),
		count,
	}

	if length > 0 {
		r.Data = data
	}

	return c.JSON(status, r)
}

// GenerateLinks is
func GenerateLinks(host string, totalPages int, currentPage int, pageSize int, filterType string) *Links {
	l := new(Links)
	filterStr := ""
	if len(filterType) > 0 {
		filterStr = "&filter%5Btype%5D=" + filterType
	}

	l.Self = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	l.First = host + "?page%5Bnumber%5D=1&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	l.Last = host + "?page%5Bnumber%5D=" + strconv.Itoa(totalPages) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	if currentPage < totalPages {
		l.Next = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage+1) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	}
	if currentPage > 1 {
		l.Prev = host + "?page%5Bnumber%5D=" + strconv.Itoa(currentPage-1) + "&page%5Bsize%5D=" + strconv.Itoa(pageSize) + filterStr
	}
	return l
}
