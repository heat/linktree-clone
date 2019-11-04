package linktree

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service interface {

	CreateLinkTree(ctx context.Context, link *LinkTree) error
	UpdateLinkTree(ctx context.Context, id string, link *LinkTree) error
	DeleteLinkTree(ctx context.Context, id string) error
	GetLinkTree(ctx context.Context, id string) (LinkTree, error)
}

type LinkTree struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Email string `json:"email"`
	Links []LinkLink `json:"links"`
}

type LinkLink struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	URL string `json:"url"`
}

type mongoService struct {
	client *mongo.Client
	logger log.Logger
}

func (m *mongoService) CreateLinkTree(ctx context.Context, link *LinkTree) error {
	collection := m.client.Database("linktree").Collection("links")
	ctx, _ = context.WithTimeout(ctx, 2 * time.Second)

	// try to find before insert
	filter := bson.D{{"slug", link.Slug}}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("slug already exists!")
	}
	var data LinkTree
	data = *link
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	m.logger.Log("InsertedID", res.InsertedID)
	return nil
}

func (m *mongoService) UpdateLinkTree(ctx context.Context, id string, link *LinkTree) error {
	collection := m.client.Database("linktree").Collection("links")
	ctx, _ = context.WithTimeout(ctx, 2 * time.Second)
	var data LinkTree
	data = *link
	filter := bson.D{{"slug", id}}
	res, err := collection.ReplaceOne(ctx, filter, data)
	if err != nil {
		return err
	}
	m.logger.Log("UpsertedID", res.UpsertedID)
	return nil
}

func (m *mongoService) DeleteLinkTree(ctx context.Context, id string) error {
	collection := m.client.Database("linktree").Collection("links")
	ctx, _ = context.WithTimeout(ctx, 2 * time.Second)
	filter := bson.D{{"slug", id}}
	res, err := collection.DeleteOne(ctx, filter)

	m.logger.Log("DeletedCount", res.DeletedCount)
	return err
}

func (m *mongoService) GetLinkTree(ctx context.Context, id string) (link LinkTree, err error) {
	collection := m.client.Database("linktree").Collection("links")
	ctx, _ = context.WithTimeout(ctx, 2 * time.Second)
	filter := bson.D{{"slug", id}}
	err = collection.FindOne(ctx, filter).Decode(&link)

	return link, err
}

func NewServiceMongo(client *mongo.Client, logger log.Logger) Service {
	return &mongoService{client, logger }
}
