package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joao3101/daniel-joao-project/token"
	"github.com/joao3101/daniel-joao-project/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	/*router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)*/

	/*authRoutes := router.Group("/").Use(app.AuthMiddleware(server.tokenMaker))
	authRoutes.POST("/athletes", server.CreateAthletes)
	authRoutes.GET("/athletes/:id", server.GetAthlete)
	authRoutes.GET("/athletes", server.ListAthletes)*/

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
