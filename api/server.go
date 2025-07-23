package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"udemy-course/db/sqlc"
	"udemy-course/token"
	"udemy-course/utils"
)

type Server struct {
	store      sqlc.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     utils.Config
}

func NewServer(config utils.Config, store sqlc.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	if v, ok := binding.Validator.Engine().(validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server := &Server{store: store, router: gin.Default(), tokenMaker: tokenMaker, config: config}
	router := server.router

	router.POST("/user", server.createUser)
	router.POST("/login", server.loginUser)

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccount)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server, nil
}

func (server *Server) setupRouter() {
	router := server.router

	router.POST("/user", server.createUser)
	router.POST("/login", server.loginUser)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("/account", server.createAccount)
	authRouter.GET("/account/:id", server.getAccount)
	authRouter.GET("/account", server.listAccount)

	authRouter.POST("/transfer", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
