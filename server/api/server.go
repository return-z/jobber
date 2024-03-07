package api

import (
	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
  "github.com/return-z/jobber/assets"
)

type Server struct {
  address string
}

func NewServer(address string) *Server {
  return &Server{
    address: address,
  }
}

func (s *Server) Start() {
  
  router := gin.Default()
  router.Use(CORSMiddleware())

  //handle code for the templ renderer
  ginHTMLRenderer := router.HTMLRender
  router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHTMLRenderer}

  router.GET("/", func(c *gin.Context){
    c.HTML(http.StatusOK, "", assets.Greeting())
  })
}
