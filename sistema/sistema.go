package sistema

import "time"

type SistemaSIGOA struct {
	aeronaves   Aeronaves
	aeropuertos Aeropuertos
	clientes    Clientes
	edificios   Edificios

	cargas   []*Carga
	vuelos   []*Vuelo
	reservas []*Reserva
}

// obtiene la hora del primer vuelo para calcular el inicio de la simulacion
func (s *SistemaSIGOA) ObtenerHoraDelPrimerVuelo() time.Time {
	primerVuelo := s.vuelos[0].ObtenerFechaHora()
	for _, vuelo := range s.vuelos[1:] {
		if vuelo.ObtenerFechaHora().Before(primerVuelo) {
			primerVuelo = vuelo.ObtenerFechaHora()
		}
	}
	return primerVuelo
}
func (s *SistemaSIGOA) ObtenerHoraDelUltimoVuelo() time.Time {
	ultimoVuelo := s.vuelos[0].ObtenerFechaHora()
	for _, vuelo := range s.vuelos[1:] {
		if vuelo.ObtenerFechaHora().After(ultimoVuelo) {
			ultimoVuelo = vuelo.ObtenerFechaHora()
		}
	}
	return ultimoVuelo
}
func NewSistemaSIGOA() *SistemaSIGOA {
	return &SistemaSIGOA{}
}

func (s *SistemaSIGOA) Aeronaves() *Aeronaves {
	return &s.aeronaves
}

func (s *SistemaSIGOA) Aeropuertos() *Aeropuertos {
	return &s.aeropuertos
}

func (s *SistemaSIGOA) Clientes() *Clientes {
	return &s.clientes
}

func (s *SistemaSIGOA) Edificios() *Edificios {
	return &s.edificios
}
