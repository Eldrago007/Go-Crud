package main

// To compile run-CompileDaemon -command="./go-curd" in terminal by going to the folder
//Path-/Users/eldrago/go/src/github.com/skondrag/go-curd

import (
	"github.com/gin-gonic/gin"
	"github.com/skondrag/go-curd/controllers"
	"github.com/skondrag/go-curd/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
