package user

import "time"

type Usuario struct {
	ID        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// AltaUsuario es una funcion que inicializa la estructura
func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	this.ID = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}
