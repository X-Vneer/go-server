package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/x-vneer/go-server/api"
	"github.com/x-vneer/go-server/internal/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NeedDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var LoginDetails *tools.LoginDetails
		LoginDetails = (*database).GetUserLoginDetails(username)

		if LoginDetails == nil || (token != (*LoginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
