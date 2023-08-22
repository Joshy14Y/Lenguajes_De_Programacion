package main

import (
	"fmt"
	"sort"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
	for i, prod := range *l {
		if strings.ToUpper(nombre) == strings.ToUpper(prod.nombre) {
			(*l)[i].cantidad++
			if precio != (*l)[i].precio {
				(*l)[i].precio = precio
			}
			return
		}
	}
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
func (l *listaProductos) agregarProductos(prods ...producto) {
	for _, prod := range prods {
		l.agregarProducto(prod.nombre, prod.precio, prod.cantidad)
	}
}

// modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
// tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
// que el uso de la nueva función ameriten
func (l *listaProductos) buscarProducto(nombre string) (*producto, int, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	for i, prod := range *l {
		if strings.ToUpper(prod.nombre) == strings.ToUpper(nombre) {
			return &(*l)[i], i, 0
		}
	}
	return nil, 0, -1
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	prod, i, _ := l.buscarProducto(nombre)
	if prod != nil {
		if prod.cantidad >= cant {
			prod.cantidad -= cant
			if prod.cantidad == 0 {
				*l = append((*l)[:i], (*l)[i+1])
				fmt.Println("Se eliminó el producto", prod.nombre)
			}
		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
		//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	}
}

// haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE
func (l *listaProductos) cambiarPrecio(nombre string, precio int) {
	prod, _, _ := l.buscarProducto(nombre)
	prod.precio = precio
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var min listaProductos = nil
	for _, prod := range *l {
		if prod.cantidad <= 10 {
			min = append(min, prod)
		}
	}
	return min
}

// Práctica 2:

// a.
func (l *listaProductos) aumentarInventarioDeMínimos() {
	min := l.listarProductosMínimos()
	for i := range *l {
		for x, prod := range min {
			if (*l)[i].nombre == prod.nombre {
				z := existenciaMinima - (*l)[i].cantidad
				(*l)[i].cantidad = (*l)[i].cantidad + z
				min[x] = min[len(min)-1] // Reemplazar el producto procesado por el último de min
				min = min[:len(min)-1]   // Reducir la longitud de min
				break
			}
		}
	}
}

// b.
func ordenarPorNombre(lp *listaProductos) {
	sort.Slice(*lp, func(i, j int) bool {
		return (*lp)[i].nombre < (*lp)[j].nombre
	})
}

func ordenarPorPrecio(lp *listaProductos) {
	sort.Slice(*lp, func(i, j int) bool {
		return (*lp)[i].precio < (*lp)[j].precio
	})
}

func ordenarPorCantidad(lp *listaProductos) {
	sort.Slice(*lp, func(i, j int) bool {
		return (*lp)[i].cantidad < (*lp)[j].cantidad
	})
}

func main() {
	llenarDatos()

	// Ejercicio 1
	fmt.Println("\nOriginal:")
	fmt.Println(lProductos)
	fmt.Println("\nCon inventario aumentado:")
	lProductos.aumentarInventarioDeMínimos()
	fmt.Println(lProductos)

	// Ejercicio 2
	fmt.Println("\nOriginal:")
	fmt.Println(lProductos)

	ordenarPorNombre(&lProductos)
	fmt.Println("\nOrdenados por nombre:")
	fmt.Println(lProductos)

	ordenarPorPrecio(&lProductos)
	fmt.Println("\nOrdenados por precio:")
	fmt.Println(lProductos)

	ordenarPorCantidad(&lProductos)
	fmt.Println("\nOrdenados por cantidad:")
	fmt.Println(lProductos)
}
