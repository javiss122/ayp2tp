package sistema

type Aeronave struct {
	matricula             string
	asientos              uint
	capacidadCarga        uint
	volumenCarga          float32
	configuracionAsientos []*ConfiguracionAsientos
}

func NewAeronave(matricula string, asientos, capacidad uint, volumen float32, configuracionAsientos []*ConfiguracionAsientos) *Aeronave {
	return &Aeronave{
		matricula:             matricula,
		asientos:              asientos,
		capacidadCarga:        capacidad,
		volumenCarga:          volumen,
		configuracionAsientos: configuracionAsientos,
	}
}

// ObtenerMatricula devuelve la matricula de la aeronave.
func (a *Aeronave) ObtenerMatricula() string {
	return a.matricula
}

// ObtenerAsientos devuelve la cantidad de asientos de la aeronave.
func (a *Aeronave) ObtenerAsientos() uint {
	return a.asientos
}

// ObtenerCapacidadCarga devuelve la capacidad de carga de la aeronave.
func (a *Aeronave) ObtenerCapacidadCarga() uint {
	return a.capacidadCarga
}

// ObtenerVolumenCarga devuelve el volumen de carga de la aeronave.
func (a *Aeronave) ObtenerVolumenCarga() float32 {
	return a.volumenCarga
}

// ObtenerConfiguracionAsientos devuelve la configuracion de asientos de la aeronave.
func (a *Aeronave) ObtenerConfiguracionAsientos() []*ConfiguracionAsientos {
	return a.configuracionAsientos
}

func (a *Aeronave) AgregarConfiguracion(config *ConfiguracionAsientos) {
	a.configuracionAsientos = append(a.configuracionAsientos, config)
}

// -----------------------------------------

type Aeronaves struct {
	aeronaves []*Aeronave
}

func (a *Aeronaves) ObtenerAeronaves() []*Aeronave {
	return a.aeronaves
}

func (a *Aeronaves) Agregar(aeronave *Aeronave) {
	a.aeronaves = append(a.aeronaves, aeronave)
}

func (a *Aeronaves) BuscarPorMatricula(matricula string) *Aeronave {
	for _, aeronave := range a.aeronaves {
		if aeronave.matricula == matricula {
			return aeronave
		}
	}
	return nil
}

// -------------------------------------

type ConfiguracionAsientos struct {
	zona           uint
	asientoInicial uint
	asientoFinal   uint
}

func NewConfiguracionAsientos(zona, inicial, final uint) *ConfiguracionAsientos {
	return &ConfiguracionAsientos{
		zona:           zona,
		asientoInicial: inicial,
		asientoFinal:   final,
	}
}

// ObtenerZona devuelve la zona de la configuracion de asientos.
func (c *ConfiguracionAsientos) ObtenerZona() uint {
	return c.zona
}

// ObtenerAsientoInicial devuelve el asiento inicial de la configuracion de asientos.
func (c *ConfiguracionAsientos) ObtenerAsientoInicial() uint {
	return c.asientoInicial
}

// ObtenerAsientoFinal devuelve el asiento final de la configuracion de asientos.
func (c *ConfiguracionAsientos) ObtenerAsientoFinal() uint {
	return c.asientoFinal
}
