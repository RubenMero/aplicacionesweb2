package main

import (
	"errors"
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {

	var repo cafeteria.Repositorio = cafeteria.NewRepoMemoria()
	repo.GuardarCliente(cafeteria.Cliente{ID: "123456789", Nombre: "Jonaikel Holguin", Carrera: "Ingeniería en Sistemas", Saldo: 100.0})
	repo.GuardarCliente(cafeteria.Cliente{ID: "159748632", Nombre: "Nasly Holguin", Carrera: "Comunicacion", Saldo: 90.0})

	repo.GuardarProducto(cafeteria.Producto{ID: 101, Nombre: "Agua 500ml", Precio: 0.50, Stock: 120, Categoria: cafeteria.Categoria{ID: 1, Nombre: "Bebidas"}})
	repo.GuardarProducto(cafeteria.Producto{ID: 102, Nombre: "Cafe 100ml", Precio: 0.50, Stock: 150, Categoria: cafeteria.Categoria{ID: 2, Nombre: "Bebidas"}})
	repo.GuardarProducto(cafeteria.Producto{ID: 103, Nombre: "Tostada Mixta", Precio: 0.75, Stock: 100, Categoria: cafeteria.Categoria{ID: 3, Nombre: "Snacks"}})

	fmt.Println("Clientes registrados:")
	c, err := repo.ObtenerClientePorID("123456789")

	if err != nil {
		fmt.Println("Error al obtener clientes:", err)
	} else {
		fmt.Println("Cliente encontrado", c.Nombre)
	}

	fmt.Println("Cliente no existe")
	c, err = repo.ObtenerClientePorID("000000000")

	if err != nil {
		fmt.Println("Error al obtener clientes:", err)

		if errors.Is(err, cafeteria.ErrCategoriaNoEncontrada) {
			fmt.Println("El cliente no existe")
		}
	} else {
		fmt.Println("Cliente encontrado", c.Nombre)
	}

	fmt.Println("Productos registrados:")
	p, err := repo.ObtenerProductoPorID(101)
	if err != nil {
		fmt.Println("Error al obtener productos:", err)
	} else {
		fmt.Println("Producto encontrado", p.Nombre)
	}
	fmt.Println("Producto no existe")
	p, err = repo.ObtenerProductoPorID(999)
	if err != nil {
		fmt.Println("Error al obtener productos:", err)
		if errors.Is(err, cafeteria.ErrProductoNoEncontrado) {
			fmt.Println("El producto no existe")
		}
	} else {
		fmt.Println("Producto encontrado", p.Nombre)
	}

	fmt.Println("Pedidos registrados:")
	Cliente, _ := repo.ObtenerClientePorID("123456789")
	Producto, _ := repo.ObtenerProductoPorID(101)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  Cliente,
		Producto: Producto,
		Cantidad: 2,
		Total:    Producto.Precio * 2,
		Fecha:    "2024-06-01",
	}
	err = repo.RegistrarPedido(pedido)
	if err != nil {
		fmt.Println("Error al registrar pedido:", err)
	} else {
		fmt.Println("Pedido registrado con éxito")
	}

	fmt.Println("Pedidos creado")
	fmt.Println(pedido)

	//c := cafeteria.Cliente{
	//	ID:      "123456789",
	//	Nombre:  "Jonaikel Holguin",
	//	Carrera: "Ingeniería en Sistemas",
	//	Saldo:   100.0,
	//}
	//fmt.Println("Cliente de prueba")
	//fmt.Println(c)
}
