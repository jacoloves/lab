package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listHolidayRe = regexp.MustCompile(`^\/holiday[\/]*$`)
	getHolidayRe  = regexp.MustCompile(`^\/holiday\/year\/(\d+)$`)
)

type Data []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

type datastore struct {
	m map[string]Data
	*sync.RWMutex
}

type dateHandler struct {
	store *datastore
}

type I interface {
	convData(title string, date string)
}

func (h *dateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listHolidayRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getHolidayRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *dateHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	holiday := make([]Data, 0, len(h.store.m))
	for _, v := range h.store.m {
		holiday = append(holiday, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.MarshalIndent(holiday, " ", " ")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *dateHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getHolidayRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	y, ok := h.store.m[matches[1]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("year not found"))
		return
	}
	jsonBytes, err := json.MarshalIndent(y, " ", " ")
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
	h1 := holiday_2021()
	h2 := holiday_2022()
	h3 := holiday_2023()

	mux := http.NewServeMux()
	dHandler := &dateHandler{
		store: &datastore{
			m: map[string]Data{
				"2021": Data(h1),
				"2022": Data(h2),
				"2023": Data(h3),
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/holiday", dHandler)
	mux.Handle("/holiday/", dHandler)

	http.ListenAndServe("localhost:8080", mux)
}
