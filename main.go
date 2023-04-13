package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/tnaucoin/stringsvc/pkg/middleware"
	"github.com/tnaucoin/stringsvc/pkg/service"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = service.New()
	svc = middleware.LoggingMiddleware{Logger: logger, Next: svc}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	logger.Log("err", http.ListenAndServe(":8080", nil))

}
