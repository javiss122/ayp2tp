package simulador

type Mostrador struct {
	cola       []*Persona
	habilitado bool
}

func NewMostrador(habilitado bool) *Mostrador {
	return &Mostrador{
		cola:       make([]*Persona, 0),
		habilitado: habilitado,
	}
}

// HacerCola asigna a una persona a la cola del mostrador
func (m *Mostrador) HacerCola(persona *Persona) {
	if persona == nil {
		return
	}
	persona.CambiarEstado(EnEspera)
	m.cola = append(m.cola, persona)
}

// Tick ejecuta un tick del mostrador, procesando las personas en la cola
func (m *Mostrador) Tick() {
	if !m.habilitado || len(m.cola) == 0 {
		return
	}

	// Atender a la primera persona en la cola
	persona := m.cola[0]
	m.cola = m.cola[1:]

	persona.CambiarEstado(SiendoAtendido)
	persona.CambiarEstado(Despachando)
	persona.CambiarEstado(Retirandose)
}

func (m *Mostrador) Habilitado() bool {
	return m.habilitado
}

func (m *Mostrador) Deshabilitar(persona *Persona) {
	m.habilitado = false
}
