package sistema

import (
	"time"
)

type Vuelo struct {
	numero         string
	fechaHora      time.Time   // DateTime
	destino        *Aeropuerto // IATA
	aeronaveCodigo *Aeronave   // Code
}

func NewVuelo(numero string, hora time.Time, destino *Aeropuerto, aeronave *Aeronave) *Vuelo {
	return &Vuelo{
		numero:         numero,
		fechaHora:      hora,
		destino:        destino,
		aeronaveCodigo: aeronave,
	}
}

// ObtenerNumero devuelve el numero del vuelo
func (v *Vuelo) ObtenerNumero() string {
	return v.numero
}

// ObtenerFechaHora devuelve la fecha y hora del vuelo
func (v *Vuelo) ObtenerFechaHora() time.Time {
	return v.fechaHora
}

// ObtenerDestino devuelve el aeropuerto de destino
func (v *Vuelo) ObtenerDestino() *Aeropuerto {
	return v.destino
}

// ObtenerAeronave devuelve el codigo de la aeronave
func (v *Vuelo) ObtenerAeronave() *Aeronave {
	return v.aeronaveCodigo
}

// -----------------------------------------

func (s *SistemaSIGOA) Vuelos() []*Vuelo {
	return s.vuelos
}

// ProgramarVuelo agrega un vuelo al sistema
func (s *SistemaSIGOA) ProgramarVuelo(vuelo *Vuelo) {
	if vuelo != nil {
		s.vuelos = append(s.vuelos, vuelo)
	}
}

func (s *SistemaSIGOA) BuscarVuelo(numero string) *Vuelo {
	for _, vuelo := range s.vuelos {
		if vuelo.ObtenerNumero() == numero {
			return vuelo
		}
	}
	return nil
}
