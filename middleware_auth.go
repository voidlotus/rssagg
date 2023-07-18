package main

import (
	"fmt"
	"net/http"

	"github.com/voidlotus/rssagg/internal/auth"
	"github.com/voidlotus/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't create user %v", err))
			return
		}
		/* Context package in the standard library
		it basically gives you a way to track
		something that happening accros multiple go routine
		the most important things that you can do with context:
		is that you can cancel it, by canceling the context it would
		effectively kill the HTTP request
		*/
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
