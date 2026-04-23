package cafeteria

import "errors"

var (
	ErrclienteNoEncontrado   = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado  = errors.New("producto no encontrado")
	ErrStockInsuficiente     = errors.New("stock insuficiente")
	ErrSaldoInsuficiente     = errors.New("saldo insuficiente")
	ErrCategoriaNoEncontrada = errors.New("categoría no encontrada")
)

type Cliente struct {
	ID      string
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria Categoria
}

type Pedido struct {
	ID       int
	Cliente  Cliente
	Producto Producto
	Cantidad int
	Total    float64
	Fecha    string
}

type Categoria struct {
	ID     int
	Nombre string
}

type Repositorio interface {
	RegistrarCliente(c Cliente) error
	GuardarCliente(c Cliente) error
	ObtenerClientePorID(id string) (Cliente, error)
	ListarClientes() ([]Cliente, error)

	GuardarProducto(p Producto) error
	ObtenerProductoPorID(id int) (Producto, error)
	ListarProductos() ([]Producto, error)

	RegistrarPedido(p Pedido) error
	ObtenerPedidosPorClienteID(clienteID string) ([]Pedido, error)
	ListarPedidos() ([]Pedido, error)
}

type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{
		clientes:  []Cliente{},
		productos: []Producto{},
		pedidos:   []Pedido{},
	}
}

func (r *RepoMemoria) RegistrarCliente(c Cliente) error {
	for _, cliente := range r.clientes {
		if cliente.ID == c.ID {
			return errors.New("cliente ya registrado")
		}
	}
	r.clientes = append(r.clientes, c)
	return nil
}

func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	for i, cliente := range r.clientes {
		if cliente.ID == c.ID {
			r.clientes[i] = c
			return nil
		}
	}
	return ErrclienteNoEncontrado
}

func (r *RepoMemoria) ObtenerClientePorID(id string) (Cliente, error) {
	for _, cliente := range r.clientes {
		if cliente.ID == id {
			return cliente, nil
		}
	}
	return Cliente{}, ErrclienteNoEncontrado
}

func (r *RepoMemoria) ListarClientes() ([]Cliente, error) {
	return r.clientes, nil
}

func (r *RepoMemoria) GuardarProducto(p Producto) error {
	for i, producto := range r.productos {
		if producto.ID == p.ID {
			r.productos[i] = p
			return nil
		}
	}
	return ErrProductoNoEncontrado
}

func (r *RepoMemoria) ObtenerProductoPorID(id int) (Producto, error) {
	for _, producto := range r.productos {
		if producto.ID == id {
			return producto, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) ListarProductos() ([]Producto, error) {
	return r.productos, nil
}

func (r *RepoMemoria) RegistrarPedido(p Pedido) error {
	cliente, err := r.ObtenerClientePorID(p.Cliente.ID)
	if err != nil {
		return ErrclienteNoEncontrado
	}
	producto, err := r.ObtenerProductoPorID(p.Producto.ID)
	if err != nil {
		return ErrProductoNoEncontrado
	}
	if producto.Stock < p.Cantidad {
		return ErrStockInsuficiente
	}
	total := producto.Precio * float64(p.Cantidad)
	if cliente.Saldo < total {
		return ErrSaldoInsuficiente
	}
	cliente.Saldo -= total
	producto.Stock -= p.Cantidad
	p.Total = total
	r.pedidos = append(r.pedidos, p)
	return nil
}

func (r *RepoMemoria) ObtenerPedidosPorClienteID(clienteID string) ([]Pedido, error) {
	var pedidos []Pedido
	for _, pedido := range r.pedidos {
		if pedido.Cliente.ID == clienteID {
			pedidos = append(pedidos, pedido)
		}
	}
	return pedidos, nil
}

func (r *RepoMemoria) ListarPedidos() ([]Pedido, error) {
	return r.pedidos, nil
}

var _ Repositorio = (*RepoMemoria)(nil)
