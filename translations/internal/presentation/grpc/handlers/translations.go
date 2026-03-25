package grpchandlers

import (
	"context"
	usecases "translations/internal/usecase/translations"

	pb "words-app.local/protos/gen/translations"
)

type handler struct {
	uc *usecases.TranslationsUseCase
}

func NewTranslationsHandler() *handler {
	return &handler{
		uc: &usecases.TranslationsUseCase{},
	}
}

func (h *handler) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	definitions, err := h.uc.Execute(req.Text, req.SourceLang, req.TargetLang)
	if err != nil {
		return nil, err
	}


		pbDefs = append(pbDefs, &pb.Definition{
			Text: d.Text,
			Pos:  d.Pos,
			Tr:   pbTrs,
		})
	}

	return &pb.TranslateResponse{Definitions: pbDefs}, nil
}
