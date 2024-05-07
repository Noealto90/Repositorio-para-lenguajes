package main

import (
	"fmt"
	"strings"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

// Crear la lista donde se almacenara la información de los clientes
type listaInfoClientes []infoCliente

var listaClientes listaInfoClientes

func filter2[P1 any](list []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func map2[P1, P2 any](list []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func reduce[T, U any](list []T, fn func(U, T) U, initial U) U {
	result := initial

	for _, v := range list {
		result = fn(result, v)
	}
	return result
}

//INICIO DE LOS EJERCICIOS

// PUNTO 1
// Función para agregarClientes.
func (lc *listaInfoClientes) agregarCliente(nombre, correo string, edad int32) {
	// Verificar si el cliente ya existe en la lista
	for _, cliente := range *lc {
		if cliente.nombre == nombre && cliente.correo == correo && cliente.edad == edad {
			fmt.Println("El cliente ya existe en la lista.")
			return
		}
	}

	// Si el cliente no está duplicado, agregarlo a la lista
	*lc = append(*lc, infoCliente{nombre: nombre, correo: correo, edad: edad})
}

// PUNTO 1
// Función cargarDatos
func cargarDatos() {
	// Agregar los clientes a la lista
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	listaClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.com", 47)
	listaClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)
}

// PUNTO 2
// Función para verificar si el correo contiene el apellido de la persona
func listaClientes_ApellidoEnCorreo(clientes *listaInfoClientes) listaInfoClientes {
	contieneApellidoEnCorreo := func(c infoCliente) bool {
		partesNombre := strings.Fields(c.nombre)
		apellidoCliente := partesNombre[len(partesNombre)-1]
		apellidoMin := strings.ToLower(apellidoCliente)
		correoMin := strings.ToLower(c.correo)
		return strings.Contains(correoMin, apellidoMin)
	}

	clientesFiltrados := filter2(*clientes, contieneApellidoEnCorreo)
	return clientesFiltrados
}

// PUNTO 3
// Función que devuelve la cantidad de clientes que tienen un correo que termina en .cr

func cantidadCorreosCostaRica(clientes []infoCliente) int {
	esDominioCR := func(c infoCliente) bool {
		correo := c.correo
		for i := len(correo) - 1; i >= 0; i-- {
			if correo[i] == '.' {
				return correo[i+1:] == "cr"
			}
		}
		return false
	}

	filtrados := filter2(clientes, esDominioCR)
	return reduce(filtrados, func(acc int, _ infoCliente) int {
		return acc + 1
	}, 0)
}

// PUNTO 4
// Función para recomendar un correo si no contiene su nombre y devuelva una nueva lista con los correos actualizados
func clientesSugerenciasCorreos(clientes []infoCliente) {
	contieneApellidoEnCorreo := func(c infoCliente) bool {
		partesNombre := strings.Fields(c.nombre)
		apellidoCliente := partesNombre[len(partesNombre)-1]
		apellidoMin := strings.ToLower(apellidoCliente)
		correoMin := strings.ToLower(c.correo)
		return strings.Contains(correoMin, apellidoMin)
	}

	cambiarCorreo := func(c infoCliente) infoCliente {
		partesNombre := strings.Fields(c.nombre)
		nombre := partesNombre[0]
		apellido := partesNombre[len(partesNombre)-1]
		c.correo = fmt.Sprintf("%s.%s@ejemplo.com", strings.ToLower(nombre), strings.ToLower(apellido))
		return c
	}

	clientesNoCumplen := filter2(clientes, func(c infoCliente) bool { return !contieneApellidoEnCorreo(c) })
	clientesModificados := map2(clientesNoCumplen, cambiarCorreo)

	for i, clienteModificado := range clientesModificados {
		for j, cliente := range clientes {
			if cliente == clientesNoCumplen[i] {
				clientes[j] = clienteModificado
			}
		}
	}
}

// PUNTO 5
// Función que ordena los correos en orden alfabético
func correosOrdenadosAlfabeticos(clientes []infoCliente) []string {
	// Extraer los correos electrónicos de los clientes
	obtenerCorreos := func(c infoCliente) string {
		return c.correo
	}

	// Obtener los correos electrónicos de los clientes
	correos := map2(clientes, obtenerCorreos)

	// Ordenar los correos electrónicos alfabéticamente
	ordenarCorreos := func(correos []string) []string {
		for i := 0; i < len(correos)-1; i++ {
			for j := i + 1; j < len(correos); j++ {
				if correos[i] > correos[j] {
					// Intercambiar correos electrónicos si están en el orden incorrecto
					correos[i], correos[j] = correos[j], correos[i]
				}
			}
		}
		return correos
	}

	// Filtrar y ordenar los correos electrónicos
	correosOrdenados := ordenarCorreos(correos)

	return correosOrdenados
}

func main() {

	//PUNTO 1
	fmt.Println("Punto 1: Muestra los datos insertados en la lista listaClientes")
	cargarDatos() //Las inserciones están dentro de esta función, porque sino aquí se hacía mucho.
	fmt.Println(listaClientes)

	//PUNTO 2
	fmt.Println()
	fmt.Println("Punto 2: Muestra a los clientes que tienen su apellido en el correo")
	// Llamar a la función listaClientes_ApellidoEnCorreo para filtrar los clientes
	clientesFiltrados := listaClientes_ApellidoEnCorreo(&listaClientes)
	// Imprimir los clientes filtrados por apellido en el correo
	fmt.Println(clientesFiltrados)

	//PUNTO 3
	fmt.Println()
	fmt.Println("Punto 3: Muestra la cantidad de correos que terminan con .cr")
	// Llamar a la función cantidadCorreosCostaRica para contar los clientes con correos de Costa Rica
	cantidad := cantidadCorreosCostaRica(listaClientes)
	// Imprimir la cantidad de clientes con correos de Costa Rica
	fmt.Println("Cantidad de clientes con correos de Costa Rica:", cantidad)

	//PUNTO 4
	fmt.Println()
	fmt.Println("Punto 4: Lista de listaClientes actualizada con los nuevos correos")
	// Llamar a la función listaClientes_ApellidoEnCorreo para filtrar los clientes
	clientesSugerenciasCorreos(listaClientes)
	// Imprimir los clientes filtrados por apellido en el correo
	fmt.Println(listaClientes)

	//PUNTO 5
	// Obtener la lista de correos ordenados alfabéticamente
	fmt.Println()
	fmt.Println("Punto 5: Nueva lista con los correos ordenados alfabéticamente")
	correosOrdenados := correosOrdenadosAlfabeticos(listaClientes)
	fmt.Println(correosOrdenados)

}
