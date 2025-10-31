package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"v1/users/model"
	"v1/users/repository"
	"v1/util/authentication"
	dbconfig "v1/util/db_config"
	"v1/util/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
	}

	var user model.User
	if err = json.Unmarshal(request, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := dbconfig.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userAuth, err := repo.FindUserByEmail(user.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = authentication.Verify(userAuth.Password, user.Password); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(strconv.FormatUint(userAuth.ID, 10))
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}
