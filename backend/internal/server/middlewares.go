package server

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"socialNetwork/entity"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from your React app
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		// Allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// Note: This is split across multiple lines for readability. You don't
	// 	// need to do this in your own code.
	// 	w.Header().Set("Content-Security-Policy",
	// 		"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
	// 	w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
	// 	w.Header().Set("X-Content-Type-Options", "nosniff")
	// 	w.Header().Set("X-Frame-Options", "deny")
	// 	w.Header().Set("X-XSS-Protection", "0")
	// 	next.ServeHTTP(w, r)
	// })
}

func (app *App) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Loger.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *App) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				//  Avoid writing if connection is hijacked
				if _, ok := w.(http.Hijacker); ok {
					app.Loger.Error.Println("RecoverPanic: connection hijacked, skipping WriteHeader")
					return
				}
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				// Call the app.serverError helper method to return a 500
				// Internal Server response.
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *App) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user is not authenticated, redirect them to the login page and
		// return from the middleware chain so that no subsequent handlers in
		// the chain are executed.
		if !app.isAuthenticated(r) {
			w.WriteHeader(http.StatusForbidden)
			//			http.Redirect(w, r, "api/login", http.StatusSeeOther)
			return
		}

		// Otherwise set the "Cache-Control: no-store" header so that pages
		// require authentication are not stored in the users browser cache (or
		// other intermediary cache).
		w.Header().Add("Cache-Control", "no-store")
		// And call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func (app *App) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the authenticatedUserID value from the session using the
		// GetInt() method. This will return the zero value for an int (0) if no
		// "authenticatedUserID" value is in the session -- in which case we
		// call the next handler in the chain as normal and return.
		// id := app.SessionManager.GetInt(r.Context(), "authenticatedUserID")
		id := app.SessionManager.GetInt(r.Context(), "authenticatedUserID")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}
		// Otherwise, we check to see if a user with that ID exists in our
		// database.
		exists, err := app.User.IsUserExist(uint(id))
		if err != nil {
			app.serverError(w, err)
			return
		}
		// If a matching user is found, we know we know that the request is
		// coming from an authenticated user who exists in our database. We
		// create a new copy of the request (with an isAuthenticatedContextKey
		// value of true in the request context) and assign it to r.
		if exists {
			ctx := context.WithValue(r.Context(), entity.IsAuthenticatedContextKey, true)
			ctx = context.WithValue(ctx, entity.ContextID, id)
			r = r.WithContext(ctx)
		}
		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}

func (app *App) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(entity.IsAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *App) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.Loger.Error.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
