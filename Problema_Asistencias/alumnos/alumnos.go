package alumnos

// Asistencias es un mapa de alumnos
var Asistencias = map[string]int{
	"HOMERO": 0,
	"BART":   0,
	"LISA":   1,
	"MARGE":  1,
	"MAGGIE": 1,
}

// TodoElMes es un mapa con la suma de todas las assistencias del mes
var TodoElMes = map[string]int{
	"Homero": 1,
	"Bart":   10,
	"Lisa":   16,
	"Marge":  14,
	"Maggie": 14,
}

// Asistentes es un slice inicialmente vacio para almacenar solo los presentes
var Asistentes []string

// Ausentes es un slice para almacenar los asusentes
var Ausentes []string
