package database

type UserRepository interface {
	GetAll()
	GetById()
	Create()
	DeleteOne()
	UpdateOne()
}
