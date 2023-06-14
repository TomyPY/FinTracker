package handlers

import (
	"github.com/TomyPY/FinTracker/internal/user"
	"github.com/gin-gonic/gin"
)

type User struct {
	s user.Service
}

func NewUser(s user.Service) *User {
	return &User{s: s}
}

// GetAll handler
// @Summary Get all users
// @Schemes
// @Description Show an array of users
// @Tags User
// @Produce json
// @Success 200 {array} user.User
// @Failure 500 {object} web.errorResponse
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Get handler
// @Summary Show a user by id
// @Schemes
// @Description Get just one user by his ID
// @Tags User
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} user.User
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /users/{id} [get]
func (u *User) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Update handler
// @Summary Create user information
// @Schemes
// @Description Create user information
// @Tags User
// @Produce json
// @Param user body user.User true "User"
// @Success 200 {object} user.User
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /users [post]
func (u *User) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Update handler
// @Summary Update user information
// @Schemes
// @Description Update user information by his id
// @Tags User
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} user.User
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /users/{id} [post]
func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// Delete handler
// @Summary Delete a user
// @Schemes
// @Description Delete a user from db
// @Tags User
// @Produce json
// @Param userId path int true "User ID"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /users/{id} [delete]
func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
