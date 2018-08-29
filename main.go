package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Gopher represents the furry creature we all love
type Gopher struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	FavoriteColor string `json:"favoritecolour"`
}

// GopherData Holds A Slice of Gophers.
type GopherData struct {
	Gophers []Gopher `json:"gophers"`
}

var (
	gophers GopherData
)

func init() {

	gophers = GopherData{
		[]Gopher{},
	}
}

func main() {

	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a list of all the Gophers in Memory

		var err error
		jsonOutput := []byte("")

		if len(gophers.Gophers) > 0 {
			jsonOutput, err = json.Marshal(gophers)
			if err != nil {
				errJSON, _ := json.Marshal(struct{ err string }{err.Error()})
				http.Error(w, string(errJSON), http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Add("Content-type", "application/json")
		w.Write(jsonOutput)
	})

	r.Methods(http.MethodPost).Path("/gopher").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotImplemented)
		return
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
