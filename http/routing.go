package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miron239/wb/config"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TaskServiceHTTPHandlers defines all the handlers the TaskService needs. It's
// possible to register routes for a different implementation (like a mock).
type TaskServiceHTTPHandlers interface {
	GetTask(c *gin.Context)
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DeleteTasks(c *gin.Context)
}

type Server struct {
	router *gin.Engine
	port   int
}

func InitServer(conf *config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		port:   conf.Server.Port,
	}
	if len(conf.Frontend.Endpoints) > 0 {
		server.router.Use(cors.New(cors.Config{
			AllowOrigins: conf.Frontend.Endpoints,
			AllowMethods: []string{"GET", "POST", "DELETE"},
			AllowHeaders: []string{"Origin", "Authorization", "Content-Type", "CallerId"},
			MaxAge:       12 * time.Hour,
		}))
	}
	if conf.Authn.NotEnforced {
		server.router.Use(authHandler(conf.Authn.ClientId, optional))
	} else {
		server.router.Use(authHandler(conf.Authn.ClientId, mandatory))
	}
	return server
}

func (s Server) Router() *gin.Engine {
	return s.router
}

func (s Server) RegisterRoutes(t TaskServiceHTTPHandlers) {
	s.router.GET("/blacklist/", t.GetTasks)
	s.router.GET("/blacklist/:taskid", t.GetTask)
	s.router.POST("/blacklist/", t.CreateTask)
	s.router.DELETE("/blacklist/:taskid", t.DeleteTask)
	s.router.DELETE("/blacklist/", t.DeleteTasks)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

}

func (s Server) Start() {
	fmt.Printf("Server at %d\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
