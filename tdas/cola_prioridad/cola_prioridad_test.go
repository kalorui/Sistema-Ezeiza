package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"
	"time"
)

var STRINGS = []string{"Go", "Python", "Java", "C", "C++", "Rust", "Swift", "Kotlin", "JavaScript", "Ruby"}
var ENTEROS = []int{5, 1, 3, 9, 2, 8, 4, 7, 6, 0}
var FLOATS = []float64{3.14, 2.71, 1.618, 0.577, 4.669, 1.732, 2.414, 0.693, 6.283, 9.81}

const VOLUMEN = 100000

func cmpInt(a, b int) int { return a - b }
func cmpFloat(a, b float64) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
func cmpString(a, b string) int {
	return strings.Compare(a, b)
}

func verificarColaVacia[T any](t *testing.T, cola TDAColaPrioridad.ColaPrioridad[T]) {
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaVacia(t *testing.T) {
	verificarColaVacia(t, TDAColaPrioridad.CrearHeap[int](cmpInt))
	verificarColaVacia(t, TDAColaPrioridad.CrearHeap[string](cmpString))
	verificarColaVacia(t, TDAColaPrioridad.CrearHeap[float64](cmpFloat))
}

func verificarEncolarDesencolar[T comparable](t *testing.T, datos []T, cmp func(T, T) int) {
	cola := TDAColaPrioridad.CrearHeap[T](cmp)
	for _, val := range datos {
		cola.Encolar(val)
	}
	require.Equal(t, len(datos), cola.Cantidad())
	require.False(t, cola.EstaVacia())

	var anterior = cola.Desencolar()
	for i := 1; i < len(datos); i++ {
		actual := cola.Desencolar()
		require.GreaterOrEqual(t, cmp(anterior, actual), 0)
		anterior = actual
	}
	require.True(t, cola.EstaVacia())
}

func TestEncolarDesencolar(t *testing.T) {
	verificarEncolarDesencolar(t, ENTEROS, cmpInt)
	verificarEncolarDesencolar(t, STRINGS, cmpString)
	verificarEncolarDesencolar(t, FLOATS, cmpFloat)
}

func TestHeapifyDesdeArreglo(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeapArr[int](ENTEROS, cmpInt)
	require.Equal(t, len(ENTEROS), cola.Cantidad())
	require.Equal(t, 9, cola.VerMax())

	val := cola.Desencolar()
	require.Equal(t, 9, val)
	require.Equal(t, 8, cola.VerMax())
}

func TestVolumen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	datos := rand.Perm(VOLUMEN)
	cola := TDAColaPrioridad.CrearHeap[int](cmpInt)

	for _, v := range datos {
		cola.Encolar(v)
	}
	require.Equal(t, VOLUMEN, cola.Cantidad())

	for i := VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestVolumenCrearHeapArr(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	datos := rand.Perm(VOLUMEN)
	cola := TDAColaPrioridad.CrearHeapArr[int](datos, cmpInt)

	require.Equal(t, VOLUMEN, cola.Cantidad())

	for i := VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestHeapSortEnteros(t *testing.T) {
	datos := make([]int, len(ENTEROS))
	copy(datos, ENTEROS)
	TDAColaPrioridad.HeapSort(datos, cmpInt)
	for i := 1; i < len(datos); i++ {
		require.LessOrEqual(t, datos[i-1], datos[i])
	}
}

func TestHeapSortStrings(t *testing.T) {
	datos := make([]string, len(STRINGS))
	copy(datos, STRINGS)
	TDAColaPrioridad.HeapSort(datos, cmpString)
	for i := 1; i < len(datos); i++ {
		require.LessOrEqual(t, cmpString(datos[i-1], datos[i]), 0)
	}
}

func TestHeapSortFloats(t *testing.T) {
	datos := make([]float64, len(FLOATS))
	copy(datos, FLOATS)
	TDAColaPrioridad.HeapSort(datos, cmpFloat)
	for i := 1; i < len(datos); i++ {
		require.LessOrEqual(t, datos[i-1], datos[i])
	}
}

func TestHeapSortVolumen(t *testing.T) {
	datos := make([]int, VOLUMEN)
	for i := 0; i < VOLUMEN; i++ {
		datos[i] = VOLUMEN - i
	}
	TDAColaPrioridad.HeapSort(datos, cmpInt)
	for i := 1; i < VOLUMEN; i++ {
		require.LessOrEqual(t, datos[i-1], datos[i])
	}
}

func TestHeapPanics(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap[int](cmpInt)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestIntercalarEncolarDesencolar(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap[int](cmpInt)

	cola.Encolar(5)
	cola.Encolar(10)
	require.Equal(t, 10, cola.VerMax())

	val := cola.Desencolar()
	require.Equal(t, 10, val)

	cola.Encolar(7)
	require.Equal(t, 7, cola.VerMax())

	val = cola.Desencolar()
	require.Equal(t, 7, val)

	val = cola.Desencolar()
	require.Equal(t, 5, val)

	require.True(t, cola.EstaVacia())
}

func TestHeapDesdeArrayVacio(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeapArr[int]([]int{}, cmpInt)
	require.True(t, cola.EstaVacia())

	cola.Encolar(42)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 42, cola.VerMax())
	require.Equal(t, 42, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}
