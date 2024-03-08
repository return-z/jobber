package api

import (
	"assets"
	"db"
	"fmt"
	"log"
	"net/http"
	"schema"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
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
  
  dbHandler, err := db.GetDBConn()
  fmt.Println("Establishing connection with the database...")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Database connected successfully")

  defer func() {
    if err := dbHandler.Close(); err != nil {
      log.Fatalf(err.Error())
    }
  }()

  router := gin.Default()
  router.Use(CORSMiddleware())

  //handle code for the templ renderer
  ginHTMLRenderer := router.HTMLRender
  router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHTMLRenderer}

  router.GET("/", func(c *gin.Context){
    applications, err := dbHandler.FetchApplications()
    if err != nil {
      log.Fatal(err)
    }
    c.HTML(http.StatusOK, "", assets.Home(applications))
  })

  router.POST("/addapp", func(c *gin.Context){
    application := &schema.Application{
      CompanyName: c.PostForm("company-name"),
      Role: c.PostForm("role"),
      ApplicationURL: c.PostForm("application-url"),
      Status: c.PostForm("status"),
    }
    err := dbHandler.AddApplication(application)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(application)
    c.HTML(http.StatusOK, "", assets.TableComponent(application))
  })

  router.Run(s.address)
}
