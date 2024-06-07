package user

import (
	"context"
	"goWebUsers/internal/domain"
	"log"
)


/*
En resumen, tu implementación define un servicio (Service) para manipular usuarios. 
Utiliza una interfaz para definir las operaciones disponibles (Create y GetAll), 
una estructura service para implementar esta interfaz, un constructor para crear instancias del servicio,
 y métodos que interactúan con un repositorio (Repository) para realizar operaciones específicas sobre los usuarios,
 todo mientras registra eventos importantes usando un logger.
 
*/


//Se define la estructura Service con su interfaz
type (
	Service interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

//Se define Constructor para instanciar Service
func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}


// 	--- Métodos para interactuar con Repository ---

//No usa asterisco* en service porque no lo modificará
func (s service) Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:   lastName,
		Email:      email,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	s.log.Println("service create")
	return user, nil
}

func (s service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	s.log.Println("service get all")
	return users, nil
}
