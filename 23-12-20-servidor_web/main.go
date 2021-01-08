package main

import "net/http"

func main() {
	//HandleFunc significa que voy a manejar el ingreso en mi localhost y eso me va a ejecutar una funcion, recibe la ruta y la funcion a ejecutar
	http.HandleFunc("/", home)
	//escucha el puerto, y una vez que detecta que hay informacion en ese puerto, sirve una peticion, en este caso nil xq solo nos interesa que escuche el puerto 3000
	http.ListenAndServe(":3000", nil) //puerto 3000
}

//En esta funcion se reciben dos paremetros del subpaquete http, son los dos elementos que tienen que ver con el flujo de datos de un servidor cuando yo envio y recibo respuestas nuesrto puerto va a estar escuchando estas dos peticiones
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
