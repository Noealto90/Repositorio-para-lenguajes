package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo del cual se deben tomar eventuales decisiones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// Buscar si el producto ya existe en la lista
	p, err := l.buscarProducto(nombre)
	if err == 0 { // Si el producto ya existe, incrementó la cantidad
		p.cantidad += cantidad
		// Si el precio es diferente, actualizarlo
		if p.precio != precio {
			p.precio = precio
		}
	} else { // Si el producto no existe, agregarlo a la lista
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// Función para insertar varios productos a la lista
func (l *listaProductos) agregarProductosVarios(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
		//Llama a la función agregarProducto para continuar con la lógica que se tenía
	}
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			return &((*l)[i]), 0 // Producto encontrado, devuelve el producto y un 0 que significa encontrado sin problemas
		}
	}
	return nil, 1 // Producto no encontrado, se devuelve nil y un 1 para indicar error
}

// Función para eliminar productos
func (l *listaProductos) eliminarProducto(nombre string) {
	for i, p := range *l {
		if p.nombre == nombre {
			// Eliminar el producto de la lista
			*l = append((*l)[:i], (*l)[i+1:]...)
			return
		}
	}
}

// Función para vender productos
func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod, err = l.buscarProducto(nombre)

	if err == 0 {
		//Verifica si hay suficientes existencias
		if (*prod).cantidad >= cant {
			//Resta la cantidad de existencias
			(*prod).cantidad -= cant
			// Verificar si la cantidad de productos es cero y eliminar el producto de la lista
			if (*prod).cantidad == 0 {
				// Eliminar el producto de la lista
				l.eliminarProducto(nombre)
				fmt.Println("Producto", nombre, "fue vendido en su totalidad y posteriormente eliminado de la lista.")
			}
		} else {
			fmt.Println("No hay suficiente cantidad de", nombre, "para vender.")
		}
	} else {
		fmt.Println("El producto", nombre, "no se encuentra en la lista.")
	}
}

// Método para modificar el precio de un producto a partir de su nombre
func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) {
	// Buscar el producto en la lista
	prod, err := l.buscarProducto(nombre)

	// Verificar si se encontró el producto
	if err == 0 {
		// Modificar el precio del producto encontrado
		prod.precio = nuevoPrecio
		fmt.Println("Precio del producto", nombre, "modificado a", nuevoPrecio)
	} else {
		// Notificar si el producto no se encuentra en la lista
		fmt.Println("El producto", nombre, "no se encuentra en la lista.")
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

	//Prueba para función agregar producto
	lProductos.agregarProducto("frijoles", 4, 2500)

	//Prueba función agregarProductosVarios
	lProductos.agregarProductosVarios(
		producto{nombre: "arroz", cantidad: 10, precio: 100},
		producto{nombre: "Producto2", cantidad: 20, precio: 200},
		// Se pueden agregar muchos más
	)

}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	newSlice := make(listaProductos, 0)

	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)

	//venda productos
	lProductos.venderProducto("arroz", 10)
	lProductos.venderProducto("frijoles", 8)

	//Prueba para la función modificarPrecio
	lProductos.modificarPrecio("arroz", 7500)

	fmt.Println(lProductos)

	pminimos := lProductos.listarProductosMinimos()
	fmt.Println(pminimos)
}
