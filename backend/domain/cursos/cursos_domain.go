package cursos

import "backend/domain/inscripciones"

type Curso struct {
	IdCurso       int                         `gorm:"primary_key;column:Id_curso"`
	Titulo        string                      `gorm:"column:Titulo"`
	FechaInicio   string                      `gorm:"column:Fecha_inicio"`
	Categoria     string                      `gorm:"column:Categoria"`
	Archivo       string                      `gorm:"column:Archivo"`
	Descripcion   string                      `gorm:"column:Descripcion"`
	Inscripciones []inscripciones.Inscripcion `gorm:"foreignkey:IdCurso"`
}
