package repository

type helloRepository struct{}

func NewHelloRepository() HelloRepositoryInterface {
	return &helloRepository{}
}

func (r *helloRepository) GetMessage() string {

	return "Hello, World ! From Repository"
}
