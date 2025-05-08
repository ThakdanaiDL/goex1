package repository

type HelloRepositoryInterface interface {
	GetMessage() string
}

type helloRepository struct{}

func NewHelloRepository() HelloRepositoryInterface {
	return &helloRepository{}
}

func (r *helloRepository) GetMessage() string {

	return "Hello, World ! From Repository"
}
