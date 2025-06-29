package pila_test

import (
	"strconv"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

var strings = []string{"Go", "Python", "Java", "C", "C++", "Rust", "Swift", "Kotlin", "JavaScript", "Ruby"}
var enteros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var floats = []float64{3.14, 2.71, 1.618, 0.577, 4.669, 1.732, 2.414, 0.693, 6.283, 9.81}

const LARGO = 10
const VOLUMEN = 10000

func TestPilaVacia(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaInt.EstaVacia())
	require.Panics(t, func() { pilaInt.VerTope() })
	require.Panics(t, func() { pilaInt.Desapilar() })

	pilaStr := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaStr.EstaVacia())
	require.Panics(t, func() { pilaStr.VerTope() })
	require.Panics(t, func() { pilaStr.Desapilar() })

	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pilaFloat.EstaVacia())
	require.Panics(t, func() { pilaFloat.VerTope() })
	require.Panics(t, func() { pilaFloat.Desapilar() })
}

func TestApilarDesapilar(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()

	for i := range LARGO {
		pilaInt.Apilar(enteros[i])
		require.Equal(t, enteros[i], pilaInt.VerTope())
		require.False(t, pilaInt.EstaVacia())

		pilaString.Apilar(strings[i])
		require.Equal(t, strings[i], pilaString.VerTope())
		require.False(t, pilaString.EstaVacia())

		pilaFloat.Apilar(floats[i])
		require.Equal(t, floats[i], pilaFloat.VerTope())
		require.False(t, pilaFloat.EstaVacia())
	}

	for i := range LARGO {
		require.False(t, pilaInt.EstaVacia())
		require.Equal(t, enteros[LARGO-1-i], pilaInt.VerTope())
		require.Equal(t, enteros[LARGO-1-i], pilaInt.Desapilar())

		require.False(t, pilaString.EstaVacia())
		require.Equal(t, strings[LARGO-1-i], pilaString.VerTope())
		require.Equal(t, strings[LARGO-1-i], pilaString.Desapilar())

		require.False(t, pilaFloat.EstaVacia())
		require.Equal(t, floats[LARGO-1-i], pilaFloat.VerTope())
		require.Equal(t, floats[LARGO-1-i], pilaFloat.Desapilar())
	}

	require.True(t, pilaInt.EstaVacia())
	require.True(t, pilaString.EstaVacia())
	require.True(t, pilaFloat.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()

	for i := range VOLUMEN {
		palabra := strconv.Itoa(i)
		decimal := float64(i) * 1.1

		pilaInt.Apilar(i)
		require.Equal(t, i, pilaInt.VerTope())
		require.False(t, pilaInt.EstaVacia())

		pilaString.Apilar(palabra)
		require.Equal(t, palabra, pilaString.VerTope())
		require.False(t, pilaString.EstaVacia())

		pilaFloat.Apilar(decimal)
		require.Equal(t, decimal, pilaFloat.VerTope())
		require.False(t, pilaFloat.EstaVacia())
	}

	for i := range VOLUMEN {
		j := VOLUMEN - 1 - i
		palabra := strconv.Itoa(j)
		decimal := float64(j) * 1.1

		require.False(t, pilaInt.EstaVacia())
		require.Equal(t, j, pilaInt.VerTope())
		require.Equal(t, j, pilaInt.Desapilar())

		require.False(t, pilaString.EstaVacia())
		require.Equal(t, palabra, pilaString.VerTope())
		require.Equal(t, palabra, pilaString.Desapilar())

		require.False(t, pilaFloat.EstaVacia())
		require.Equal(t, decimal, pilaFloat.VerTope())
		require.Equal(t, decimal, pilaFloat.Desapilar())
	}

	require.True(t, pilaInt.EstaVacia())
	require.True(t, pilaString.EstaVacia())
	require.True(t, pilaFloat.EstaVacia())

}

func TestCondicionesBorde(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()

	//Pila recién creada se comporta como debe
	require.True(t, pilaInt.EstaVacia())
	require.True(t, pilaString.EstaVacia())
	require.True(t, pilaFloat.EstaVacia())

	require.Panics(t, func() { pilaInt.VerTope() })
	require.Panics(t, func() { pilaString.VerTope() })
	require.Panics(t, func() { pilaFloat.VerTope() })

	require.Panics(t, func() { pilaInt.Desapilar() })
	require.Panics(t, func() { pilaString.Desapilar() })
	require.Panics(t, func() { pilaFloat.Desapilar() })

	//Apilamos y desapilamos para comprobar que al final se comporta como una pila vacía
	for i := range LARGO {
		pilaInt.Apilar(i)
		pilaString.Apilar(strings[i])
		pilaFloat.Apilar(floats[i])
	}

	for range LARGO {
		pilaInt.Desapilar()
		pilaString.Desapilar()
		pilaFloat.Desapilar()
	}

	require.True(t, pilaInt.EstaVacia())
	require.True(t, pilaString.EstaVacia())
	require.True(t, pilaFloat.EstaVacia())
}
