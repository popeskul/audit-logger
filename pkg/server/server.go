package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/popeskul/audit-logger/pkg/domain"
)

type Server struct {
	grpcSrv      *grpc.Server
	loggerServer domain.AuditServiceServer
}

func NewServer(auditServer domain.AuditServiceServer, interceptor grpc.UnaryServerInterceptor) *Server {
	return &Server{
		grpcSrv: grpc.NewServer(
			grpc.UnaryInterceptor(interceptor),
		),
		loggerServer: auditServer,
	}
}

func (s *Server) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.registerServices()

	go func() {
		log.Fatalln(s.grpcSrv.Serve(lis))
	}()

	log.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	s.grpcSrv.GracefulStop()

	return nil
}

func (s *Server) registerServices() {
	domain.RegisterAuditServiceServer(s.grpcSrv, s.loggerServer)
}

func LoggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	logMessage(ctx, info.FullMethod, time.Since(start), err)
	return resp, err
}

func logMessage(
	ctx context.Context,
	method string,
	latency time.Duration,
	err error,
) {
	var requestId string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Print("No metadata")
	} else {
		if len(md.Get("Request-Id")) != 0 {
			requestId = md.Get("Request-Id")[0]
		}
	}
	log.Printf("Method:%s, Duration:%s, Error:%v, Request-Id:%s",
		method,
		latency,
		err,
		requestId,
	)
}
