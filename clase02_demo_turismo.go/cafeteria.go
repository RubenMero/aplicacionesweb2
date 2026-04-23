package main

import (
	"errors"
	"fmt"
)

type Cliente struct {
	Id      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	Id        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	Id         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

// Clientes
func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarIndiceCliente(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.Id == id {
			return i
		}
	}
	return -1
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarIndiceCliente(clientes, id)
	if idx == -1 {
		fmt.Println("Cliente no encontrado")
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

func ListarClientes(clientes []Cliente) {
	fmt.Println("==============================================")
	fmt.Println("ID | NOMBRE            | CARRERA     | SALDO")
	fmt.Println("==============================================")

	if len(clientes) == 0 {
		fmt.Println("(no hay clientes)")
		return
	}

	for _, c := range clientes {
		fmt.Printf("%-2d | %-17s | %-11s | $%6.2f\n",
			c.Id, c.Nombre, c.Carrera, c.Saldo)
	}
	fmt.Println("==============================================")
}

// Productos
func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarIndiceProducto(productos []Producto, id int) int {
	for i, p := range productos {
		if p.Id == id {
			return i
		}
	}
	return -1
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarIndiceProducto(productos, id)
	if idx == -1 {
		fmt.Println("Producto no encontrado")
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func ListarProductos(productos []Producto) {
	fmt.Println("==============================================================")
	fmt.Println("ID | NOMBRE            | CATEGORIA  | PRECIO  | STOCK")
	fmt.Println("==============================================================")

	if len(productos) == 0 {
		fmt.Println("(no hay productos)")
		return
	}

	for _, p := range productos {
		fmt.Printf("%-2d | %-17s | %-10s | $%6.2f | %-5d\n",
			p.Id, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
}

// Punteros
func DescontarSaldo(cliente *Cliente, monto float64) error {
	if monto <= 0 {
		return errors.New("monto debe ser positivo")
	}
	if cliente.Saldo < monto {
		return errors.New("saldo insuficiente")
	}
	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("cantidad debe ser positiva")
	}
	if producto.Stock < cantidad {
		return errors.New("stock insuficiente")
	}
	producto.Stock -= cantidad
	return nil
}

// Pedidos
func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {

	idxCliente := BuscarIndiceCliente(clientes, clienteID)
	if idxCliente == -1 {
		return pedidos, errors.New("cliente no encontrado")
	}

	idxProducto := BuscarIndiceProducto(productos, productoID)
	if idxProducto == -1 {
		return pedidos, errors.New("producto no encontrado")
	}

	total := float64(cantidad) * productos[idxProducto].Precio

	err := DescontarStock(&productos[idxProducto], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxCliente], total)
	if err != nil {
		productos[idxProducto].Stock += cantidad
		return pedidos, err
	}

	nuevoPedido := Pedido{
		Id:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevoPedido)
	return pedidos, nil
}

func main() {

	clientes := []Cliente{
		{1, "Juan Perez", "Ingeniería", 100},
		{2, "Maria Gomez", "Medicina", 150},
		{3, "Carlos Sanchez", "Derecho", 200},
	}

	productos := []Producto{
		{1, "Café", 1.50, 10, "bebida"},
		{2, "Sandwich", 3.00, 5, "snack"},
		{3, "Jugo", 2.00, 8, "bebida"},
		{4, "Almuerzo", 5.00, 3, "comida"},
	}

	pedidos := []Pedido{}

	fmt.Println("\n--- CLIENTES ---")
	ListarClientes(clientes)

	fmt.Println("\n--- PRODUCTOS ---")
	ListarProductos(productos)

	fmt.Println("\n--- PEDIDO EXITOSO ---")
	var err error

	pedidos, err = RegistrarPedido(clientes, productos, pedidos, 1, 1, 2, "2024-06-01")
	if err != nil {
		fmt.Println("Error al registrar pedido:", err)
	} else {
		fmt.Println("Pedido registrado exitosamente")
	}

	ListarClientes(clientes)
	ListarProductos(productos)

	fmt.Println("\n--- ERROR: CLIENTE NO EXISTE ---")
	_, err = RegistrarPedido(clientes, productos, pedidos, 999, 1, 1, "2024-06-02")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- ERROR: PRODUCTO NO EXISTE ---")
	_, err = RegistrarPedido(clientes, productos, pedidos, 1, 999, 1, "2024-06-02")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- ERROR: STOCK INSUFICIENTE ---")
	_, err = RegistrarPedido(clientes, productos, pedidos, 1, 2, 1000, "2024-06-02")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- ERROR: SALDO INSUFICIENTE ---")
	clientes[0].Saldo = 1
	_, err = RegistrarPedido(clientes, productos, pedidos, 1, 4, 10, "2024-06-02")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
