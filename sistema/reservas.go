package sistema

import (
	"time"
)

type Reservas struct {
	reservas []*Reserva
}

// -----------------------------

type EstadoReserva uint8

const (
	Confirmada EstadoReserva = iota
	Cancelada
)

func ParseEstadoReserva(str string) (EstadoReserva, bool) {
	switch str {
	case "Confirmada":
		return Confirmada, true
	case "Cancelada":
		return Cancelada, true
	default:
		return 0, false
	}
}

func (e EstadoReserva) String() string {
	switch e {
	case Confirmada:
		return "Confirmada"
	default:
		return "Cancelada"
	}
}

type Reserva struct {
	codigo    uint
	pasajeros []*Cliente
	vuelo     *Vuelo
	fecha     time.Time
	estado    EstadoReserva
}

func NewReserva(codigo uint, pasajeros []*Cliente, vuelo *Vuelo, fecha time.Time, estado EstadoReserva) *Reserva {
	return &Reserva{
		codigo:    codigo,
		pasajeros: pasajeros,
		vuelo:     vuelo,
		fecha:     fecha,
		estado:    estado,
	}
}

// ObtenerCodigo devuelve el codigo de la reserva
func (r *Reserva) ObtenerCodigo() uint {
	return r.codigo
}

// ObtenerPasajeros devuelve los pasajeros de la reserva
func (r *Reserva) ObtenerPasajeros() []*Cliente {
	return r.pasajeros
}

// ObtenerVuelo devuelve el vuelo de la reserva
func (r *Reserva) ObtenerVuelo() *Vuelo {
	return r.vuelo
}

// ObtenerFecha devuelve la fecha de la reserva
func (r *Reserva) ObtenerFecha() time.Time {
	return r.fecha
}

// ObtenerEstado devuelve el estado de la reserva
func (r *Reserva) ObtenerEstado() EstadoReserva {
	return r.estado
}

func (r *Reserva) AgregarPasajero(pasajero *Cliente) {
	r.pasajeros = append(r.pasajeros, pasajero)
}

// -----------------------------------------

func (s *SistemaSIGOA) Reservas() []*Reserva {
	return s.reservas
}

// RegistrarReserva agrega una reserva al sistema
func (s *SistemaSIGOA) RegistrarReserva(reserva *Reserva) {
	if reserva != nil {
		s.reservas = append(s.reservas, reserva)
	}
}

// BuscarReservaPorPasajero busca una reserva por pasajero
func (s *SistemaSIGOA) BuscarReservaPorPasajero(documento uint) *Reserva {
	for _, reserva := range s.reservas {
		for _, pasajero := range reserva.ObtenerPasajeros() {
			if pasajero.ObtenerDocumento() == documento {
				return reserva
			}
		}
	}
	return nil
}

// BuscarReservasPorCodigo busca una reserva por codigo
func (s *SistemaSIGOA) BuscarReservasPorCodigo(codigo uint) *Reserva {
	for _, reserva := range s.reservas {
		if reserva.ObtenerCodigo() == codigo {
			return reserva
		}
	}
	return nil
}
