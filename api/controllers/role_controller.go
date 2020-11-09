package controllers

import (
	"net/http"

	"github.com/kiojio/gorest/api/models"
	"github.com/kiojio/gorest/api/responses"
)

func (server *Server) GetRoles(w http.ResponseWriter, r *http.Request) {

	role := models.Role{}
	roles, err := role.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, roles)
}
