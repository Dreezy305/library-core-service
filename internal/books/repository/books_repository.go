package repository

type BookRepository struct {
	gormRepo *GormBookRepository
}

func NewBookRepository(gormRepo *GormBookRepository) *BookRepository {
	return &BookRepository{gormRepo: gormRepo}
}

func (s *BookRepository) GetBooks() error {
	return s.gormRepo.GetBooks()
}

func (s *BookRepository) GetBook(bookId string) error {
	return s.gormRepo.GetBook(bookId)
}

func (s *BookRepository) CreateBook() error {
	return s.gormRepo.CreateBook()
}

func (s *BookRepository) UpdateBook() error {
	return s.gormRepo.UpdateBook()
}
