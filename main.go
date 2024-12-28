package main

import (
	"ToDoList/Initializers"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	Initializers.EnvInitializer()
	Initializers.DBInitializer()
	Initializers.DBMigration()
}

func main() {

	r := gin.Default()
	Initializers.SetupRoutes(r)
	r.Run(os.Getenv("PORT"))
}
