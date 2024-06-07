package user

import (
	"context"
	"encoding/json"
	"goWebUsers/internal/domain"
	"net/http"
)

type(
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}

	CreateReq struct{
		FirstName string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller{
	return func (w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case http.MethodGet:
			GetAllUser(ctx, s, w)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var user domain.User 
			if err := decode.Decode(&user); err != nil{
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w,user)
		default:
			InvalidMethod(w)
		}
	}
}

func GetAllUser(ctx context.Context, s Service, w http.ResponseWriter ){
	users, err := s.GetAll(ctx)
	if err != nil{
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}
