package sistema

import (
	"fmt"
	"strconv"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/lista"
	"time"
)

const _MINIMO_PRIORIDAD_ = 1

type ClaveFecha struct {
	Fecha  time.Time
	Codigo string
}

func cmpClaveFechaAsc(a, b ClaveFecha) int {
	if a.Fecha.Before(b.Fecha) {
		return -1
	}
	if a.Fecha.After(b.Fecha) {
		return 1
	}
	if a.Codigo < b.Codigo {
		return -1
	}
	if a.Codigo > b.Codigo {
		return 1
	}
	return 0
}

func cmpClaveFechaDesc(a, b ClaveFecha) int {
	return cmpClaveFechaAsc(b, a)
}

type gestorVuelos struct {
	vuelosPorCodigo    diccionario.Diccionario[string, Vuelo]
	vuelosPorFechaAsc  diccionario.DiccionarioOrdenado[ClaveFecha, string]
	vuelosPorFechaDesc diccionario.DiccionarioOrdenado[ClaveFecha, string]
}

func CrearGestorVuelos() SistemaVuelos {
	return &gestorVuelos{
		vuelosPorCodigo:    diccionario.CrearHash[string, Vuelo](),
		vuelosPorFechaAsc:  diccionario.CrearABB[ClaveFecha, string](cmpClaveFechaAsc),
		vuelosPorFechaDesc: diccionario.CrearABB[ClaveFecha, string](cmpClaveFechaDesc),
	}
}

func (g *gestorVuelos) AgregarVuelos(vuelos []Vuelo) {
	for _, nuevoVuelo := range vuelos {
		claveNueva := ClaveFecha{Fecha: nuevoVuelo.ObtenerFecha(), Codigo: nuevoVuelo.ObtenerCodigo()}

		if g.vuelosPorCodigo.Pertenece(nuevoVuelo.ObtenerCodigo()) {
			vueloAntiguoPublico := g.vuelosPorCodigo.Obtener(nuevoVuelo.ObtenerCodigo())
			claveAntigua := ClaveFecha{Fecha: vueloAntiguoPublico.ObtenerFecha(), Codigo: vueloAntiguoPublico.ObtenerCodigo()}
			g.vuelosPorFechaAsc.Borrar(claveAntigua)
			g.vuelosPorFechaDesc.Borrar(claveAntigua)
		}

		g.vuelosPorCodigo.Guardar(nuevoVuelo.ObtenerCodigo(), nuevoVuelo)
		g.vuelosPorFechaAsc.Guardar(claveNueva, nuevoVuelo.ObtenerCodigo())
		g.vuelosPorFechaDesc.Guardar(claveNueva, nuevoVuelo.ObtenerCodigo())
	}
}

func (g *gestorVuelos) VerTablero(k int, modo, desdeStr, hastaStr string) (lista.Lista[string], error) {
	tablero := lista.CrearListaEnlazada[string]()
	if k <= 0 {
		return tablero, fmt.Errorf("cantidad de vuelos debe ser mayor a 0")
	}
	if modo != "asc" && modo != "desc" {
		return tablero, fmt.Errorf("modo de ordenamiento inválido")
	}

	desde, err := time.Parse(TIME_FORMAT, desdeStr)
	if err != nil {
		return tablero, fmt.Errorf("formato de fecha 'desde' inválido")
	}
	hasta, err := time.Parse(TIME_FORMAT, hastaStr)
	if err != nil {
		return tablero, fmt.Errorf("formato de fecha 'hasta' inválido")
	}
	if hasta.Before(desde) {
		return tablero, fmt.Errorf("la fecha 'hasta' no puede ser anterior a 'desde'")
	}
	claveDesde := ClaveFecha{Fecha: desde, Codigo: ""}
	claveHasta := ClaveFecha{Fecha: hasta, Codigo: strings.Repeat("~", 10)}
	iter := g.vuelosPorFechaAsc.IteradorRango(&claveDesde, &claveHasta)

	if modo == "desc" {
		claveDesde, claveHasta = claveHasta, claveDesde
		iter = g.vuelosPorFechaDesc.IteradorRango(&claveDesde, &claveHasta)
	}

	i := 0
	for iter.HaySiguiente() && i < k {
		clave, codigo := iter.VerActual()
		info := fmt.Sprintf("%s - %s", clave.Fecha.Format(TIME_FORMAT), codigo)
		tablero.InsertarUltimo(info)
		iter.Siguiente()
		i++
	}
	return tablero, nil
}

func (g *gestorVuelos) InfoVuelo(codigo string) (string, error) {
	var vacio string
	if !g.vuelosPorCodigo.Pertenece(codigo) {
		return vacio, fmt.Errorf("el vuelo no pertenece al sistema")
	}
	vueloPublico := g.vuelosPorCodigo.Obtener(codigo)

	return vueloPublico.ObtenerInfoGeneral(), nil
}

func (g *gestorVuelos) SiguienteVuelo(origen, destino, fechaStr string) (string, error) {
	var vacio string
	parseFecha, err := time.Parse(TIME_FORMAT, fechaStr)
	if err != nil {
		return vacio, fmt.Errorf("error al parsear fecha: %w", err)
	}
	fecha := ClaveFecha{Fecha: parseFecha, Codigo: ""}
	for it := g.vuelosPorFechaAsc.IteradorRango(&fecha, nil); it.HaySiguiente(); it.Siguiente() {
		_, codigo := it.VerActual()
		vuelo := g.vuelosPorCodigo.Obtener(codigo)

		if vuelo.ObtenerOrigen() == origen && vuelo.ObtenerDestino() == destino {
			return vuelo.ObtenerInfoGeneral(), nil
		}
	}
	vacio = fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", origen, destino, fecha.Fecha.Format(TIME_FORMAT))
	return vacio, nil
}

func (g *gestorVuelos) Borrar(desde, hasta string) (lista.Lista[string], error) {
	aBorrar := lista.CrearListaEnlazada[ClaveFecha]()
	borrados := lista.CrearListaEnlazada[string]()

	parseDesde, err := time.Parse(TIME_FORMAT, desde)
	parseHasta, err := time.Parse(TIME_FORMAT, hasta)
	if err != nil {
		return borrados, fmt.Errorf("error al parsear fecha: %w", err)
	}
	if parseHasta.Before(parseDesde) {
		return borrados, fmt.Errorf("rango de fechas invalido")
	}

	claveDesde := ClaveFecha{Fecha: parseDesde, Codigo: ""}
	claveHasta := ClaveFecha{Fecha: parseHasta, Codigo: strings.Repeat("~", 10)}
	g.vuelosPorFechaAsc.IterarRango(&claveDesde, &claveHasta, func(clave ClaveFecha, codigo string) bool {
		info, _ := g.InfoVuelo(codigo)
		borrados.InsertarUltimo(info)
		aBorrar.InsertarUltimo(clave)
		return true
	})

	aBorrar.Iterar(func(clave ClaveFecha) bool {
		g.vuelosPorFechaAsc.Borrar(clave)
		g.vuelosPorFechaDesc.Borrar(clave)
		g.vuelosPorCodigo.Borrar(clave.Codigo)
		return true
	})
	return borrados, nil
}

func cmpPrioridad(vuelo1, vuelo2 Vuelo) int {
	comparadorPrioridad := vuelo1.ObtenerPrioridad() - vuelo2.ObtenerPrioridad()
	if comparadorPrioridad != 0 {
		return comparadorPrioridad
	}
	return strings.Compare(vuelo2.ObtenerCodigo(), vuelo1.ObtenerCodigo())
}

func (g *gestorVuelos) PrioridadVuelos(k int) (lista.Lista[string], error) {
	prioritarios := lista.CrearListaEnlazada[string]()
	if k < _MINIMO_PRIORIDAD_ {
		return prioritarios, fmt.Errorf("rango de prioridades no valido")
	}
	vuelos := make([]Vuelo, g.vuelosPorCodigo.Cantidad())
	i := 0
	g.vuelosPorCodigo.Iterar(func(clave string, vuelo Vuelo) bool {
		vuelos[i] = vuelo
		i++
		return true
	})

	heap := cola_prioridad.CrearHeapArr[Vuelo](vuelos, cmpPrioridad)
	desencolados := 0

	for !heap.EstaVacia() && desencolados < k {
		vuelo := heap.Desencolar()
		info := []string{strconv.Itoa(vuelo.ObtenerPrioridad()), vuelo.ObtenerCodigo()}
		prioritarios.InsertarUltimo(strings.Join(info, " - "))
		desencolados++
	}
	return prioritarios, nil
}
