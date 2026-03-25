package translation

import (
	"auth/pkg/postgres"
	"translations/pkg/httpclient"
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

func (r *repository) GetTranslation(text, sourceLang, targetLang string) ([]string, error) {

	return nil, nil
}
