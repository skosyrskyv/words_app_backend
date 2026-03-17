package interfaces

type TranslationRepository interface {
	Translate(text string, sourceLang string, targetLang string) (interface{}, error)
}
