package lista_test

import (
	"github.com/stretchr/testify/require"
	"strconv"
	TDALista "tdas/lista"
	"testing"
)

var STRINGS = []string{"Go", "Python", "Java", "C", "C++", "Rust", "Swift", "Kotlin", "JavaScript", "Ruby"}
var ENTEROS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var FLOATS = []float64{3.14, 2.71, 1.618, 0.577, 4.669, 1.732, 2.414, 0.693, 6.283, 9.81}

const VOLUMEN = 10000

func verificarListaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

func TestListaVacia(t *testing.T) {
	verificarListaVacia(t, TDALista.CrearListaEnlazada[int]())
	verificarListaVacia(t, TDALista.CrearListaEnlazada[string]())
	verificarListaVacia(t, TDALista.CrearListaEnlazada[float64]())
}

func verificarEnlistarVaciar[T comparable](t *testing.T, lista TDALista.Lista[T], datos []T) {
	for i := range datos {
		lista.InsertarPrimero(datos[i])
		require.False(t, lista.EstaVacia())
		require.Equal(t, datos[i], lista.VerPrimero())
		require.Equal(t, i+1, lista.Largo())
	}

	require.Equal(t, datos[0], lista.VerUltimo())

	largo := len(datos)
	for i := range datos {
		require.False(t, lista.EstaVacia())
		require.Equal(t, datos[largo-1-i], lista.VerPrimero())
		require.Equal(t, datos[largo-1-i], lista.BorrarPrimero())
		require.Equal(t, len(datos)-1-i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())
}

func TestEnlistarVaciar(t *testing.T) {
	verificarEnlistarVaciar(t, TDALista.CrearListaEnlazada[int](), ENTEROS)
	verificarEnlistarVaciar(t, TDALista.CrearListaEnlazada[string](), STRINGS)
	verificarEnlistarVaciar(t, TDALista.CrearListaEnlazada[float64](), FLOATS)
}

func verificarVolumenLista[T comparable](t *testing.T, lista TDALista.Lista[T], generarDato func(int) T) {
	for i := 0; i < VOLUMEN; i++ {
		dato := generarDato(i)
		lista.InsertarPrimero(dato)
		require.False(t, lista.EstaVacia())
	}

	for i := 0; i < VOLUMEN; i++ {
		dato := generarDato(VOLUMEN - 1 - i)
		require.False(t, lista.EstaVacia())
		require.Equal(t, dato, lista.VerPrimero())
		require.Equal(t, dato, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

func TestVolumen(t *testing.T) {
	verificarVolumenLista(t, TDALista.CrearListaEnlazada[int](), func(i int) int { return i })
	verificarVolumenLista(t, TDALista.CrearListaEnlazada[string](), func(i int) string { return strconv.Itoa(i) })
	verificarVolumenLista(t, TDALista.CrearListaEnlazada[float64](), func(i int) float64 { return float64(i) * 1.1 })
}

func verificarInsertarIterExterno[T comparable](t *testing.T, lista TDALista.Lista[T], datos []T) {
	iterador := lista.Iterador()
	iterador.Insertar(datos[0])
	require.Equal(t, datos[0], lista.VerPrimero())
	require.Equal(t, datos[0], lista.VerUltimo()) //se inserta como primero y ultimo

	iterador.Siguiente()
	iterador.Insertar(datos[1])
	require.Equal(t, datos[1], lista.VerUltimo()) //el segundo dato se inserta al final

	iterador = lista.Iterador()
	iterador.Siguiente()
	iterador.Insertar(datos[2])
	require.Equal(t, datos[2], iterador.VerActual()) //insertar al medio
	iterador.Siguiente()
	require.Equal(t, datos[1], iterador.VerActual()) //verificar que el segundo dato ingresado sigue siendo el último
}

func TestInsertarIterExterno(t *testing.T) {
	verificarInsertarIterExterno(t, TDALista.CrearListaEnlazada[int](), ENTEROS)
	verificarInsertarIterExterno(t, TDALista.CrearListaEnlazada[string](), STRINGS)
	verificarInsertarIterExterno(t, TDALista.CrearListaEnlazada[float64](), FLOATS)
}

func verificarBorradoInicio[T comparable](t *testing.T, lista TDALista.Lista[T], datos []T) {
	iterador := lista.Iterador()
	for i, elem := range datos {
		iterador.Insertar(elem)
		iterador.Siguiente()
		require.Equal(t, i+1, lista.Largo())
	}

	iterador = lista.Iterador()
	for i, elem := range datos {
		borrado := iterador.Borrar()
		require.Equal(t, elem, borrado)
		require.Equal(t, len(datos)-i-1, lista.Largo())
	}
}

func verificarBorradoMedio[T comparable](t *testing.T, lista TDALista.Lista[T]) {
	largoOriginal := lista.Largo()
	mitad := lista.Largo() / 2
	iterador := lista.Iterador()
	for i := 0; i < mitad; i++ {
		iterador.Siguiente()
	}
	borrado := iterador.Borrar()
	require.Equal(t, largoOriginal-1, lista.Largo())

	//Buscar si borrado se encuentra en la lista. Funciona cuando no hay elementos repetidos en la lista
	iterador = lista.Iterador()
	encontrado := false
	for iterador.HaySiguiente() {
		if iterador.VerActual() == borrado {
			encontrado = true
			break
		}
		iterador.Siguiente()
	}
	require.False(t, encontrado)
}

func verificarBorradoUltimo[T comparable](t *testing.T, lista TDALista.Lista[T]) {
	ultimo := lista.VerUltimo()
	iterador := lista.Iterador()
	largoOriginal := lista.Largo()

	i := 0
	for iterador.HaySiguiente() && i < lista.Largo()-1 {
		iterador.Siguiente()
		i++
	}
	require.Equal(t, ultimo, iterador.Borrar())
	require.Equal(t, largoOriginal-1, lista.Largo())
	require.False(t, iterador.HaySiguiente())
}

func TestBorrarIterExterno(t *testing.T) {
	verificarBorradoInicio(t, TDALista.CrearListaEnlazada[int](), ENTEROS)
	verificarBorradoInicio(t, TDALista.CrearListaEnlazada[string](), STRINGS)
	verificarBorradoInicio(t, TDALista.CrearListaEnlazada[float64](), FLOATS)

	listaInt := TDALista.CrearListaEnlazada[int]()
	listaStr := TDALista.CrearListaEnlazada[string]()
	listaFloat := TDALista.CrearListaEnlazada[float64]()

	for i := 0; i < len(ENTEROS); i++ {
		listaInt.InsertarPrimero(ENTEROS[i])
		listaStr.InsertarPrimero(STRINGS[i])
		listaFloat.InsertarPrimero(FLOATS[i])
	}
	verificarBorradoMedio(t, listaInt)
	verificarBorradoMedio(t, listaStr)
	verificarBorradoMedio(t, listaFloat)

	verificarBorradoUltimo(t, listaInt)
	verificarBorradoUltimo(t, listaStr)
	verificarBorradoUltimo(t, listaFloat)
}

func TestSumaIterInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	for _, elem := range ENTEROS {
		lista.InsertarUltimo(elem)
	}
	lista.Iterar(func(valor int) bool {
		suma += valor
		return true
	})
	require.Equal(t, 55, suma)
}

func TestCondicionCorteIterInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	for _, elem := range STRINGS {
		lista.InsertarUltimo(elem)
	}

	var ultimaLeida string
	lista.Iterar(func(palabra string) bool {
		ultimaLeida = palabra
		if []rune(palabra)[0] == 'R' {
			return false
		}
		return true
	})
	require.Equal(t, "Rust", ultimaLeida)
}

func TestIteradorPuntaAPunta(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for _, elem := range ENTEROS {
		lista.InsertarUltimo(elem)
	}

	iter := lista.Iterador()
	for i := 0; i < len(ENTEROS); i++ {
		require.True(t, iter.HaySiguiente())
		require.Equal(t, ENTEROS[i], iter.VerActual())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
}

func TestInsertarIntercalado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())

	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())

	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}
