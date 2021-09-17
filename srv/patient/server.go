package patient

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/blendle/zapdriver"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogSettable "github.com/grpc-ecosystem/go-grpc-middleware/logging/settable"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCTXTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpenTracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	"http2/srv/patient/handler"
	"http2/srv/patient/proto"
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

func setupLogger(cmd *cobra.Command) *zap.Logger {
	setupFunc := zapdriver.NewProduction
	if debug {
		setupFunc = zapdriver.NewDevelopment
	}

	l, err := setupFunc()
	if err != nil {
		panic(fmt.Errorf("failed to initialise logger: %w", err))
	}
	return l.Named(cmd.Name())
}

func runE(cmd *cobra.Command, _ []string) error {
	log := setupLogger(cmd)
	defer func() {
		err := log.Sync()
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, fmt.Errorf("failed to sync logger: %w", err))
		}
	}()

	log.Info("setting up gRPC server")
	srv := grpc.NewServer()

	grpc.NewServer(
		grpc.StreamInterceptor(
			grpcMiddleware.ChainStreamServer(
				grpcRecovery.StreamServerInterceptor(),
				grpcOpenTracing.StreamServerInterceptor(),
				// grpc_prometheus.StreamServerInterceptor,
				grpcZap.StreamServerInterceptor(log),
				// grpcAuth.StreamServerInterceptor(myAuthFunction),
			)),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcRecovery.UnaryServerInterceptor(),
				grpcCTXTags.UnaryServerInterceptor(),
				grpcOpenTracing.UnaryServerInterceptor(),
				// grpc_prometheus.UnaryServerInterceptor,
				grpcZap.UnaryServerInterceptor(log),
				// grpcAuth.UnaryServerInterceptor(myAuthFunction),
			)),
	)

	log.Debug("setting debugger")
	grpcZap.SetGrpcLoggerV2(grpcLogSettable.ReplaceGrpcLoggerV2(), log)

	log.Debug("creating handler")
	backingArr := []*proto.Patient{
		{Id: "0_test_patient", Name: "Joan W"},
		{Id: "1_test_patient", Name: "Jack L"},
		{
			Id:   "2_test_patient",
			Name: "Janice T",
			Prescriptions: []*proto.Prescription{
				{
					MedicineId:   "Ventolin Evohaler (Blue)",
					DoctorId:     "0_test_doctor",
					PrescribedAt: timestamppb.New(time.Now().UTC()),
					Dosage:       &proto.Dosage{Units: "mg", Quantity: 20},
				},
			},
		},
	}
	h := handler.New(
		log,
		backingArr,
	)

	log.Debug("registering handler to grpc")
	proto.RegisterPatientServiceServer(srv, h)

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
