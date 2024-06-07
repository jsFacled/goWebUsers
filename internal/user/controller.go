package user

/*
Video:
- Hasta el minuto 8 se define el mètodo GetAllUser
- 8´: metodo Post.
- 12: cmd/main


*/
import (
	"context"
	"encoding/json"
	"fmt"
	"goWebUsers/internal/domain"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			GetAllUser(ctx, s, w)

		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var user domain.User
			if err := decode.Decode(&user); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, user)

		default:
			InvalidMethod(w)
		}
	}
}

func GetAllUser(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

// Método
func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateReq)

	if req.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "first name is required")
		return
	}
	if req.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "last name is required")
		return
	}
	if req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "email is required")
		return
	}

	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusCreated, user)
}

// Método que da un mensaje en la Respuesta
func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}

// Método para convertir Entidiad Usuarios a json.
// /Usamos paquete encoding, la funcion Marshall transforma una Estructura en un Json
func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value) // En el campo data, el valor %s va sin comillas para que tome el json.

}

func InvalidMethod(w http.ResponseWriter){
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "metohod doesn´t exist}`, status)
}
