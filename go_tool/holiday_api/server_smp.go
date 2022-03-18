package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listHolidayRe = regexp.MustCompile(`^\/holiday[\/]*$`)
	getHolidayRe  = regexp.MustCompile(`^\/holiday\/?year=(\d+)$`)
)

type Data []struct {
	Title string `json:"Title"`
	Date  string `json:"Date"`
}

type datastore struct {
	m map[string]string
	*sync.RWMutex
}

type HolidayHandler struct {
	store *datastore
}

func (h *HolidayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listHolidayRe.MatchString(r.URL.Path):
		return
	case r.Method == http.MethodGet && getHolidayRe.MatchString(r.URL.Path):
		return
	default:
		return
	}
	/*
		_, err := fmt.Fprint(w, len(h.store.m))
		if err != nil {
			return
		}
		for _, v := range h.store.m {
			_, err := fmt.Fprint(w, v+",")
			if err != nil {
				return
			}
		}
	*/
	/*
			_, err := fmt.Fprint(w, h.store.m)
		if err != nil {
			return
		}
	*/
}

func main() {
	h1 := holiday_2021()

	s1, err := json.MarshalIndent(h1, " ", " ")
	if err != nil {
		panic(err)
	}

	h2 := holiday_2022()

	s2, err := json.MarshalIndent(h2, " ", " ")
	if err != nil {
		panic(err)
	}

	h3 := holiday_2023()

	s3, err := json.MarshalIndent(h3, " ", " ")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	hHandler := &HolidayHandler{
		store: &datastore{
			m: map[string]string{
				"2021": string(s1),
				"2022": string(s2),
				"2023": string(s3),
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/holiday", hHandler)

	http.ListenAndServe("localhost:8080", mux)
}
