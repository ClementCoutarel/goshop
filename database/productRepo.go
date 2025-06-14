package database

type ProductRepository interface {
	GetAll()
	GetById()
	Create()
	UpdateOne()
	DeleteOne()
}
