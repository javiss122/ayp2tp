package main

import (
	"errors"
	"fmt"
	"sigoa/entrada"
	"sigoa/simulador"
	"sigoa/sistema"
)

func main() {
	fmt.Println("Iniciando sistema SIGOA...")
	sistema := sistema.NewSistemaSIGOA()
	fmt.Println("Cargando datos...")
	if err := entrada.CargarDatosEntrada(sistema); err != nil {
		panic(errors.New("No se ha podido inicializar el sistema SIGOA: " + err.Error()))
	}

	fmt.Println("Preparando simulador...")
	simulador := simulador.NewSimulador(sistema)
	simulador.Iniciar()

	fmt.Println("Iniciando simulación.")
	for simulador.DebeContinuar() {
		simulador.Tick()
	}
	fmt.Println("Simulación finalizada.")
}
