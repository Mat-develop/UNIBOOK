package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"v1/monorepo/users/model"
	"v1/monorepo/users/repository"
	"v1/monorepo/util/authentication"
	dbconfig "v1/monorepo/util/db_config"
	"v1/monorepo/util/response"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

// Isso Tbm é handler
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	//throws to user struct
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := dbconfig.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.New(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, user.ID)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := dbconfig.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.New(db)
	if nameOrNick != "" {
		users, err := repo.FindUserByID(nameOrNick)
		if err != nil {
			response.Erro(w, http.StatusInternalServerError, err)
			return
		}
		response.JSON(w, http.StatusOK, users)
	} else {
		users, err := repo.FindUsers()
		if err != nil {
			response.Erro(w, http.StatusInternalServerError, err)
			return
		}
		response.JSON(w, http.StatusOK, users)
	}
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	tokenUserId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println(tokenUserId)
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("edit"); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	db, err := dbconfig.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.New(db)
	if err = repo.Update(userId, user); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := dbconfig.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.New(db)
	if err = repo.Delete(userId); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
