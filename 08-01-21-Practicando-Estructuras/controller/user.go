package controller

import "time"

//Usuario ...
type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//AltaUsuario es una funcion que agrega un usuario. this *Usuario nos permite hacer referencia a Usuario cada vez que ponga this.
// func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
// 	this.ID = id
// 	this.Nombre = nombre
// 	this.FechaAlta = fechaAlta
// 	this.Status = status
// }
func AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) Usuario {
	return Usuario{
		ID:        id,
		Nombre:    nombre,
		FechaAlta: fechaAlta,
		Status:    status,
	}

}
