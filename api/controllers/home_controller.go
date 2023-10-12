package controllers

import (
	"net/http"

	"github.com/agushermanto22/go-gorm-jwt-postgres-mysql-and-testing/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
