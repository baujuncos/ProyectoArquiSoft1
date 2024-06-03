package db

import (
	"backend/dao"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB
var sqlDB *sql.DB

func InitDB() {
	//dsn := "root:ladrillo753@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	dsn := "root:RaTa8855@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	//dsn := "root:ladrillo753@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	//dsn := "root:ladrillo753@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Obtener la conexión SQL nativa de gorm
	sqlDB, err = DB.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm: ", err)
	}

	// Migrar las estructuras a la base de datos
	Migrate()

	// Seed the database with initial data
	SeedDB()
}

func Migrate() {

	// Migrar tablas en el orden correcto
	err := DB.Migrator().CreateTable(&dao.User{})
	if err != nil {
		log.Fatal("failed to migrate User table: ", err)
	}

	err = DB.Migrator().CreateTable(&dao.Curso{})
	if err != nil {
		log.Fatal("failed to migrate Curso table: ", err)
	}

	err = DB.Migrator().CreateTable(&dao.Inscripcion{})
	if err != nil {
		log.Fatal("failed to migrate Inscripcion table: ", err)
	}
}

func SeedDB() {
	// Crear usuarios
	users := []dao.User{
		{NombreUsuario: "pauliiortizz", Nombre: "paulina", Apellido: "ortiz", Email: "pauliortiz@example.com", Contrasena: "contraseña1", Tipo: false},
		{NombreUsuario: "baujuncos", Nombre: "bautista", Apellido: "juncos", Email: "baujuncos@example.com", Contrasena: "contraseña2", Tipo: true},
		{NombreUsuario: "belutreachi", Nombre: "belen", Apellido: "treachi", Email: "belutreachi2@example.com", Contrasena: "contraseña3", Tipo: false},
		{NombreUsuario: "virchu", Nombre: "virginia", Apellido: "rodriguez", Email: "virchurodiguez@example.com", Contrasena: "contraseña4", Tipo: false},
		{NombreUsuario: "johndoe", Nombre: "John", Apellido: "Doe", Email: "johndoe@example.com", Contrasena: "contraseña5", Tipo: false},
		{NombreUsuario: "alicesmith", Nombre: "Alice", Apellido: "Smith", Email: "alicesmith@example.com", Contrasena: "contraseña6", Tipo: true},
		{NombreUsuario: "bobjohnson", Nombre: "Bob", Apellido: "Johnson", Email: "bobjohnson@example.com", Contrasena: "contraseña7", Tipo: false},
		{NombreUsuario: "janedoe", Nombre: "Jane", Apellido: "Doe", Email: "janedoe@example.com", Contrasena: "contraseña8", Tipo: false},
		{NombreUsuario: "emilywilliams", Nombre: "Emily", Apellido: "Williams", Email: "emilywilliams@example.com", Contrasena: "contraseña9", Tipo: true},
	}

	for i, user := range users {
		// Hashear la contraseña con SHA-1
		hasher := sha1.New()
		hasher.Write([]byte(user.Contrasena))
		hashedPassword := hex.EncodeToString(hasher.Sum(nil))
		users[i].Contrasena = hashedPassword
		DB.FirstOrCreate(&users[i], dao.User{Email: user.Email})
	}

	// Crear cursos
	cursos := []dao.Curso{
		{Titulo: "Curso de Go", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso.pdf", Descripcion: "Curso avanzado de Go"},
		{Titulo: "Curso de Python", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_python.pdf", Descripcion: "Curso básico de Python"},
		{Titulo: "Curso de Java", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_java.pdf", Descripcion: "Curso intermedio de Java"},
		{Titulo: "Curso de C++", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_cpp.pdf", Descripcion: "Curso básico de C++"},
		{Titulo: "Curso de JavaScript", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_js.pdf", Descripcion: "Curso completo de JavaScript"},
	}
	for _, curso := range cursos {
		DB.FirstOrCreate(&curso, dao.Curso{Titulo: curso.Titulo})
	}

	// Crear inscripciones
	inscripciones := []dao.Inscripcion{
		{IdUsuario: 1, IdCurso: 1, FechaInscripcion: time.Now()},
		{IdUsuario: 2, IdCurso: 2, FechaInscripcion: time.Now()},
		{IdUsuario: 3, IdCurso: 3, FechaInscripcion: time.Now()},
		{IdUsuario: 4, IdCurso: 4, FechaInscripcion: time.Now()},
		{IdUsuario: 5, IdCurso: 5, FechaInscripcion: time.Now()},
		{IdUsuario: 1, IdCurso: 2, FechaInscripcion: time.Now()},
		{IdUsuario: 2, IdCurso: 3, FechaInscripcion: time.Now()},
		{IdUsuario: 3, IdCurso: 4, FechaInscripcion: time.Now()},
		{IdUsuario: 4, IdCurso: 5, FechaInscripcion: time.Now()},
		{IdUsuario: 5, IdCurso: 1, FechaInscripcion: time.Now()},
	}
	for _, inscripcion := range inscripciones {
		DB.FirstOrCreate(&inscripcion, dao.Inscripcion{IdUsuario: inscripcion.IdUsuario, IdCurso: inscripcion.IdCurso})
	}
}

func DeleteCursoByID(IdCurso string) error {
	result := DB.Delete(&dao.Curso{}, "Id_curso = ?", IdCurso)
	return result.Error
}

func DeleteInscripcionesByCursoID(cursoID string) error {
	query := `DELETE FROM inscripciones WHERE Id_curso = ?`
	_, err := sqlDB.Exec(query, cursoID)
	return err
}

func GetCursosUsuario(userId int) ([]dao.Curso, error) {
	// Obtener los IDs de los cursos del usuario desde la tabla de inscripciones
	var inscripciones []struct {
		IdCurso int `gorm:"column:Id_curso"`
	}
	if err := DB.Table("inscripciones").
		Where("Id_usuario = ?", userId).
		Pluck("Id_curso", &inscripciones).
		Error; err != nil {
		log.Printf("Error al obtener inscripciones del usuario: %v\n", err)
		return nil, err
	}

	// Extraer los IDs de los cursos de la lista de inscripciones
	var cursoIDs []int
	for _, inscripcion := range inscripciones {
		log.Printf("Inscripción encontrada: %+v\n", inscripcion)
		cursoIDs = append(cursoIDs, inscripcion.IdCurso)
	}

	if len(cursoIDs) == 0 {
		log.Println("El usuario no está inscrito en ningún curso")
		return []dao.Curso{}, nil
	}

	// Buscar los cursos correspondientes a los IDs obtenidos
	var cursos []dao.Curso
	if err := DB.Where("Id_curso IN ?", cursoIDs).Find(&cursos).Error; err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}

	// Verificar que los datos se han obtenido correctamente
	for _, curso := range cursos {
		log.Printf("Curso obtenido: %+v\n", curso)
	}

	return cursos, nil
}

func GetUsuarioByUsername(username string) (dao.User, error) {
	var usuario dao.User

	result := DB.Where("Nombre_Usuario = ?", username).First(&usuario)

	if result.Error != nil {
		return dao.User{}, fmt.Errorf("No se encontro el usuario: %s", username)
	}

	return usuario, nil
}

func FindCoursesByQuery(query string) ([]dao.Curso, error) {
	var cursos []dao.Curso
	result := DB.Where("Titulo LIKE ? OR Descripcion LIKE ?", "%"+query+"%", "%"+query+"%").Find(&cursos)
	if result.Error != nil {
		return nil, result.Error
	}
	return cursos, nil
}

func FindCourseByID(id int) (dao.Curso, error) {
	var curso dao.Curso
	result := DB.First(&curso, id)
	if result.Error != nil {
		return dao.Curso{}, fmt.Errorf("no se encontró el curso con el ID: %d", id)
	}
	return curso, nil
}

func SubscribeUserToCourse(userID int, courseID int) error {
	var inscripcion dao.Inscripcion
	result := DB.Where("Id_usuario = ? AND Id_curso = ?", userID, courseID).First(&inscripcion)
	if result.Error == nil {
		return fmt.Errorf("el usuario %d ya está suscrito al curso %d", userID, courseID)
	}

	inscripcion = dao.Inscripcion{
		IdUsuario:        userID,
		IdCurso:          courseID,
		FechaInscripcion: time.Now().UTC(),
	}

	result = DB.Create(&inscripcion)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func SelectUserByID(id int) (dao.User, error) {
	var user dao.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("No se encontro el usuario con el id: %d", id)
	}
	return user, nil
}
