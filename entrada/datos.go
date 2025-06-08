package entrada

import (
	"fmt"
	"sigoa/sistema"
	"strconv"
	"time"
)

func CargarDatosAeronaves(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}
	for data := entrada.Next(); data != nil; data = entrada.Next() {
		asientos, err := strconv.ParseUint(data["Asientos"], 10, 32)
		if err != nil {
			return err
		}
		capacidadCarga, err := strconv.ParseUint(data["CapacidadCarga"], 10, 32)
		if err != nil {
			return err
		}
		volumenCarga, err := strconv.ParseFloat(data["VolumenCarga"], 32)
		if err != nil {
			return err
		}
		aeronave := sistema.NewAeronave(
			data["Matricula"],
			uint(asientos),
			uint(capacidadCarga),
			float32(volumenCarga),
			make([]*sistema.ConfiguracionAsientos, 0, 3),
		)
		s.Aeronaves().Agregar(aeronave)
	}
	return nil
}

func CargarDatosAeropuertos(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}
	for data := entrada.Next(); data != nil; data = entrada.Next() {
		aeropuerto := sistema.NewAeropuerto(
			data["COD_IATA"],
			data["NOMBRE AEROPUERTO"],
			data["CIUDAD"],
			data["PROVINCIA"],
		)
		s.Aeropuertos().Agregar(aeropuerto)
	}
	return nil
}

func CargarDatosClientes(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}
	for data := entrada.Next(); data != nil; data = entrada.Next() {
		documento, err := strconv.ParseUint(data["DNI"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se pudo parsear el DNI %s", data["DNI"])
		}
		categoria, ok := sistema.ParseCategoriaCliente(data["Categoria"])
		if !ok {
			return fmt.Errorf("No se reconoce la categoría %s", data["Categoria"])
		}
		cliente := sistema.NewCliente(
			uint(documento), data["Nombre"], data["Apellido"], categoria,
		)
		s.Clientes().Agregar(cliente)
	}
	return nil
}

func CargarDatosEdificios(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}

	for data := entrada.Next(); data != nil; data = entrada.Next() {
		xiVal, err := strconv.ParseUint(data["xi"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se puedo parsear la posicion inicial %d", xiVal)
		}
		alturaVal, err := strconv.ParseUint(data["altura"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se puedo parsear la altura %d", alturaVal)
		}
		xfVal, err := strconv.ParseUint(data["xf"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se puedo parsear la posicion final %d", xfVal)
		}

		edificio := sistema.NewEdificio(
			uint(xiVal),
			uint(alturaVal),
			uint(xfVal),
		)
		s.Edificios().Agregar(edificio)
	}
	return nil
}

func CargarDatosCargas(s *sistema.SistemaSIGOA, path string) error {
	entradas, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}
	for data := entradas.Next(); data != nil; data = entradas.Next() {
		aeropuerto := s.Aeropuertos().BuscarPorIATA(data["Destino"])
		if aeropuerto == nil {
			return fmt.Errorf("No existe un aeropuerto %s", data["Destino"])
		}
		peso, err := strconv.ParseUint(data["Peso"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se pudo parsear el peso %s", data["Peso"])
		}
		volumen, err := strconv.ParseFloat(data["Volumen"], 32)
		if err != nil {
			return fmt.Errorf("No se pudo parsear el volumen %s", data["Volumen"])
		}
		carga := sistema.NewCarga(
			aeropuerto,
			uint(peso),
			float32(volumen),
		)
		s.RegistrarCarga(carga)
	}
	return nil
}

func CargarDatosVuelos(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}

	for data := entrada.Next(); data != nil; data = entrada.Next() {
		numero := data["numero"]
		fechaHoraStr := data["fecha_hora"]
		destinoIATA := data["destino"]
		aeronaveMatricula := data["aeronave"]

		fechaHora, err := time.Parse("2006-01-02 15:04:05", fechaHoraStr)
		if err != nil {
			return fmt.Errorf("Error parseando fecha/hora: %v", err)
		}

		destino := s.Aeropuertos().BuscarPorIATA(destinoIATA)
		if destino == nil {
			return fmt.Errorf("Aeropuerto no encontrado: %s", destinoIATA)
		}

		aeronave := s.Aeronaves().BuscarPorMatricula(aeronaveMatricula)
		if aeronave == nil {
			return fmt.Errorf("Aeronave no encontrada: %s", aeronaveMatricula)
		}

		vuelo := sistema.NewVuelo(
			numero,
			fechaHora,
			destino,
			aeronave,
		)
		s.ProgramarVuelo(vuelo)
	}

	return nil
}

func CargarDatosConfiguracionesAsientos(s *sistema.SistemaSIGOA, path string) error {
	entrada, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}

	for data := entrada.Next(); data != nil; data = entrada.Next() {
		aeronave := s.Aeronaves().BuscarPorMatricula(data["CódigoAeronave"])
		if aeronave == nil {
			return fmt.Errorf("No se reconoce el codigo de aeronave %s", data["CódigoAeronave"])
		}

		zona, err := strconv.ParseUint(data["Zona"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se reconoce la zona %s", data["Zona"])
		}
		inicial, err := strconv.ParseUint(data["AsientoInicial"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se reconoce el asiento inicial %s", data["AsientoInicial"])
		}
		final, err := strconv.ParseUint(data["AsientoFinal"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se reconoce el asiento final %s", data["AsientoFinal"])
		}

		config := sistema.NewConfiguracionAsientos(
			uint(zona),
			uint(inicial),
			uint(final),
		)
		aeronave.AgregarConfiguracion(config)

	}
	return nil
}

func CargarDatosReservas(s *sistema.SistemaSIGOA, path string) error {
	entradas, err := NewDatosEntrada(path)
	if err != nil {
		return err
	}
	for data := entradas.Next(); data != nil; data = entradas.Next() {
		codReserva, err := strconv.ParseUint(data["CodReserva"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se pudo parsear el codReserva %s", data["CodReserva"])
		}

		dniPasajero, err := strconv.ParseUint(data["DNIPasajero"], 10, 32)
		if err != nil {
			return fmt.Errorf("No se pudo parsear el DNIPasajero %s", data["DNIPasajero"])
		}
		cliente := s.Clientes().BuscarPorDocumento(uint(dniPasajero))
		if cliente == nil {
			return fmt.Errorf("No existe un cliente %s", data["DNIPasajero"])
		}

		vuelo := s.BuscarVuelo(data["NroVuelo"])
		if vuelo == nil {
			return fmt.Errorf("No existe un vuelo %s", data["NroVuelo"])
		}

		estado, ok := sistema.ParseEstadoReserva(data["EstadoReserva"])
		if !ok {
			return fmt.Errorf("No se reconoce el estado %s", data["EstadoReserva"])
		}

		reserva := s.BuscarReservasPorCodigo(uint(codReserva))
		if reserva == nil {
			reserva = sistema.NewReserva(
				uint(codReserva),
				make([]*sistema.Cliente, 0, 5),
				vuelo,
				time.Now(), // FIXME:
				estado)
			s.RegistrarReserva(reserva)
		}

		reserva.AgregarPasajero(cliente)
	}
	return nil
}
