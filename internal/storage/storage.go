package storage

type Storage struct {
	// In a real application, this would hold a database connection pool.
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) GetMessage() string {
	// In a real application, this would fetch data from a database.
	return "Hello from the storage layer!"
}
