package linktree

import (
	"context"
	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	next Service
	logger log.Logger
}

func (l *loggingMiddleware) CreateLinkTree(ctx context.Context, link *LinkTree) error {
	l.logger.Log("method", "CreateLinkTree", "slug", link.Slug)
	return l.next.CreateLinkTree(ctx, link)
}

func (l *loggingMiddleware) UpdateLinkTree(ctx context.Context, id string, link *LinkTree) error {
	l.logger.Log("method", "UpdateLinkTree", "slug", id)
	return l.next.UpdateLinkTree(ctx, id, link)
}

func (l *loggingMiddleware) DeleteLinkTree(ctx context.Context, id string) error {
	l.logger.Log("method", "DeleteLinkTree", "slug", id)
	return l.next.DeleteLinkTree(ctx, id)
}

func (l *loggingMiddleware) GetLinkTree(ctx context.Context, id string) (LinkTree, error) {
	l.logger.Log("method", "GetLinkTree", "slug", id)
	return l.next.GetLinkTree(ctx, id)
}

func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &loggingMiddleware{
		next:   s,
		logger: logger,
	}
}
