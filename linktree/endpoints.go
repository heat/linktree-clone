package linktree

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PostLinkEndpoint endpoint.Endpoint
	GetLinkEndpoint endpoint.Endpoint
	PutLinkEndpoint endpoint.Endpoint
	DeleteLinkEndpoint endpoint.Endpoint
}

type refResponse struct {
	ID string
}

type dataResponse struct {
	ID string
	Link LinkTree
}

func MakeServerEndpoints(s Service) Endpoints {

	return Endpoints{
		PostLinkEndpoint:   makePostLinkEndpoint(s),
		GetLinkEndpoint:    makeGetLinkEndpoint(s),
		PutLinkEndpoint:    makePutLinkEndpoint(s),
		DeleteLinkEndpoint: makeDeleteLinkEndpoint(s),
	}
}

func makeDeleteLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		linkRequest := request.(refLinkRequest)
		err := s.DeleteLinkTree(ctx, linkRequest.ID)
		if err != nil {
			return nil, err
		}
		return refResponse{ linkRequest.ID }, nil
	}
}

func makePutLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		linkRequest := request.(dataLinkRequest)
		// id as slug
		linkRequest.Link.Slug = linkRequest.ID
		err := s.UpdateLinkTree(ctx, linkRequest.ID, &linkRequest.Link)
		if err != nil {
			return nil, err
		}

		return refResponse{ linkRequest.ID }, nil
	}
}

func makeGetLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		linkRequest := request.(refLinkRequest)
		link, err := s.GetLinkTree(ctx, linkRequest.ID)
		if err != nil {
			return nil, err
		}

		return dataResponse{ linkRequest.ID, link }, nil
	}
}

func makePostLinkEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		linkRequest := request.(dataLinkRequest)
		err := s.CreateLinkTree(ctx, &linkRequest.Link)
		if err != nil {
			return nil, err
		}

		return refResponse{ linkRequest.Link.Slug }, nil
	}
}


