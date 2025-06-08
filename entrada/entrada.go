package entrada

import (
	"bufio"
	"fmt"
	"os"
	"sigoa/sistema"
	"strings"
)

// Relativo al archivo main.go (cuando haces go run .)
const CarpetaData = "./data/"

// CargarArchivosEntrada carga los todos los archivos de entrada
// de la carpeta data
func CargarDatosEntrada(s *sistema.SistemaSIGOA) error {
	// Cargar Aeronaves
	err := CargarDatosAeronaves(s, CarpetaData+"aeronaves.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar las aeronaves: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d aeronaves\n", len(s.Aeronaves().ObtenerAeronaves()))

	// Cargar Aeropuertos
	err = CargarDatosAeropuertos(s, CarpetaData+"aeropuertos.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar los aeropuertos: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d aeropuertos\n", len(s.Aeropuertos().ObtenerAeropuertos()))

	// Cargar Clientes
	err = CargarDatosClientes(s, CarpetaData+"clientes.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar los clientes: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d clientes\n", len(s.Clientes().ObtenerClientes()))

	// Cargar Edificios
	err = CargarDatosEdificios(s, CarpetaData+"edificios.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar los edificios: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d edificios\n", len(s.Edificios().ObtenerEdificios()))

	// Cargar Cargas
	err = CargarDatosCargas(s, CarpetaData+"cargas.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar las cargas: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d cargas\n", len(s.Cargas()))

	// Cargar Vuelos
	err = CargarDatosVuelos(s, CarpetaData+"vuelos.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar los vuelos: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d vuelos\n", len(s.Vuelos()))

	// Cargar Configuración de Asientos
	err = CargarDatosConfiguracionesAsientos(s, CarpetaData+"configuracion_asientos.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar las configuración de asientos: %v", err)
	}
	fmt.Println("(i) Se han cargado las configuraciones de asientos")

	// Cargar Reservas
	err = CargarDatosReservas(s, CarpetaData+"reservas.txt")
	if err != nil {
		return fmt.Errorf("No se pudieron cargar las reservas: %v", err)
	}
	fmt.Printf("(i) Se han cargado %d reservas\n", len(s.Reservas()))

	return nil
}

type DatosEntrada struct {
	file        *os.File
	scanner     *bufio.Scanner
	columnNames *[]string
}

func NewDatosEntrada(csvPath string) (*DatosEntrada, error) {
	file, err := os.OpenFile(csvPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return &DatosEntrada{
		file:    file,
		scanner: scanner,
	}, nil
}

func (d *DatosEntrada) Next() map[string]string {
	if d.scanner.Scan() {
		data := strings.Split(d.scanner.Text(), ";")
		if d.columnNames == nil {
			d.columnNames = &data
			return d.Next()
		}

		result := make(map[string]string)
		for i, name := range *d.columnNames {
			result[name] = data[i]
		}
		return result
	} else {
		d.file.Close()
		return nil
	}
}
