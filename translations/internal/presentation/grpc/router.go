package grpcrouter

import (
	grpchandlers "translations/internal/presentation/grpc/handlers"
	usecases "translations/internal/usecase"

	"google.golang.org/grpc"
	pb "words-app.local/protos/gen/translations"
)

type GRPCRouter struct {
	usecases *usecases.UseCases
}

func (r *GRPCRouter) Register(s *grpc.Server) {
	pb.RegisterTranslationsServer(s, grpchandlers.New(r.usecases.Translations.Translate))
}

func NewGRPCRouter(uc *usecases.UseCases) *GRPCRouter {
	return &GRPCRouter{
		usecases: uc,
	}
}
