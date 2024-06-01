package user

import (
	"context"
	"goWebUsers/internal/domain"
	"log"
)

// Generamos un esquema de Base de Datos.
//DB es una estructura que simula una base de datos en memoria
type DB struct {
	Users     []domain.User	//slice de usuarios
	MaxUserID uint64
}

type (
	// Interfaz para definir metodos.
	//Cualquier tipo que implemente estos métodos puede considerarse un Repository.
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepo(db DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

// - - - Métodos de repo - - - 

//El asterisco* permite ingresar a la instancia repo y modificarla, lo mismo con User
//Aquí no estoy creando un usuario, simplemente agrego un user a la BD el cual vendrá por parámetro
func (r *repo) Create(ctx context.Context, user *domain.User) error{
	r.db.MaxUserID++						//Accede al campo db de repo y aumenta el id
	user.ID = r.db.MaxUserID				//asigna el valor del nuevo id
	r.db.Users = append(r.db.Users, *user)	//agrega el user al slice
	r.log.Println("repository create")		//registra un mensaje en el log de repo 
	return nil								//Devuelve nil si está todo ok
}

//Devuelve el slice de usuarios almacenados y nil para indicar que no hubo errores.
func (r *repo) GetAll(ctx context.Context) ([]domain.User, error){
r.log.Println("repositoryget all")
return r.db.Users, nil
}

