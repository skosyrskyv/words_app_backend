package interfaces

import "translations/internal/domain/entity"

type CollectionRepository interface {
	AddItem(item entity.CollectionItem) error
	GetCollection(user string) ([]entity.CollectionItem, error)
	RemoveItem(uuid string, user string) error
	UpdateItem(item entity.CollectionItem) (entity.CollectionItem, error)
	UpdateItemsBatch(items []entity.CollectionItem) error
}
