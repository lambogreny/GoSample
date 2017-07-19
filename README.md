## GoAPI

## Example code
- [Main](#main)
- [Router](#router)
- [Middleware](#middleware)
  - [Authen](#authen-middleware)
  - [AllowRoles](#allowroles-middleware)
- MVC (Comming Zoon)

## Main
```go
package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/GoAPI/routes"
)

func main() {
	h := routes.Router()
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Println(err)
	}
}
```

## Router
```go
package routes

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/middleware"
)

// Router provider
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.Handle("/staff", middleware.Chain(
		middleware.Authen("Bearer ABCD1234"),
		middleware.AllowRoles("staff"),
	)(http.HandlerFunc(staffHandler)))

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Staff!"))
}
```

## Middleware
```go
package middleware

import "net/http"

// Middleware is the http middleware
type Middleware func(http.Handler) http.Handler

// Chain is the helper function for chain middlewares into one middleware
func Chain(hs ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}
```

## Authen Middleware
```go
package middleware

import (
	"log"
	"net/http"
)

// Authen middleware
func Authen(token string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Authen Pass")
			h.ServeHTTP(w, r)
		})
	}
}
```

## AllowRoles Middleware
```go
package middleware

import (
	"net/http"
)

// AllowRoles middleware
func AllowRoles(roles ...string) Middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
```