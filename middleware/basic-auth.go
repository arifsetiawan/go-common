package middleware

import (
	"os"

	auth "github.com/abbot/go-http-auth"
	"github.com/labstack/echo"
	"gitlab.com/nextid/common/model"
)

// BasicAuth is
func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		secrets := auth.HtpasswdFileProvider(os.Getenv("BASIC_HTPASSWD"))
		a := &auth.BasicAuth{Realm: os.Getenv("BASIC_REALM"), Secrets: secrets}

		user := a.CheckAuth(c.Request())
		if user == "" {
			c.Set("basicAuthd", false)
			return next(c)
		}

		userEngine := &model.User{
			Name: user,
			ID:   "user:" + user,
		}

		c.Set("basicAuthd", true)
		c.Set("me", userEngine)
		return next(c)
	}
}
