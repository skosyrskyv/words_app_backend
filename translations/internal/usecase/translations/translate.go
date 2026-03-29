package translations

import (
	"log/slog"
	"translations/internal/domain/entity"
	"translations/internal/domain/interfaces"
)

type TranslateUseCase struct {
	repo   interfaces.TranslationRepository
	logger *slog.Logger
}

func (uc *TranslateUseCase) Execute(text, sourceLang, targetLang string) ([]entity.Definition, error) {
	return nil, nil
}

func NewTranslateUseCase(repo interfaces.TranslationRepository, logger *slog.Logger) *TranslateUseCase {
	return &TranslateUseCase{
		repo:   repo,
		logger: logger,
	}
}
