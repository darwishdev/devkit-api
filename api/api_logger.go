package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func GrpcLogger(isDev bool) connect.UnaryInterceptorFunc {

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			startTime := time.Now()
			result, err := next(ctx, req)
			duration := time.Since(startTime)
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			logger := log.Info()
			if isDev {
				if err != nil {
					logger = log.Error().Err(err)
				}

				logger.
					Str("Procedure", req.Spec().Procedure).
					Interface("request", req.Any()).
					Interface("response", result).
					Dur("duration", duration).
					Msg("received a gRPC request")

			}
			return result, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
