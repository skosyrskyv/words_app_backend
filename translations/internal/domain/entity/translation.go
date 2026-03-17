package entity

type Translation struct {
	Text     string
	Pos      string
	Synonyms []Synonym
	Meanings []Meaning
}
