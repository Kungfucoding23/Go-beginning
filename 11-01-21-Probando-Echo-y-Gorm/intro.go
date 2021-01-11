/*
GORM:
Es un framework que utilizaremos para todo lo que es el manejo de
bases de datos, ya que nos brinda una gran facilidad a la hora de trabajar con
las mismas,es una ORM (es decir que nos convierte la base de datos en
‘objetos’), esto nos ayuda a hacer que no tengamos que estar escribiendo
cada sentencia SQL básica(inserts, updates, wheres, etc) veamos un breve
ejemplo de lo que podemos hacer con GORM:
*/
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product ...
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Connect with your database
	db, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	if err != nil {
		panic("failed connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	/*
		En este ejemplo lo que hacemos es crear una estructura que implementa
		gorm.Model que lo que hace gorm.Model es llamar por defecto a esta
		estructura que nos brinda gorm:
	*/
	// Model definition
	type Model struct {
		ID        uint `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	/*
		Nosotros a las estructuras podemos decirles ya que tipo de dato van a ser en
		nuestra tabla de la base de datos, veamos otro ejemplo:
	*/

	type RequestDecision struct {
		ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
		RequestID  uint      `json:"request_id"`
		UUID       uuid.UUID `json:"uuid" gorm: "type:uuid; not null;"`
		Comment    string    `json:"comment" gorm:"type:text;"`
		CreatedAt  time.Time `json:"created_at" gorm:"not null;"`
		Department string    `json:"approved" gorm:"not null;"`
	}

	/*
		Luego, nos conectamos a la base de datos y como se puede ver en los
		ejemplos vamos llamando distintas funcionalidades que nos permite
		crear,actualizar,buscar y realizar sentencias de forma fácil y sencilla. GORM
		se puede trabajar en conjunto con Echo
	*/
}
