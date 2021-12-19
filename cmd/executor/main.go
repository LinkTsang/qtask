package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "qtask/api/proto/v1"
	"qtask/data"
	"qtask/pkg/endpoint"
	"qtask/pkg/sd"
	"qtask/pkg/service"
	"qtask/pkg/transport"
)

type Args struct {
	consulHost   string
	consulPort   string
	grpcHost     string
	grpcPort     string
	grpcTls      bool
	grpcCertFile string
	grpcKeyFile  string
}

func main() {
	args := new(Args)
	flag.StringVar(&args.consulHost, "consul.host", "", "The host address of Consul")
	flag.StringVar(&args.consulPort, "consul.port", "8500", "The port of Consul")
	flag.StringVar(&args.grpcHost, "grpc.host", "localhost", "The gRPC server address")
	flag.StringVar(&args.grpcPort, "grpc.port", "50051", "The gRPC server port")
	flag.BoolVar(&args.grpcTls, "grpc.tls", false, "The gRPC server connection uses TLS if true, else plain TCP")
	flag.StringVar(&args.grpcCertFile, "grpc.cert_file", "", "The gRPC server TLS cert file")
	flag.StringVar(&args.grpcKeyFile, "grpc.key_file", "", "The gRPC server TLS key file")
	flag.Parse()

	logger := setupLogger()

	var g run.Group

	setupExecutorServer(&g, logger, args)

	waitForTermination(&g)

	_ = logger.Log("exit", g.Run())

}

func setupLogger() log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

// setup gRPC server and service discovering
func setupExecutorServer(g *run.Group, logger log.Logger, args *Args) {
	healthCheck := health.NewServer()

	executorService := service.NewExecutorService()
	endpoints := endpoint.MakeServerEndpoints(executorService, logger)
	grpcServer := transport.NewGRPCServer(endpoints, logger)

	register := sd.ConsulRegister(
		args.consulHost, args.consulPort,
		args.grpcHost, args.grpcPort,
		pb.Executor_ServiceDesc.ServiceName,
		logger)

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", args.grpcHost, args.grpcPort))
	if err != nil {
		_ = logger.Log("transport", "gRPC", "message", "failed to listen", "error", err)
		os.Exit(1)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "gRPC", "addr", args.grpcHost, "port", args.grpcPort)
		var opts []grpc.ServerOption
		if args.grpcTls {
			if args.grpcCertFile == "" {
				args.grpcCertFile = data.Path("x509/server_cert.pem")
			}
			if args.grpcKeyFile == "" {
				args.grpcKeyFile = data.Path("x509/server_key.pem")
			}
			tlsCredentials, err := credentials.NewServerTLSFromFile(args.grpcCertFile, args.grpcKeyFile)
			if err != nil {
				_ = logger.Log("message", "Failed to generate credentials", "error", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(tlsCredentials)}
		}
		grpcBaseServer := grpc.NewServer(opts...)
		healthpb.RegisterHealthServer(grpcBaseServer, healthCheck)
		pb.RegisterExecutorServer(grpcBaseServer, grpcServer)
		register.Register()
		return grpcBaseServer.Serve(listener)
	}, func(err error) {
		register.Deregister()
		_ = listener.Close()
	})

}

// wait for termination
func waitForTermination(g *run.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
