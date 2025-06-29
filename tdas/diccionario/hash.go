package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const TAMANIO_INICIAL = 17
const MENSAJE_CLAVE_NO_ENCONTRADA = "La clave no pertenece al diccionario"
const ERROR_HASHING = "Error al hashear"
const TAMANIO_CARGA_MAXIMO = 4.0
const TAMANIO_CARGA_MINIMO = 0.25
const FACTOR_REDIMENSION = 2

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}
type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iteradorHashAbierto[K comparable, V any] struct {
	diccionario *hashAbierto[K, V]
	posActual   int
	iterLista   TDALista.IteradorLista[parClaveValor[K, V]]
}

func crearTablaHash[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)

	for i := 0; i < tam; i++ {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{tabla: crearTablaHash[K, V](TAMANIO_INICIAL), tam: TAMANIO_INICIAL, cantidad: 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashearFnv[K comparable](elem K, tamTabla int) int {
	bytes := convertirABytes(elem)
	hasher := fnv.New64a()
	_, err := hasher.Write(bytes)
	if err != nil {
		panic(ERROR_HASHING)
	}
	return int(hasher.Sum64() % uint64(tamTabla))
}

func (diccionario *hashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func crearPar[K comparable, V any](clave K, dato V) parClaveValor[K, V] {
	return parClaveValor[K, V]{clave: clave, dato: dato}
}

func (diccionario *hashAbierto[K, V]) buscar(clave K) TDALista.IteradorLista[parClaveValor[K, V]] {
	pos := hashearFnv(clave, diccionario.tam)
	celdaLista := diccionario.tabla[pos]
	iterador := celdaLista.Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			break
		}
		iterador.Siguiente()
	}

	return iterador
}

func (diccionario *hashAbierto[K, V]) Pertenece(clave K) bool {
	iterLista := diccionario.buscar(clave)
	return iterLista.HaySiguiente()
}

func (diccionario *hashAbierto[K, V]) Obtener(clave K) V {
	iterLista := diccionario.buscar(clave)
	if !iterLista.HaySiguiente() {
		panic(MENSAJE_CLAVE_NO_ENCONTRADA)
	}
	return iterLista.VerActual().dato
}

func (diccionario *hashAbierto[K, V]) redimensionar(nuevaCap int) {
	antiguaTabla := diccionario.tabla
	diccionario.tabla = crearTablaHash[K, V](nuevaCap)
	diccionario.cantidad = 0
	diccionario.tam = nuevaCap
	for _, lista := range antiguaTabla {
		if lista.EstaVacia() {
			continue
		}
		lista.Iterar(func(par parClaveValor[K, V]) bool {
			diccionario.Guardar(par.clave, par.dato)
			return true
		})
	}
}

func (diccionario *hashAbierto[K, V]) hayQueAgrandar() bool {
	return float64(diccionario.cantidad)/float64(diccionario.tam) > TAMANIO_CARGA_MAXIMO
}

func (diccionario *hashAbierto[K, V]) hayQueAchicar() bool {
	cantidad := float64(diccionario.cantidad)
	tam := float64(diccionario.tam)
	return cantidad/tam < TAMANIO_CARGA_MINIMO && tam > TAMANIO_INICIAL
}

func (diccionario *hashAbierto[K, V]) Guardar(clave K, dato V) {

	iterLista := diccionario.buscar(clave)
	if iterLista.HaySiguiente() {
		iterLista.Borrar() // Actualizar el dato asociado a la clave
		diccionario.cantidad--
	}
	iterLista.Insertar(crearPar(clave, dato))
	diccionario.cantidad++

	if diccionario.hayQueAgrandar() {
		diccionario.redimensionar(diccionario.tam * FACTOR_REDIMENSION)
	}
}

func (diccionario *hashAbierto[K, V]) Borrar(clave K) V {
	iterLista := diccionario.buscar(clave)
	if !iterLista.HaySiguiente() {
		panic(MENSAJE_CLAVE_NO_ENCONTRADA)
	}
	datoBorrado := iterLista.VerActual().dato
	iterLista.Borrar()
	diccionario.cantidad--

	if diccionario.hayQueAchicar() {
		diccionario.redimensionar(diccionario.tam / FACTOR_REDIMENSION)
	}

	return datoBorrado
}

func (diccionario *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorHashAbierto[K, V]{diccionario: diccionario}
	iter.siguienteListaNoVacia()
	return iter
}

func (iter *iteradorHashAbierto[K, V]) siguienteListaNoVacia() {
	for iter.posActual < iter.diccionario.tam {
		lista := iter.diccionario.tabla[iter.posActual]
		if !lista.EstaVacia() {
			iter.iterLista = lista.Iterador()
			if iter.iterLista.HaySiguiente() {
				return
			}
		}
		iter.posActual++
	}
}

func (iter *iteradorHashAbierto[K, V]) HaySiguiente() bool {
	return iter.posActual < iter.diccionario.tam && iter.iterLista.HaySiguiente()
}

func (iter *iteradorHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	par := iter.iterLista.VerActual()
	return par.clave, par.dato
}

func (iter *iteradorHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.iterLista.Siguiente()
	if !iter.iterLista.HaySiguiente() {
		iter.posActual++
		iter.siguienteListaNoVacia()
	}
}

func (diccionario *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for i := 0; i < diccionario.tam; i++ {
		lista := diccionario.tabla[i]
		if lista.EstaVacia() {
			continue
		}
		continuar := true
		lista.Iterar(func(par parClaveValor[K, V]) bool {
			continuar = visitar(par.clave, par.dato)
			return continuar
		})
		if !continuar {
			return
		}
	}
}
