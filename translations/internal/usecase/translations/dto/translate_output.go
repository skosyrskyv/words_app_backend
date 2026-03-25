package dto

import "translations/internal/domain/entity"

type TranslateOutput struct {
	Definitions []entity.Definition
}
