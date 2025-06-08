package sistema

import (
	"fmt"
	"strings"
)

type Clientes struct {
	clientes []*Cliente
}

// ObtenerClientes devuelve la lista de clientes.
func (c *Clientes) ObtenerClientes() []*Cliente {
	return c.clientes
}

// Agregar agrega un cliente al sistema.
func (c *Clientes) Agregar(cliente *Cliente) {
	c.clientes = append(c.clientes, cliente)
}

// BuscarPorDocumento busca un cliente por documento.
func (c *Clientes) BuscarPorDocumento(documento uint) *Cliente {
	for _, cliente := range c.clientes {
		if cliente.ObtenerDocumento() == documento {
			return cliente
		}
	}
	return nil
}

// String devuelve una representación legible de los clientes.
func (c *Clientes) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Clientes (%d): [\n", len(c.clientes)))
	for _, cliente := range c.clientes {
		builder.WriteString("\t" + cliente.String() + "\n")
	}
	builder.WriteString("]")
	return builder.String()
}

// ----------------------------------------

type CategoriaCliente uint8

const (
	Platino CategoriaCliente = iota
	Oro
	Plata
	NoFrecuente
)

func (c CategoriaCliente) String() string {
	switch c {
	case Platino:
		return "Platino"
	case Oro:
		return "Oro"
	case Plata:
		return "Plata"
	default:
		return ""
	}
}

// ParseCategoriaCliente parsea una cadena de texto y devuelve la categoría correspondiente.
// Si no se reconoce la cadena se devuelve 0 y false.
func ParseCategoriaCliente(str string) (CategoriaCliente, bool) {
	switch str {
	case "Platino":
		return Platino, true
	case "Oro":
		return Oro, true
	case "Plata":
		return Plata, true
	case "":
		return NoFrecuente, true
	default:
		return 0, false
	}
}

// ----------------------------------------

type Cliente struct {
	documento uint // DNI o Pasaporte
	nombre    string
	apellido  string
	categoria CategoriaCliente
}

// NewCliente crea un nuevo cliente.
func NewCliente(documento uint, nombre, apellido string, categoria CategoriaCliente) *Cliente {
	return &Cliente{
		documento: documento,
		nombre:    nombre,
		apellido:  apellido,
		categoria: categoria,
	}
}

// ObtenerDocumento devuelve el DNI del cliente.
func (c *Cliente) ObtenerDocumento() uint {
	return c.documento
}

// ObtenerNombre devuelve el nombre del cliente.
func (c *Cliente) ObtenerNombre() string {
	return c.nombre
}

// ObtenerApellido devuelve el apellido del cliente.
func (c *Cliente) ObtenerApellido() string {
	return c.apellido
}

// ObtenerCategoria devuelve la categoría del cliente.
func (c *Cliente) ObtenerCategoria() CategoriaCliente {
	return c.categoria
}

// String devuelve una representación legible del cliente.
func (c *Cliente) String() string {
	return fmt.Sprintf(
		"Cliente{Documento=%d, Nombre=%s, Apellido=%s, Categoria=%s}",
		c.documento,
		c.nombre,
		c.apellido,
		c.categoria)
}
