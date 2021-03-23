package main

import (
	"flag"
	"github.com/marcusthierfelder/mpi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var logLevel = flag.String("log-level", "info", "logging level")

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stdout,
		PartsOrder: []string{
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		}})
	l, err := zerolog.ParseLevel(*logLevel)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")

		return
	}
	zerolog.SetGlobalLevel(l)

	mpi.Init()

	log.Info().Msgf(
		"comm size %v comm rank %v",
		mpi.Comm_size(mpi.COMM_WORLD), mpi.Comm_rank(mpi.COMM_WORLD))

	mpi.Finalize()
}
