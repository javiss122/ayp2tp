package sistema

type Edificio struct {
	posInicial uint // Xi
	altura     uint
	posFinal   uint // Xf
}

func NewEdificio(inicial, altura, final uint) *Edificio {
	return &Edificio{
		posInicial: inicial,
		altura:     altura,
		posFinal:   final,
	}
}

// ObtenerPosInicial devuelve la posició inicial del edificio.
func (e *Edificio) ObtenerPosInicial() uint {
	return e.posInicial
}

// ObtenerPosFinal devuelve la posició final del edificio.
func (e *Edificio) ObtenerPosFinal() uint {
	return e.posFinal
}

// ObtenerAltura devuelve la altura del edificio.
func (e *Edificio) ObtenerAltura() uint {
	return e.altura
}

// ------------------------------------------

type Edificios struct {
	edificios []*Edificio
}

func (e *Edificios) ObtenerEdificios() []*Edificio {
	return e.edificios
}

func (e *Edificios) Agregar(edificio *Edificio) {
	e.edificios = append(e.edificios, edificio)
}
