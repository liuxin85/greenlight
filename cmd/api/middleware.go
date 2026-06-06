package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate" // New import
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {
	// allows an average of 2 requests per second,
	// with a maximum of 4 requests in a single "burst".
	limiter := rate.NewLimiter(2, 4)

	//  The function we are returning is a closure, which 'closes over' the limiter
	// variable.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			app.rateLimitExeededResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimitExeededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}
