package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os/signal"
	"syscall"
	"time"

	linktreesvc "github.com/heat/linktree/linktree"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)
func main() {

	var(
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		mongoUrl = flag.String("database.url", "mongodb://admin:admin@127.0.0.1:27017", "MongoDB database URL")
		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}


	var mongoClient *mongo.Client
	{
		var err error
		mongoClient, err = mongo.NewClient(options.Client().ApplyURI(*mongoUrl))

		mctx, _ := context.WithTimeout(ctx, 2*time.Second)
		mongoClient.Connect(ctx)
		err = mongoClient.Ping(mctx, readpref.Primary())
		if err != nil {
			logger.Log("error", err.Error())
			panic("database is out!")
		}
	}

	var svc linktreesvc.Service
	{
		svc = linktreesvc.NewServiceMongo(mongoClient, log.With(logger, "component", "service"))
		svc = linktreesvc.NewLoggingMiddleware(log.With(logger, "component", "linktree"), svc)
	}
	var h http.Handler
	{
		h = linktreesvc.MakeHTTPHandler(svc, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <- c)
	}()
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
