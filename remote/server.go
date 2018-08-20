package remote

import (
	"io/ioutil"
	slog "log"
	"net"
	"os"
	"time"

	"github.com/aergoio/aergo-actor/actor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	s         *grpc.Server
	edpReader *endpointReader
)

// Start the remote server
func Start(address string, options ...RemotingOption) {
	grpclog.SetLogger(slog.New(ioutil.Discard, "", 0))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		plog.Error().Err(err).Msg("failed to listen")
		os.Exit(1)
	}
	config := defaultRemoteConfig()
	for _, option := range options {
		option(config)
	}

	address = lis.Addr().String()
	actor.ProcessRegistry.RegisterAddressResolver(remoteHandler)
	actor.ProcessRegistry.Address = address

	spawnActivatorActor()
	startEndpointManager(config)

	s = grpc.NewServer(config.serverOptions...)
	edpReader = &endpointReader{}
	RegisterRemotingServer(s, edpReader)
	plog.Info().Str("address", address).Msg("Starting Proto.Actor server")

	go s.Serve(lis)
}

func Shutdown(graceful bool) {
	if graceful {
		edpReader.suspend(true)
		stopEndpointManager()
		stopActivatorActor()

		//For some reason GRPC doesn't want to stop
		//Setup timeout as walkaround but need to figure out in the future.
		//TODO: grpc not stopping
		c := make(chan bool, 1)
		go func() {
			s.GracefulStop()
			c <- true
		}()

		select {
		case <-c:
			plog.Info().Msg("Stopped Proto.Actor server")
		case <-time.After(time.Second * 10):
			s.Stop()
			plog.Info().Str("err", "timeout").Msg("Stopped Proto.Actor server")
		}
	} else {
		s.Stop()
		plog.Info().Msg("Killed Proto.Actor server")
	}
}
