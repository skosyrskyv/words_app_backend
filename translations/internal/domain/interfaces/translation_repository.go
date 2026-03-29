package interfaces

import "translations/internal/domain/entity"

type TranslationRepository interface {
	Translate(text string, sourceLang string, targetLang string) ([]entity.Definition, error)
}
