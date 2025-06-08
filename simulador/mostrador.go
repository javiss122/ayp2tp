package simulador

type Mostrador struct {
	// TODO: cola queue.Queue[*Persona]
	habilitado bool
}

func NewMostrador(habilitado bool) *Mostrador {
	return &Mostrador{
		habilitado: habilitado,
	}
}

// HacerCola asigna a una persona a la cola del mostrador
func (m *Mostrador) HacerCola(persona *Persona) {

}

// Tick ejecuta un tick del mostrador, procesando las personas en la cola
func (m *Mostrador) Tick() {

}

func (m *Mostrador) Habilitado() bool {
	return m.habilitado
}

func (m *Mostrador) Deshabilitar(persona *Persona) {
	m.habilitado = false
}
