package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listUserRe = regexp.MustCompile(`^\/holiday[\/]*$`)
	getUserRe  = regexp.MustCompile(`^\/holiday\/(\d+)$`)
)

// holiday represents our REST resource
type holiday struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

// our in-memory datastore
// remember to guard map access with a mutex for  concurrent access
type datastore struct {
	m map[string]holiday
	*sync.RWMutex
}

type holidayHandler struct {
	store *datastore
}

func (h *holidayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// all users request are going to be routed here
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *holidayHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	holiday := make([]holiday, 0, len(h.store.m))
	for _, v := range h.store.m {
		holiday = append(holiday, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(holiday)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *holidayHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getUserRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	mux := http.NewServeMux()
	holidayH := &holidayHandler{
		store: &datastore{
			m: map[string]holiday{
				"1": holiday{Title: "元旦", Date: "2022-01-01"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/holiday", holidayH)
	mux.Handle("/holiday/", holidayH)

	http.ListenAndServe("localhost:8080", mux)
}
