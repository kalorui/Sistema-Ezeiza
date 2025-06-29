package diccionario_test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
	"time"
)

var NUMS_DESORDENADOS = []int{16, 9, 1, 5, 21, 17, 0, 23, 30}

func funcCmpInt(a, b int) int {
	return a - b
}
func TestDiccionarioVacioAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioClaveDefaultAbb(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](funcCmpInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardarAbb(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoAbb(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotchAbb(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	desordenados := []int{16, 9, 1, 5, 21, 17, 0, 23, 30}

	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, actual := range desordenados {
		dic.Guardar(actual, actual)
	}
	for _, actual := range desordenados {
		dic.Guardar(actual, 2*actual)
	}
	ok := true
	for _, actual := range desordenados {
		if !ok {
			break
		}
		ok = dic.Obtener(actual) == 2*actual
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrarAbb(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestConClavesNumericasAbb(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](funcCmpInt)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructsAbb(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearABB[avanzado, int](func(a, b avanzado) int {
		sumaA := a.w + a.x.b + a.y.b
		sumaB := b.w + b.x.b + b.y.b
		if sumaA > sumaB {
			return 1
		} else if sumaA < sumaB {
			return -1
		}
		return 0
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVaciaAbb(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloAbb(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestGuardarYBorrarTodos(t *testing.T) {
	t.Log("Esta prueba guarda y borra todos los elementos.")

	desordenados := []int{16, 9, 1, 5, 21, 17, 0, 23, 30}

	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, actual := range desordenados {
		dic.Guardar(actual, actual)
		require.True(t, dic.Pertenece(actual))
	}

	for _, actual := range desordenados {
		borrado := dic.Borrar(actual)
		require.False(t, dic.Pertenece(actual))
		require.Equal(t, actual, borrado)
	}
}

func buscarClave(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestGuardaryBorrarBalanceado(t *testing.T) {
	t.Log("Borra la ra[iz y asegura que se reemplace por el más chico de su subarbol derecho" +
		"Borra un subarbol y se asegura que ya no esté en el árbol")
	preOrden := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range preOrden {
		dic.Guardar(numero, numero)
	}
	borrado := dic.Borrar(8)
	require.False(t, dic.Pertenece(borrado))
	iter := dic.Iterador()
	for range 6 { // llega hasta la posicion de la raíz
		iter.Siguiente()
	}
	claveActual, _ := iter.VerActual()
	hasta := 6
	require.Equal(t, 7, claveActual)
	iter = dic.IteradorRango(nil, &hasta)
	for iter.HaySiguiente() { // borrs subarbol izquierdo
		claveActual, _ = iter.VerActual()
		dic.Borrar(claveActual)
		iter.Siguiente()
	}

	for i := 0; i < 7; i++ {
		require.False(t, dic.Pertenece(i))
	}

	for i := 9; i < 16; i++ {
		require.True(t, dic.Pertenece(i))
	}

}

func TestBorrarHoja(t *testing.T) {
	preOrden := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range preOrden {
		dic.Guardar(numero, numero)
	}
	dic.Borrar(1)
	iter := dic.Iterador()
	claveActual, _ := iter.VerActual()
	require.NotEqualValues(t, 1, claveActual)
	iter.Siguiente() //Actual estaba en el nodo que era padre de 1 y voy a su hijo derecho
	claveActual, _ = iter.VerActual()
	require.Equal(t, 3, claveActual)
	dic.Borrar(9)
	require.False(t, dic.Pertenece(9))
}

func TestIteradorInternoClavesAbb(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno en orden")
	clave1 := "Perro"
	clave2 := "Vaca"
	clave3 := "Gato"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	ordenado := []string{"Gato", "Perro", "Vaca"}
	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.EqualValues(t, cs[0], ordenado[0])
	require.EqualValues(t, cs[1], ordenado[1])
	require.EqualValues(t, cs[2], ordenado[2])
}

func TestIteradorInternoValoresAbb(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoRangos(t *testing.T) {
	t.Log("Valida que los datos sean recorridas por un rango dado correctamente (y una única vez) con el iterador interno")

	desordenados := []int{16, 9, 1, 5, 21, 17, 0, 23, 30}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range desordenados {
		dic.Guardar(numero, numero)
	}
	desde := 21
	hasta := 31
	multiplicacion := 1
	dic.IterarRango(&desde, &hasta, func(_, b int) bool {
		multiplicacion *= b
		return true
	})
	require.EqualValues(t, 14490, multiplicacion)
}

func TestIteradorInternoRangoConNil(t *testing.T) {
	t.Log("Valida que el iterador interno con rangos se comporte como un iterador normal cuando desde y hasta son nil")
	desordenados := []int{16, 9, 1, 5, 21, 17, 2, 23, 30}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range desordenados {
		dic.Guardar(numero, numero)
	}
	multiplicacion := 1
	dic.IterarRango(nil, nil, func(_, b int) bool {
		multiplicacion *= b
		return true
	})
	require.EqualValues(t, 354715200, multiplicacion)
}

func TestIteradorInternoConCortes(t *testing.T) {
	t.Log("Valida que funcione correctamente al momento de iterar hasta cierto punto con ambos iteradores")
	desordenados := []int{16, 9, 1, 5, 21, 17, 0, 23, 30}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range desordenados {
		dic.Guardar(numero, numero)
	}

	suma := 0
	dic.Iterar(func(_, dato int) bool {
		if suma == 31 { //Solo recorre el subarbol izquierdo
			return false
		}
		suma += dato
		return true
	})
	require.Equal(t, 31, suma)

	multiplicacion := 1
	desde := 16

	dic.IterarRango(&desde, nil, func(_, dato int) bool { // Recorre una parte del subarbol derecho
		if multiplicacion == 5712 {
			return false
		}
		multiplicacion *= dato
		return true
	})
	require.Equal(t, 5712, multiplicacion)
}

/*PRUEBAS DE VOLUMEN ITERADOR INTERNO -> INICIO*/

func generarClavesUnicas(cantidad int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	claves := make([]int, cantidad)
	for i := 0; i < cantidad; i++ {
		claves[i] = i
	}
	r.Shuffle(cantidad, func(i, j int) { claves[i], claves[j] = claves[j], claves[i] })
	return claves
}

func TestABBVolumen(t *testing.T) {
	t.Log("Prueba de volumen con inserción desordenada")
	dic := TDADiccionario.CrearABB[int, string](funcCmpInt)
	cantidad := 100000
	claves := generarClavesUnicas(cantidad)

	for _, clave := range claves {
		dic.Guardar(clave, strconv.Itoa(clave))
	}
	require.Equal(t, cantidad, dic.Cantidad())

	clavesOrdenadas := make([]int, cantidad)
	copy(clavesOrdenadas, claves)
	sort.Ints(clavesOrdenadas)

	visitados := make([]int, 0, cantidad)
	dic.Iterar(func(clave int, valor string) bool {
		visitados = append(visitados, clave)
		require.Equal(t, strconv.Itoa(clave), valor)
		return true
	})

	require.Equal(t, clavesOrdenadas, visitados)
}

func TestABBVolumenBorrar(t *testing.T) {
	t.Log("Prueba de volumen con inserción desordenada")
	dic := TDADiccionario.CrearABB[int, string](funcCmpInt)
	cantidad := 100000
	claves := generarClavesUnicas(cantidad)

	for _, clave := range claves {
		dic.Guardar(clave, strconv.Itoa(clave))
	}

	require.Equal(t, cantidad, dic.Cantidad())
	var esperadas = []int{}
	for _, clave := range claves {
		if clave%2 == 0 {
			dic.Borrar(clave)
		} else {
			esperadas = append(esperadas, clave)
		}
	}

	sort.Ints(esperadas)

	i := 0
	dic.Iterar(func(clave int, valor string) bool {
		require.Equal(t, esperadas[i], clave)
		i++
		return true
	})

	require.Equal(t, len(esperadas), dic.Cantidad())
	require.Equal(t, i, dic.Cantidad())
}

func TestABBVolumenIteradorRango(t *testing.T) {
	t.Log("Prueba de volumen con iterador en un rango.")
	dic := TDADiccionario.CrearABB[int, string](funcCmpInt)
	cantidad := 100000
	claves := generarClavesUnicas(cantidad)

	for _, clave := range claves {
		dic.Guardar(clave, strconv.Itoa(clave))
	}

	desde := cantidad * 3 / 10
	hasta := cantidad * 7 / 10

	visitados := make([]int, 0)
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		visitados = append(visitados, clave)
		require.GreaterOrEqual(t, clave, desde)
		require.LessOrEqual(t, clave, hasta)
		require.Equal(t, strconv.Itoa(clave), valor)
		return true
	})

	clavesEnRango := 0
	for _, clave := range claves {
		if clave >= desde && clave <= hasta {
			clavesEnRango++
		}
	}
	require.Equal(t, clavesEnRango, len(visitados))
}

/*PRUEBAS DE VOLUMEN ITERADOR INTERNO -> FIN*/

/*PRUEBAS DE ITERADOR EXTERNO -> INICIO*/

func TestIterarDiccionarioVacioAbb(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIteradorAbb(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(segundo, claves))
	require.EqualValues(t, valores[buscarClave(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarClave(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIteradorConRangos(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Abeja"
	clave5 := "Foca"

	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "bee"
	valor5 := "sonido de foca"
	claves := []string{clave1, clave2, clave3, clave4, clave5}
	valores := []string{valor1, valor2, valor3, valor4, valor5}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}
	desde := "Foca"
	hasta := "Perro"
	iter := dic.IteradorRango(&desde, &hasta)
	claveActual, datoActual := iter.VerActual()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, clave5, claveActual)
	require.Equal(t, valor5, datoActual)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	claveActual, datoActual = iter.VerActual()
	require.Equal(t, clave1, claveActual)
	require.Equal(t, valor1, datoActual)
}

func TestIteradorFueraDeRanfo(t *testing.T) {
	t.Log("Valida que un iterador con un rango 'inaceptable' para el ABB, no haga nada")
	desordenados := []int{16, 9, 1, 5, 21, 17, 0, 23, 30}
	dic := TDADiccionario.CrearABB[int, int](funcCmpInt)
	for _, numero := range desordenados {
		dic.Guardar(numero, numero)
	}
	desde := 31
	hasta := 47
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.Panics(t, iter.Siguiente)
}

/*PRUEBAS DE ITERADOR EXTERNO -> FIN*/
