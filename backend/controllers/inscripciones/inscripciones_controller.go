package inscripciones

import (
	"backend/domain"
	"backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Subscribe(c *gin.Context) {
	var request domain.Inscripcion

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	subscribed, err := services.IsSubscribed(request.IdUsuario, request.IdCurso)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Message: fmt.Sprintf("Error checking subscription: %s", err.Error()),
		})
		return
	}

	if subscribed {
		c.JSON(http.StatusConflict, domain.Response{
			Message: fmt.Sprintf("Usuario %d ya está suscrito al curso %d", request.IdUsuario, request.IdCurso),
		})
		return
	}

	// Agregar fecha de inscripción y comentario
	fechaInscripcion := time.Now()
	if err := services.Subscribe(request.IdUsuario, request.IdCurso, fechaInscripcion, request.Comentario); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Message: fmt.Sprintf("Error al subscribirse: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Response{
		Message: fmt.Sprintf("Usuario %d inscripto exitosamente al curso %d", request.IdUsuario, request.IdCurso),
	})
}
