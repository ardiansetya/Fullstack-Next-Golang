package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"jamal/api/api/config"
	"jamal/api/api/helper"
	"jamal/api/api/models/domain"
	web2 "jamal/api/api/models/web"
	"jamal/api/api/repository"
	"net/http"
	"time"
)

type AuthServiceImpl struct {
	DB             *gorm.DB
	AuthRepository repository.AuthRepository
}

func NewAuthService(db *gorm.DB, authRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{
		DB:             db,
		AuthRepository: authRepository,
	}
}

func (service AuthServiceImpl) Register(ctx *gin.Context, register web2.AuthRequestRegister) web2.WebResponse {
	var response web2.WebResponse
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		//<-- hash password
		hashPsw, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
		if err != nil { // <-- response ketika gagal hash
			response = web2.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "InternalServerError",
				Data: map[string]interface{}{
					"Error": "Error Hash Password",
				},
			}
			return nil
		}
		register.Password = string(hashPsw)

		//<-- add format request register
		registerData := domain.User{
			UserName: register.UserName,
			Password: register.Password,
			Name:     register.Name,
		}

		err = service.AuthRepository.Register(tx, registerData)
		if err != nil { //<-- response ketika gagal create
			response = web2.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   err.Error(),
			}
			return nil
		}

		//<-- response sukses
		response = web2.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Register Success",
		}
		return nil
	})
	helper.HandleErrorResponse(&response, err) // <-- ketika error transaction

	return response
}

func (service AuthServiceImpl) Login(ctx *gin.Context, login web2.AuthRequestLogin) web2.WebResponse {
	var response web2.WebResponse
	err := service.DB.Transaction(func(tx *gorm.DB) error {

		// <-- add adta format
		requestLogin := domain.User{
			UserName: login.UserName,
			Password: login.Password,
			Name:     "",
		}

		fmt.Println("service ", requestLogin)

		err := service.AuthRepository.Login(tx, requestLogin)
		if err != nil { //<-- response ketika gagal create
			response = web2.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   err.Error(),
			}
			return nil
		}

		//<-- response sukses
		response = web2.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Login Success",
		}

		return nil
	})
	helper.HandleErrorResponse(&response, err) // <-- ketika error transaction

	//<-- Generate JWT TOKEN
	expTime := time.Now().Add(time.Minute * 2) // <-- token kadaluarsa dalam 2min
	claimToken := &config.JWTClaim{
		UserName: login.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "king_jamal",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// <-- generate algorithm yg akan di gunakan untuk login
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	// <-- sign token
	token, err := tokenAlgo.SignedString([]byte(config.JWT_KEY))

	// <-- set token ke cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	return response
}

func (service AuthServiceImpl) Logout(ctx *gin.Context) web2.WebResponse {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	response := web2.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]interface{}{
			"message": "Logout Success",
		},
	}

	return response
}
