package blog

import "net/http"



type comments interface {
	ServerHTTP(w http.ResponseWriter, r *http.Request)
}

type tags interface {
	ServerHTTP(w http.ResponseWriter, r *http.Request)
}

