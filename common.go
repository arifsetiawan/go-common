package common

import (
	"errors"
)

var (
	ErrorAuthenticationFailure   = errors.New("username or password didn't match")
	ErrorNoSuchUser              = errors.New("no such user")
	ErrorBadRequest              = errors.New("bad request")
	ErrorBadChallengeToken       = errors.New("bad challenge token format")
	ErrorNoChallengeCookie       = errors.New("challenge token isn't stored in a cookie")
	ErrorTooMuchChallengeCookies = errors.New("too much challenge cookies, don't know which to use")
	ErrorDatabaseProvider        = errors.New("can't connect to provider")
	ErrorUserSuspended           = errors.New("user is suspended")
)
