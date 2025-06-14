package database

type ProductService struct {
	Repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}
