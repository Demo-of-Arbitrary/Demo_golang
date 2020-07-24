package main

import (
  "fmt"
)
import "github.com/gin-gonic/gin"

type CreateTodoRequest struct {
  Description string 
}

func init() {
  model.InitModel()
}

func main() {
  fmt.Println("Hello, playground")
  g := gin.New()
  g.Use(gin.Logger())
  g.Use(gin.Recovery())

  todoRoute := g.Group("/api/todo")
  todoRoute.use()
  {
    todoRoute.Get("", func(c *gin.Context){
      c.Status(http.StatusOK)
    })

    todoRoute.POST("", func (c *gin.Context) {
      var req CreateTodoRequest
      todo, err := c.BindJSON(&req)

      if err != nil {
          c.Status(http.StatusInternalServerError)
      }
    })

  }

  err := g.Run("8888")

  if err != nil {
    fmt.Println("There is an error")
  }
}
