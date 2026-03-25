package usecases

import "translations/internal/domain/entity"

type TranslationsUseCase struct {
}

func (*TranslationsUseCase) Execute(text, sourceLang, targetLang string) ([]entity.Definition, error) {
	return nil, nil
}
