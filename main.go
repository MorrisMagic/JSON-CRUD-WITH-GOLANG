package main

import (
	"crud/controllers"
	"crud/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.Migrate()

}
func main() {
	app := gin.Default()
	app.POST("/create", controllers.Posts)
	app.GET("/posts", controllers.PostIndex)
	app.GET("/post/:id", controllers.OnePost)
	app.PUT("/post/:id", controllers.UpdatePost)
	app.DELETE("/dpost/:id", controllers.DeletePost)

	app.Run(os.Getenv("PORT"))
}
