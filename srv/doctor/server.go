package doctor

import (
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogSettable "github.com/grpc-ecosystem/go-grpc-middleware/logging/settable"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCTXTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpenTracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"http2/lib/logger"
	"http2/srv/doctor/handler"
	"http2/srv/doctor/proto"
)

var (
	CMD = &cobra.Command{
		Use:   "patient",
		Short: "starts a patient service",
		RunE:  runE,
	}

	port  = 8091
	debug = false
)

func init() {
	CMD.PersistentFlags().IntVar(&port, "port", 8091, "the port to run the service on")
	CMD.PersistentFlags().BoolVar(&debug, "debug", false, "whether to run in debug mode")
}

func setupHandlerStorage() []*proto.Doctor {
	return []*proto.Doctor{
		{Id: "0_test_patient", Name: "Joan W"},
		{Id: "1_test_patient", Name: "Jack L"},
		{
			Id:   "2_test_patient",
			Name: "Janice T",
		},
	}
}

func runE(cmd *cobra.Command, _ []string) error {
	log := logger.Setup(debug, cmd.Name())
	defer log.Sync()

	// patientProto.NewPatientServiceClient(grpc.NewClientStream())

	log.Debug("creating handler")
	h := handler.New(
		setupHandlerStorage(),
	)

	log.Info("setting up gRPC server")
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		grpcRecovery.UnaryServerInterceptor(), // Automatic panic recovery
		grpcCTXTags.UnaryServerInterceptor(),
		grpcOpenTracing.UnaryServerInterceptor(), // Automated OpenTracing span support
		grpcValidator.UnaryServerInterceptor(),   // Validate where possible
		// TODO: grpc_prometheus.UnaryServerInterceptor,
		// TODO: grpcAuth.UnaryServerInterceptor(myAuthFunction),
	)))
	grpcZap.SetGrpcLoggerV2(grpcLogSettable.ReplaceGrpcLoggerV2(), log)

	log.Debug("registering handler to grpc")
	proto.RegisterDoctorServiceServer(srv, h)

	if debug {
		log.Debug("registering reflection handler")
		reflection.Register(srv)
	}

	log.Debug("setting up gRPC listener")
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to create tcp listener for grpc", zap.Error(err))
	}

	log.Info("service ready - listening...")
	return srv.Serve(grpcListener)
}
