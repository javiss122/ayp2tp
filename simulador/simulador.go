package simulador

import (
	"fmt"
	"math/rand"
	"sigoa/sistema"
	"time"
)

type Simulador struct {
	sigoa       *sistema.SistemaSIGOA
	personas    []*Persona
	mostradores []*Mostrador
	hora        time.Time // Hora dentro de la simulación
}

func NewSimulador(sigoa *sistema.SistemaSIGOA) *Simulador {
	return &Simulador{
		sigoa: sigoa,
	}
}

func (s *Simulador) Iniciar() {
	// 1. Obtener la hora del primer vuelo y restar 3 horas
	primeraHora := s.sigoa.ObtenerHoraDelPrimerVuelo()
	//ultimaHora := s.sigoa.ObtenerHoraDelUltimoVuelo()
	s.hora = primeraHora.Add(-3 * time.Hour)

	// 2. (Opcional) Inicializar personas y mostradores aquí
	rand.Seed(time.Now().UnixNano())
	//tinicio := primeraHora.Add(-3 * time.Hour)

	// 3. Imprimir información de arranque
	fmt.Println("────────────────────────────────────────────")
	fmt.Println("🚀 Se ha iniciado la simulación SIGOA")
	fmt.Println("⏰ Hora de comienzo de la simulacion:", s.hora.Format("2006-01-02 15:04"))
	fmt.Println("✈️ Vuelos programados:", len(s.sigoa.Vuelos()))
	fmt.Println("📦 Cargas registradas:", len(s.sigoa.Cargas()))
	fmt.Println("👤 Reservas registradas:", len(s.sigoa.Reservas()))
	fmt.Println("────────────────────────────────────────────")
}

// Devuelve true si aún hay vuelos por realizar
func (s *Simulador) DebeContinuar() bool {
	return false
}

// Ejecuta un tick de la simulación
func (s *Simulador) Tick() {
	// TODO: procesar
	s.hora = s.hora.Add(time.Minute)
}
