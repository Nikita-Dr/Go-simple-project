package entity

type Author struct {
	ID   int `bun:"id,pk,autoincrement"`
	Name string
}

func NewAuthorUponCreation(name string) Author {
	return Author{
		Name: name,
	}
}
