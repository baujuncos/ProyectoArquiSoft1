package router

import (
	"backend/controllers/cursos"
	"backend/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ruta para el login
	//r.POST("/api/v1/login", users.Login)

	// Ruta para eliminar un curso
	r.DELETE("/cursos/:id", cursos.DeleteCurso)

	// Ruta para actualizar un curso
	r.PUT("/cursos/:id", cursos.UpdateCurso)

	r.POST("/cursos", cursos.CreateCurso)

	return r
}

func main() {
	db.InitDB() // Inicializar la base de datos antes de empezar el servidor
	r := SetupRouter()
	r.Run(":8080") // Por defecto corre en :8080
}
