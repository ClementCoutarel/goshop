package database

type AuthRepo interface {
	Register()
	Signin()
}
