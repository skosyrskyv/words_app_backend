package grpchandlers

import (
	"context"

	usecase "translations/internal/usecase/translations"

	pb "words-app.local/protos/gen/translations"
)

type TranslationsHandler struct {
	pb.UnimplementedTranslationsServer
	uc *usecase.TranslateUseCase
}

func New(uc *usecase.TranslateUseCase) *TranslationsHandler {
	return &TranslationsHandler{
		uc: uc,
	}
}

func (h *TranslationsHandler) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	return &pb.TranslateResponse{
			Definitions: []*pb.Definition{
				{
					Text: "Example translation",
					Pos:  "noun",
					Tr: []*pb.Translation{
						{
							Text:     "Пример перевода",
							Pos:      "существительное",
							Synonyms: []*pb.Synonym{{Text: "пример"}, {Text: "образец"}},
							Meanings: []*pb.Meaning{{Text: "пример для иллюстрации"}},
						},
					},
				},
			},
		},
		nil
}
