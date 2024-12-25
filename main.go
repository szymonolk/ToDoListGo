package main

import (
	"ToDoList/Initializers"
	"github.com/gin-gonic/gin"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func init() {
	Initializers.EnvInitializer()
	Initializers.DBInitializer()
	Initializers.DBMigration()
}

func main() {
	r := gin.Default()
	r.Run(os.Getenv("PORT"))

}
