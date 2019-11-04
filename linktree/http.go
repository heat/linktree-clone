package linktree

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	contentTypeApplicationJson = "application/json; charset=utf-8"
)

var (
	errBadRequest = errors.New("invalid request")
)

type errorResponse struct {
	Error string `json:"error"`
	success int `json: "success"`
}

type refLinkRequest struct {
	ID string
}

type dataLinkRequest struct {
	ID string
	Link LinkTree
}
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	e := MakeServerEndpoints(s)
	r.Methods("POST").Path("/links").Handler(httptransport.NewServer(
		e.PostLinkEndpoint,
		decodePostDataLinkRequest,
		encodeResponse,
		options...
	))
	r.Methods("GET").Path("/links/{id}").Handler(httptransport.NewServer(
		e.GetLinkEndpoint,
		decodeRefLinkRequest,
		encodeResponse,
		options...
		))
	r.Methods("PUT").Path("/links/{id}").Handler(httptransport.NewServer(
		e.PutLinkEndpoint,
		decodeDataLinkRequest,
		encodeResponse,
		options...
		))

	r.Methods("DELETE").Path("/links/{id}").Handler(httptransport.NewServer(
		e.DeleteLinkEndpoint,
		decodeRefLinkRequest,
		encodeResponse,
		options...
	))
	return r
}

func decodePostDataLinkRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	var link LinkTree
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		return nil, err
	}

	return dataLinkRequest{
		ID: "",
		Link: link,
	}, nil
}

func decodeDataLinkRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRequest
	}

	var link LinkTree
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		return nil, err
	}

	return dataLinkRequest{
		ID: id,
		Link: link,
	}, nil
}

func decodeRefLinkRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		return nil, errBadRequest
	}
	return refLinkRequest{ ID: id}, nil
}

func encodeResponse( ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.Header().Set("Content-Type", contentTypeApplicationJson)
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("no error to encode")
	}
	w.Header().Set("Content-Type", contentTypeApplicationJson)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(errorResponse{
		Error:   err.Error(),
		success: 0,
	})
}