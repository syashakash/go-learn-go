package context

import (
	"context"
	"net/http"
)

type Store interface {
	Fetch(context.Context) (string,error)
}


func Server(store Store) (http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
