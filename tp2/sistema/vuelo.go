package sistema

import (
	"fmt"
	"strconv"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02T15:04:05"
)

type datosVuelo struct {
	codigo        string
	aerolinea     string
	origen        string
	destino       string
	tailNumber    string
	prioridad     int
	fecha         time.Time
	delay         int
	tiempoDeVuelo int
	cancelado     bool
}

func (v *datosVuelo) ObtenerCodigo() string {
	return v.codigo
}

func (v *datosVuelo) ObtenerFecha() time.Time {
	return v.fecha
}

func (v *datosVuelo) ObtenerPrioridad() int {
	return v.prioridad
}

func (v *datosVuelo) ObtenerInfoGeneral() string {
	cancelado := 0
	if v.cancelado {
		cancelado = 1
	}
	return fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d", v.codigo, v.aerolinea, v.origen, v.destino, v.tailNumber, v.prioridad, v.fecha.Format(TIME_FORMAT), v.delay, v.tiempoDeVuelo, cancelado)
}

func (v *datosVuelo) ObtenerDestino() string { return v.destino }

func (v *datosVuelo) ObtenerOrigen() string { return v.origen }

func convertirDato(a string) (int, error) {
	dato, err := strconv.Atoi(a)
	if err != nil {
		return -1, fmt.Errorf("error al parsear: %w", err)
	}
	return dato, nil
}

func CrearVuelo(data []string) (Vuelo, error) {
	if len(data) != 10 {
		return nil, fmt.Errorf("registro de vuelo con formato incorrecto")
	}
	prioridad, err := convertirDato(data[5])

	fecha, err := time.Parse(TIME_FORMAT, data[6])
	if err != nil {
		return nil, fmt.Errorf("error al parsear fecha: %w", err)
	}

	delay, err := convertirDato(data[7])
	tiempoVuelo, err := convertirDato(data[8])

	cancelado := data[9] == "1"

	return &datosVuelo{
		codigo:        data[0],
		aerolinea:     data[1],
		origen:        data[2],
		destino:       data[3],
		tailNumber:    data[4],
		prioridad:     prioridad,
		fecha:         fecha,
		delay:         delay,
		tiempoDeVuelo: tiempoVuelo,
		cancelado:     cancelado,
	}, nil
}
