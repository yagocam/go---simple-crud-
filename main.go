package main

import (
	"awesomeProject/internal/controllers"
	"awesomeProject/internal/models"

	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	r := gin.Default()
	models.ConnectDatabase()
	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
	}
	r.Run(":8080")
}
