package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type animal interface {
	respirar()
	//comer() // lo puedo comentar xq no lo estoy usando
	especie() string
}

type perro struct {
	cantPatas  int
	respirando bool
	comiendo   bool
	ladrando   bool
	vivo       bool
}
type gato struct {
	cantPatas  int
	respirando bool
	comiendo   bool
	maullando  bool
}

// Deben implementarse todos los metodos de la interfaz
func (p *perro) respirar() { p.respirando = true }

//func (p *perro) comer()          { p.comiendo = true }
func (p *perro) especie() string { return "Perro" }

// Puede implementarse un metodo de otra interfaz
func (p *perro) estaVivo() bool { return p.vivo }

func (p *gato) respirar() { p.respirando = true }

//func (p *gato) comer()          { p.comiendo = true }
func (p *gato) especie() string { return "Gato" }

//AnimalRespirando es una funcion que toma como parametro anim del tipo animal
func AnimalRespirando(anim animal) {
	anim.respirar()
	fmt.Printf("Soy un %s y estoy respirando\n", anim.especie())
}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	Michi := new(gato)
	AnimalRespirando(Michi)

	Kira := new(perro)
	AnimalRespirando(Kira)
	Kira.vivo = true

	fmt.Printf("Estoy vivo = %t", estoyVivo(Kira))
}
