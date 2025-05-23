package api

import (
	"net/http"

	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/devkit-api/config"
	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/pkg/auth"
	"github.com/darwishdev/devkit-api/pkg/redisclient"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1/devkitv1connect"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	config      config.Config
	store       db.Store
	tokenMaker  auth.Maker
	redisClient redisclient.RedisClientInterface
	validator   *protovalidate.Validator
	api         devkitv1connect.DevkitServiceHandler
}

func NewServer(config config.Config, store db.Store, tokenMaker auth.Maker, redisClient redisclient.RedisClientInterface, validator *protovalidate.Validator) (*Server, error) {
	api, err := NewApi(config, store, tokenMaker, redisClient, validator)

	if err != nil {
		return nil, err
	}
	return &Server{
		config:      config,
		validator:   validator,
		redisClient: redisClient,
		tokenMaker:  tokenMaker,
		store:       store,
		api:         api,
	}, nil
}

const maxMessageSize = 10 * 1024 * 1024 // 10MB

func (s Server) NewGrpcHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler("https://darwishdev.com", http.StatusFound))
	// here we can find examples of diffrent compression method 	https://connectrpc.com/docs/go/serialization-and-compression/#compression
	compress1KB := connect.WithCompressMinBytes(1024)
	log.Debug().Interface("stete", s.config.State).Msg("state ios")
	interceptors := connect.WithInterceptors(s.NewValidateInterceptor(), s.NewAuthenticationInterceptor(), s.NewAuthorizationInterceptor(), s.NewLoggerInterceptor())

	mux.Handle(devkitv1connect.NewDevkitServiceHandler(
		s.api,
		interceptors,
		compress1KB,
		connect.WithReadMaxBytes(maxMessageSize),
		connect.WithSendMaxBytes(maxMessageSize),
	))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	// Add upload handler here
	mux.HandleFunc("/upload", s.api.(*Api).HandleHttpFileUpload)
	cors := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
	server := &http.Server{
		Addr:    s.config.GRPCServerAddress,
		Handler: h2c.NewHandler(cors.Handler(mux), &http2.Server{}),
	}
	return server

}
