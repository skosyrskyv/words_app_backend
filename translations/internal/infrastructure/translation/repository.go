package translation

import (
	"translations/internal/domain/entity"
	"translations/pkg/httpclient"
	"translations/pkg/postgres"
)

type repository struct {
	httpClient *httpclient.Client
	postgres   *postgres.Postgres
}

func NewRepository(httpClient *httpclient.Client, postgres *postgres.Postgres) *repository {
	return &repository{
		httpClient: httpClient,
		postgres:   postgres,
	}
}

func (r *repository) Translate(text, sourceLang, targetLang string) ([]entity.Definition, error) {

	return nil, nil
}
