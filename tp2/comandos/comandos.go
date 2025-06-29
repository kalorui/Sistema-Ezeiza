package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	sistema "tp2/sistema"
)

const _PARAMETROS_AGREGAR_ = 1
const _PARAMETROS_VER_TABLERO_ = 4
const _PARAMETROS_INFO_VUELO_ = 1
const _PARAMETROS_BORRAR_ = 2
const _PARAMETROS_PRIORIDAD_ = 1
const _PARAMETROS_SIGUIENTE_VUELO_ = 3

func EjecutarComando(sistema sistema.SistemaVuelos, linea string) error {
	partes := strings.Fields(linea)
	if len(partes) == 0 {
		return nil
	}

	comando := partes[0]
	args := partes[1:]
	var err error
	switch comando {
	case "agregar_archivo":

		err = manejarAgregarArchivo(sistema, args)
	case "ver_tablero":

		err = manejarVerTablero(sistema, args)
	case "info_vuelo":

		err = manejarInfoVuelo(sistema, args)
	case "prioridad_vuelos":

		err = manejarPrioridadVuelos(sistema, args)
	case "siguiente_vuelo":

		err = manejarSiguienteVuelo(sistema, args)
	case "borrar":
		err = manejarBorrar(sistema, args)
	default:
		return fmt.Errorf("%s", comando)
	}
	if err != nil {
		return fmt.Errorf("%s", comando)
	}
	return nil
}

func manejarAgregarArchivo(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_AGREGAR_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}

	ruta := args[0]

	archivo, err := os.Open(ruta)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo %s", ruta)
	}
	defer archivo.Close()

	var vuelos []sistema.Vuelo

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		if linea == "" {
			continue
		}
		campos := strings.Split(linea, ",")

		nuevoVuelo, err := sistema.CrearVuelo(campos)
		if err != nil {
			continue
		}
		vuelos = append(vuelos, nuevoVuelo)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	s.AgregarVuelos(vuelos)
	return nil
}
func manejarVerTablero(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_VER_TABLERO_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	k, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("el primer parámetro (K) debe ser un número")
	}
	tablero, err1 := s.VerTablero(k, args[1], args[2], args[3])
	if err1 != nil {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	tablero.Iterar(func(info string) bool {
		fmt.Printf("%s\n", info)
		return true
	})
	return nil
}

func manejarInfoVuelo(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_INFO_VUELO_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	info, err := s.InfoVuelo(args[0])
	if err != nil {
		return fmt.Errorf("error al manejar información del vuelo")
	}
	fmt.Printf("%s\n", info)
	return nil
}

func manejarPrioridadVuelos(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_PRIORIDAD_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	k, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("el parámetro debe ser un número")
	}
	prioritarios, err1 := s.PrioridadVuelos(k)
	if err1 != nil {
		return fmt.Errorf("error al manejar prioridad vuelos")
	}
	prioritarios.Iterar(func(info string) bool {
		fmt.Printf("%s\n", info)
		return true
	})
	return nil
}

func manejarSiguienteVuelo(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_SIGUIENTE_VUELO_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	info, err := s.SiguienteVuelo(args[0], args[1], args[2])
	if err != nil {
		return fmt.Errorf("error al manejar siguiente vuelo")
	}
	fmt.Printf("%s\n", info)
	return nil
}

func manejarBorrar(s sistema.SistemaVuelos, args []string) error {
	if len(args) != _PARAMETROS_BORRAR_ {
		return fmt.Errorf("cantidad de parámetros incorrecta")
	}
	borrados, err := s.Borrar(args[0], args[1])
	if err != nil {
		return fmt.Errorf("error al manejar Borrar")
	}
	borrados.Iterar(func(borrado string) bool {
		fmt.Printf("%s\n", borrado)
		return true
	})
	return nil
}
