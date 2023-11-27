package book

type Book struct {
	ID    int
	Title string
}

func (Book) TableName() string {
	return "books"
}
