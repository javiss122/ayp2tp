package sistema

type Carga struct {
	destino *Aeropuerto // IATA
	peso    uint
	volumen float32
}

func NewCarga(destino *Aeropuerto, peso uint, volumen float32) *Carga {
	return &Carga{
		destino: destino,
		peso:    peso,
		volumen: volumen,
	}
}

// ObtenerPeso devuelve el peso de la carga.
func (c *Carga) ObtenerPeso() uint {
	return c.peso
}

// ObtenerVolumen devuelve el volumen de la carga.
func (c *Carga) ObtenerVolumen() float32 {
	return c.volumen
}

// ObtenerDestino devuelve el aeropuerto de destino
func (c *Carga) ObtenerDestino() *Aeropuerto {
	return c.destino
}

// -------------------------------------

func (s *SistemaSIGOA) Cargas() []*Carga {
	return s.cargas
}

func (s *SistemaSIGOA) RegistrarCarga(carga *Carga) {
	if carga != nil {
		s.cargas = append(s.cargas, carga)
	}
}
