package pipes

import "net/http"

type Handler interface {
	CreatePipe(w http.ResponseWriter, r *http.Request)
	ListPipes(w http.ResponseWriter, r *http.Request)
	GetPipe(w http.ResponseWriter, r *http.Request)
	UpdatePipe(w http.ResponseWriter, r *http.Request)
	DeletePipe(w http.ResponseWriter, r *http.Request)
}

type Module interface {
}
