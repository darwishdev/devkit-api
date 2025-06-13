package api

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/devkit-api/config"
	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/pkg/auth"
	"github.com/darwishdev/devkit-api/pkg/redisclient"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1/devkitv1connect"
	"github.com/rs/zerolog/log"
)

var (
	testConfig      config.Config
	store           db.Store
	redisClient     redisclient.RedisClientInterface
	validator       *protovalidate.Validator
	isRedisDisabled bool
	testServer      *http.Server
	testClient      devkitv1connect.DevkitServiceClient
	tokenMaker      auth.Maker
)

func startRealServer(store db.Store) *http.Server {
	server, err := NewServer(testConfig, store, tokenMaker, redisClient, validator)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the API")
	}
	httpServer := server.NewGrpcHttpServer()
	go func() {
		err = httpServer.ListenAndServe()
		if err != nil && !strings.Contains(err.Error(), "Server closed") {
			log.Fatal().Err(err).Msg("cannot start the API")
		}
	}()
	return httpServer

}
func setupTestServer(store db.Store) (devkitv1connect.DevkitServiceClient, *http.Server) {
	// Create a test HTTP server with the API handler
	serv := startRealServer(store)
	// Initialize a client that will connect to the test server
	testClient := devkitv1connect.NewDevkitServiceClient(
		http.DefaultClient,
		"http://localhost:9090",
		connect.WithGRPC(),
	)
	return testClient, serv
}

func TestMain(m *testing.M) {
	var err error
	_, err = config.LoadState("../config")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config state")
	}

	testConfig, err = config.LoadConfig("../config", "test")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load test config")
	}
	testConfig.State = "TEST"
	redisClient = redisclient.NewRedisClient(
		testConfig.RedisHost,
		testConfig.RedisPort,
		testConfig.RedisPassword,
		testConfig.RedisDatabase,
		isRedisDisabled,
	)
	validator, err = protovalidate.New()
	if err != nil {
		log.Fatal().Err(err).Msg("can't get the validator")
	}

	tokenMaker, err = auth.NewPasetoMaker(testConfig.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Paseto maker")
	}

	store, _, err = db.InitDB(context.Background(), testConfig.DBSource, false)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the database")
	}

	testClient, testServer = setupTestServer(store)
	// Set up the test server and client
	// Run tests
	code := m.Run()
	// Clean up after tests
	defer testServer.Close()
	os.Exit(code)
}
