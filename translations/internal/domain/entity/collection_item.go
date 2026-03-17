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
	Source       string
	Translations []string
	User         uuid.UUID
	Tag          CollectionTag
}

func NewCollectionItem(source string, translations []string, user uuid.UUID, tag CollectionTag) *CollectionItem {
	return &CollectionItem{
		UUID:         uuid.New(),
		Source:       source,
		Translations: translations,
		User:         user,
		Tag:          tag,
	}
}
