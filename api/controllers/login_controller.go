package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kiojio/gorest/api/auth"
	"github.com/kiojio/gorest/api/models"
	"github.com/kiojio/gorest/api/responses"
	"github.com/kiojio/gorest/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	ret,err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, ret)
}

func (server *Server) SignIn(email, password string) (interface{}, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	dataResponse := map[string]interface{}{}
	dataResponse["nickname"] = user.Nickname
	dataResponse["email"] = user.Email
	dataResponse["token"] = token

	return dataResponse, err
}
