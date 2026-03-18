package entity

import (
	"github.com/google/uuid"
)

type CollectionTag = string

const (
	RED    CollectionTag = "red"
	PINK   CollectionTag = "pink"
	PURPLE CollectionTag = "purple"
	BLUE   CollectionTag = "blue"
	GREEN  CollectionTag = "green"
	YELLOW CollectionTag = "yellow"
	ORANGE CollectionTag = "orange"
	GREY   CollectionTag = "grey"
	NON    CollectionTag = "non"
)

type CollectionItem struct {
	UUID         uuid.UUID
	User         uuid.UUID
	Collection   uuid.UUID
	Tag          CollectionTag
	Source       string
	Translations []string
}

func NewCollectionItem(source string, translations []string, user uuid.UUID, collectionUUID uuid.UUID, tag CollectionTag) *CollectionItem {
	return &CollectionItem{
		UUID:         uuid.New(),
		User:         user,
		Collection:   collectionUUID,
		Tag:          tag,
		Source:       source,
		Translations: translations,
	}
}
