package checkin

type TarjetaEmbarque struct {
	NumeroVuelo     string
	NumeroDocumento uint
	Categoria       string
	PosicionOrden   uint
	Zona            uint
}

type TicketEquipaje struct {
	NumeroVuelo       string
	NumeroDocumento   uint
	BultosDespachados uint
	PesoTotalEquipaje uint
}
