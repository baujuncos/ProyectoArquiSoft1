package cursos

import (
	cursosDomain "backend/domain/cursos"
	cursosServices "backend/services/cursos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteCurso(c *gin.Context) {
	cursoID := c.Param("id")

	err := cursosServices.DeleteCurso(cursoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso eliminado correctamente"})
}
func UpdateCurso(c *gin.Context) {
	cursoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var updatedCurso cursosDomain.Curso
	if err := c.ShouldBindJSON(&updatedCurso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cursosServices.UpdateCurso(cursoID, updatedCurso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso editado correctamente"})
}

func CreateCurso(c *gin.Context) {
	var curso cursosDomain.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cursosServices.CreateCurso(curso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso creado correctamente"})
}
