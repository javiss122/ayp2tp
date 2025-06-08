package sistema

type Aeropuertos struct {
	aeropuertos []*Aeropuerto
}

func (a *Aeropuertos) ObtenerAeropuertos() []*Aeropuerto {
	return a.aeropuertos
}

func (a *Aeropuertos) Agregar(aeropuerto *Aeropuerto) {
	a.aeropuertos = append(a.aeropuertos, aeropuerto)
}

func (a *Aeropuertos) BuscarPorIATA(codigo string) *Aeropuerto {
	for _, aeropuerto := range a.aeropuertos {
		if aeropuerto.ObtenerCodigoIATA() == codigo {
			return aeropuerto
		}
	}
	return nil
}

// ------------------------------------------

type Aeropuerto struct {
	codigoIATA string
	nombre     string
	ciudad     string
	provincia  string
}

func NewAeropuerto(codigo, nombre, ciudad, provincia string) *Aeropuerto {
	return &Aeropuerto{
		codigoIATA: codigo,
		nombre:     nombre,
		ciudad:     ciudad,
		provincia:  provincia,
	}
}

// ObtenerCodigoIATA devuelve el codigo IATA del aeropuerto.
func (a *Aeropuerto) ObtenerCodigoIATA() string {
	return a.codigoIATA
}

// ObtenerNombre devuelve el nombre del aeropuerto.
func (a *Aeropuerto) ObtenerNombre() string {
	return a.nombre
}

// ObtenerCiudad devuelve la ciudad del aeropuerto.
func (a *Aeropuerto) ObtenerCiudad() string {
	return a.ciudad
}

// ObtenerProvincia devuelve la provincia del aeropuerto.
func (a *Aeropuerto) ObtenerProvincia() string {
	return a.provincia
}
