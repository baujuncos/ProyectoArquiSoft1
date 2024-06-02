package dao

type User struct {
	IdUsuario     int    `gorm:"primary_key;column:Id_usuario;autoIncrement"`
	NombreUsuario string `gorm:"column:Nombre_Usuario"`
	Nombre        string `gorm:"column:Nombre"`
	Apellido      string `gorm:"column:Apellido"`
	Email         string `gorm:"column:Email"`
	Contrasena    string `gorm:"column:Contraseña"`
	Tipo          bool   `gorm:"column:Tipo"`
}
