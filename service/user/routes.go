package user

import (
	"fmt"
	"net/http"

	"github.com/ayushn2/go_ecom.git/service/auth"
	"github.com/ayushn2/go_ecom.git/types"
	"github.com/ayushn2/go_ecom.git/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler{
	return &Handler{store : store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin).Methods("POST")
	router.HandleFunc("/register",h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest,err)
		return
	}

	// Verify the payload
	if err := utils.Validate.Struct(payload); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid payload %v",errors))
		return
	}

	// Check if the user exists
	_,err := h.store.GetUserByEmail(payload.Email)
	if err == nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return 
	}

	hashedPassword,err := auth.HashPassword(payload.Password)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError,err)
		return 
	}
	// if it doesn't we create the new user 
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email : payload.Email,
		Password : hashedPassword,
	})

	if err != nil{
		utils.WriteError(w,http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w,http.StatusCreated,nil)
}