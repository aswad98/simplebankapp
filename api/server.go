package api

import (
	"fmt"

	"github.com/minibank/token"
	"github.com/minibank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/minibank/db/sqlc"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("token is not genrated: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/login", server.loginUser)
	router.POST("/user", server.creatUser)

	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoute.POST("/account", server.creatAccount)
	authRoute.GET("/account/:id", server.getAccount)
	authRoute.GET("/accounts", server.listAccount)
	authRoute.POST("/transfer", server.createTransfer)

	server.router = router

}
