package sistema

import (
	"tdas/lista"
	"time"
)

type SistemaVuelos interface {
	// Guarda en el sistema los vuelos contenidos en el array, si un vuelo ya se encontraba cargado, lo actualiza.
	AgregarVuelos(vuelos []Vuelo)

	//Devuelve una lista con la información de 'k' vuelos en orden (por fecha)
	//ascendente o descendentemente según el 'modo' pasado por parámetro.
	VerTablero(k int, modo, desde, hasta string) (lista.Lista[string], error)

	//Devuelve la información completa del vuelo asociado al código pasado por parámetro.
	//Si el código no está en el sistema devuelve un error.
	InfoVuelo(codigo string) (string, error)

	//Devuelve la información completa del primer vuelo posterior a la fecha pasada por
	//parámetro desde 'origen' a 'destino'. Si no fue posible encontrar un vuelo, retorna
	//un mensaje indicándolo.
	SiguienteVuelo(origen, destino, fecha string) (string, error)

	//Devuelve los k vuelos con mayor prioridad, siendo 1 la prioridad más baja.
	//Los devuelve en una lista de strings en formato 'prioridad - codigo'
	// Si dos vuelos tienen la misma prioridad, se desempata por el código de
	//vuelo mostrándolos de menor a mayor (tomado como cadena).
	PrioridadVuelos(k int) (lista.Lista[string], error)

	//Borra todos los vuelos que estén dentro del intervalo 'desde' - 'hasta'
	//Devuelve una lista de strings con la información de los vuelos borrados.
	Borrar(desde, hasta string) (lista.Lista[string], error)
}

type Vuelo interface {
	//Devuelve el código asociado al vuelo.
	ObtenerCodigo() string

	//Devuelve la fecha del vuelo.
	ObtenerFecha() time.Time

	//Devuelve la prioridad del vuelo.
	ObtenerPrioridad() int

	//Devuelve toda la información del vuelo en forma de string.
	ObtenerInfoGeneral() string

	//Devuelve el la ciudad 'destino' del vuelo.
	ObtenerDestino() string

	//Devuelve la ciudad 'origen' del vuelo
	ObtenerOrigen() string
}
