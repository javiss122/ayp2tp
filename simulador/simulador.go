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
	if len(s.sigoa.Vuelos()) == 0 {
		return false
	}

	ultima := s.sigoa.ObtenerHoraDelUltimoVuelo()
	return s.hora.Before(ultima.Add(1 * time.Hour))
}

// Ejecuta un tick de la simulación
func (s *Simulador) Tick() {
	// Generar llegada aleatoria de personas
	if rand.Intn(10) < 3 {
		p := NewPersona(uint(rand.Intn(1000000)), &Equipaje{Bultos: 1, PesoTotal: 20})
		s.personas = append(s.personas, p)
		if len(s.mostradores) == 0 {
			s.mostradores = append(s.mostradores, NewMostrador(true))
		}
		s.mostradores[rand.Intn(len(s.mostradores))].HacerCola(p)
	}

	// Actualizar mostradores
	for _, m := range s.mostradores {
		m.Tick()
	}

	s.hora = s.hora.Add(time.Minute)
}
