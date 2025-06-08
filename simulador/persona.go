package simulador

import (
	"sigoa/simulador/checkin"
	"time"
)

type Estado uint8

const (
	Llegando       Estado = iota // Llegando al aeropuerto para hacer el check-in
	EnEspera                     // En espera en el mostrador de check-in o en la zona de embarque
	SiendoAtendido               // Siendo atendido en el mostrador de check-in
	Despachando                  // Despachando equipaje
	Embarcando                   // En la zona de embarque
	Subiendose                   // Subiendo al avión
	Despegando                   // Despegando
	Retirandose                  // No se pudo registrar o no logró embarcar
)

type Persona struct {
	Documento uint
	Equipaje  *Equipaje

	InicioAccion    time.Time
	Estado          Estado
	TarjetaEmbarque *checkin.TarjetaEmbarque
	TicketEquipaje  *checkin.TicketEquipaje
}

type Equipaje struct {
	Bultos    uint
	PesoTotal uint
}

func NewPersona(documento uint, equipaje *Equipaje) *Persona {
	return &Persona{
		Documento:       documento,
		Equipaje:        equipaje,
		Estado:          Llegando,
		TarjetaEmbarque: nil,
		TicketEquipaje:  nil,
	}
}
