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
	doctorpb "http2/srv/doctor/proto"
	patientpb "http2/srv/patient/proto"
)

var (
	CMD = &cobra.Command{
		Use:   "doctor",
		Short: "starts a doctor service",
		RunE:  runE,
	}

	port        = 8092
	patientPort = 8091
	debug       = false
)

func init() {
	CMD.PersistentFlags().IntVar(&port, "port", 8092, "the port to run the service on")
	CMD.PersistentFlags().IntVar(&patientPort, "patient-port", 8091, "the port to connect to the patient service on")
	CMD.PersistentFlags().BoolVar(&debug, "debug", false, "whether to run in debug mode")
}

func setupHandlerStorage() []*doctorpb.Doctor {
	return []*doctorpb.Doctor{
		{Id: "0_test_doctor", Name: "Joan W"},
		{Id: "1_test_doctor", Name: "Jack L"},
		{
			Id:   "2_test_doctor",
			Name: "Janice T",
		},
	}
}

func runE(cmd *cobra.Command, _ []string) error {
	log := logger.Setup(debug, cmd.Name())
	defer log.Sync()

	patientConn, err := grpc.DialContext(
		cmd.Context(),
		fmt.Sprintf("localhost:%d", patientPort),
		grpc.WithChainUnaryInterceptor(
			grpcZap.UnaryClientInterceptor(log),
			grpcOpenTracing.UnaryClientInterceptor(),
		),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("failed to dial patient service", zap.Error(err))
	}
	patientSVC := patientpb.NewPatientServiceClient(patientConn)

	log.Debug("creating handler")
	h := handler.New(log, setupHandlerStorage(), patientSVC)

	log.Info("setting up grpc server")
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		grpcZap.UnaryServerInterceptor(log),
		grpcRecovery.UnaryServerInterceptor(),  // Automatic panic recovery
		grpcValidator.UnaryServerInterceptor(), // Validate where possible
		grpcCTXTags.UnaryServerInterceptor(),
		grpcOpenTracing.UnaryServerInterceptor(), // Automated OpenTracing span support
		// TODO: grpc_prometheus.UnaryServerInterceptor,
		// TODO: grpcAuth.UnaryServerInterceptor(myAuthFunction),
	)))
	grpcZap.SetGrpcLoggerV2(grpcLogSettable.ReplaceGrpcLoggerV2(), log)

	log.Debug("registering handler to grpc")
	doctorpb.RegisterDoctorServiceServer(srv, h)

	if debug {
		log.Debug("registering reflection handler")
		reflection.Register(srv)
	}

	log.Debug("setting up grpc listener")
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to create tcp listener for grpc", zap.Error(err))
	}

	log.Info("service ready - listening...")
	return srv.Serve(grpcListener)
}
