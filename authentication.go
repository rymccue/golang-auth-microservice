package main

import (
	"database/sql"
	"github.com/goadesign/goa"
	"github.com/rymccue/golang-auth-microservice/app"
	"github.com/rymccue/golang-auth-microservice/repositories"
	"github.com/rymccue/golang-auth-microservice/utils/crypto"
	"github.com/rymccue/golang-auth-microservice/utils/jwt"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
	*sql.DB
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service, db *sql.DB) *AuthenticationController {
	return &AuthenticationController{
		Controller: service.NewController("AuthenticationController"),
		DB:         db,
	}
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	payload := ctx.Payload
	u, err := repositories.GetUserByEmail(c.DB, payload.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
		}
		c.Service.LogError("Login User", "err", err)
		return ctx.InternalServerError()
	}
	hashedPassword := crypto.HashPassword(payload.Password, *u.Salt)
	if *u.Password != hashedPassword {
		return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
	}
	token, err := jwt.CreateJWTToken(*u.Email)
	if err != nil {
		c.Service.LogError("Login User", "err", err)
		return ctx.InternalServerError()
	}
	res := &app.Token{
		Token: &token,
	}
	return ctx.OK(res)
}

// Register runs the register action.
func (c *AuthenticationController) Register(ctx *app.RegisterAuthenticationContext) error {
	payload := ctx.Payload
	exists, err := repositories.CheckEmailExists(c.DB, payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	if exists {
		return ctx.BadRequest(goa.ErrBadRequest("Email already exists"))
	}
	err = repositories.AddUserToDatabase(c.DB, payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	token, err := jwt.CreateJWTToken(payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	res := &app.Token{Token: &token}
	return ctx.OK(res)
}
