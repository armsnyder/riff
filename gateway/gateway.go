package gateway

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case "/artists/search":
		handleSearchArtists(w, r)
	case "/graph":
		handleGraph(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleSearchArtists(w http.ResponseWriter, r *http.Request) {
	q := strings.ToLower(r.URL.Query().Get("q"))
	if q == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("content-type", "application/json")
		if strings.HasPrefix("david bowie", q) || strings.HasPrefix("bowie", q) {
			_, _ = fmt.Fprint(w, `[{"name":"David Bowie","id":1}]`)
		} else {
			_, _ = fmt.Fprint(w, "[]")
		}
	}
}

func handleGraph(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("content-type", "application/json")
	_, _ = fmt.Fprint(w, `[`+
		`[{"name":"David Bowie","id":1},{"name":"Velvet Underground","id":2}],`+
		`[{"name":"David Bowie","id":1},{"name":"Bob Dylan","id":3}],`+
		`[{"name":"David Bowie","id":1},{"name":"The Who","id":4}],`+
		`[{"name":"David Bowie","id":1},{"name":"Little Richard","id":6}],`+
		`[{"name":"Velvet Underground","id":2},{"name":"Babatunde Olatunji","id":5}],`+
		`[{"name":"Bob Dylan","id":3},{"name":"Little Richard","id":6}],`+
		`[{"name":"The Who","id":4},{"name":"Ray Charles","id":7}],`+
		`[{"name":"The Who","id":4},{"name":"John Lee Hooker","id":8}]`+
		`]`)
}
