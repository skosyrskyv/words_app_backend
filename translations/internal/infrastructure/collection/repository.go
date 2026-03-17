package collection

import (
	"translations/internal/domain/entity"
	"translations/pkg/postgres"
)

type repository struct {
	postgres *postgres.Postgres
}

func NewRepository(pg *postgres.Postgres) *repository {
	return &repository{
		postgres: pg,
	}
}

func (*repository) AddItem(item entity.CollectionItem) error {
	return nil
}

func (*repository) GetCollection(user string) ([]entity.CollectionItem, error) {
	return nil, nil
}

func (*repository) RemoveItem(uuid string, user string) error {
	return nil
}

func (*repository) UpdateItem(item entity.CollectionItem) (entity.CollectionItem, error) {
	return item, nil
}

func UpdateItemsBatch(items []entity.CollectionItem) error {
	return nil
}
