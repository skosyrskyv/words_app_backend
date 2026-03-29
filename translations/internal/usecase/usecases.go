package usecases

import (
	"log/slog"
	"translations/internal/domain/interfaces"
	"translations/internal/usecase/translations"
)

type UseCases struct {
	Translations *TranslationUseCases
}

type TranslationUseCases struct {
	Translate *translations.TranslateUseCase
}

func NewUseCases(translations *TranslationUseCases) *UseCases {
	return &UseCases{
		Translations: translations,
	}
}

func NewTranslationUseCases(
	repo interfaces.TranslationRepository,
	logger *slog.Logger,
) *TranslationUseCases {
	return &TranslationUseCases{
		Translate: translations.NewTranslateUseCase(repo, logger),
	}
}
