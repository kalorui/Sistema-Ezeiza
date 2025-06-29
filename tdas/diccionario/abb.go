package diccionario

import (
	TDAPILA "tdas/pila"
)

const _MENSAJE_ERROR_CLAVE = "La clave no pertenece al diccionario"
const _MENSAJE_ERROR_ITERADOR = "El iterador termino de iterar"

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iteradorAbb[K comparable, V any] struct {
	diccionario  *abb[K, V]
	pila         TDAPILA.Pila[*nodoAbb[K, V]]
	desde, hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp, cantidad: 0}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.insertarNodo(&a.raiz, clave, dato)
}

func (a *abb[K, V]) insertarNodo(nodo **nodoAbb[K, V], clave K, dato V) {
	ptr, encontrado := a.buscarNodo(nodo, clave)
	if encontrado {
		(*ptr).dato = dato
	} else {
		*ptr = &nodoAbb[K, V]{clave: clave, dato: dato}
		a.cantidad++
	}
}

func (a *abb[K, V]) buscarNodo(nodo **nodoAbb[K, V], clave K) (**nodoAbb[K, V], bool) {
	if *nodo == nil {
		return nodo, false
	}
	comp := a.cmp(clave, (*nodo).clave)
	if comp == 0 {
		return nodo, true
	} else if comp < 0 {
		return a.buscarNodo(&(*nodo).izquierdo, clave)
	} else {
		return a.buscarNodo(&(*nodo).derecho, clave)
	}
}

func (a *abb[K, V]) borrarNodo(nodo **nodoAbb[K, V], clave K) V {
	ptr, encontrado := a.buscarNodo(nodo, clave)
	if !encontrado {
		panic(_MENSAJE_ERROR_CLAVE)
	}

	valor := (*ptr).dato

	if (*ptr).izquierdo == nil {
		*ptr = (*ptr).derecho
		return valor
	}

	if (*ptr).derecho == nil {
		*ptr = (*ptr).izquierdo
		return valor
	}

	minNodo := *(a.nodoMinimo(&(*ptr).derecho))

	(*ptr).clave = minNodo.clave
	(*ptr).dato = minNodo.dato

	a.borrarNodo(&(*ptr).derecho, minNodo.clave)

	return valor
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	_, encontrado := a.buscarNodo(&a.raiz, clave)
	return encontrado

}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, encontrado := a.buscarNodo(&a.raiz, clave)
	if !encontrado {
		panic(_MENSAJE_ERROR_CLAVE)
	}
	return (*nodo).dato
}

func (a *abb[K, V]) Borrar(clave K) V {
	valor := a.borrarNodo(&a.raiz, clave)
	a.cantidad--
	return valor
}

func (a *abb[K, V]) nodoMinimo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	actual := nodo
	for (*actual).izquierdo != nil {
		actual = &(*actual).izquierdo
	}
	return actual
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

//ITERADORES

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	var iterador iteradorAbb[K, V]
	iterador.diccionario = a
	iterador.pila = TDAPILA.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador.desde = desde
	iterador.hasta = hasta
	iterador.apilarIzquierdos(a.raiz)

	return &iterador
}

func (iter *iteradorAbb[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}

	cmpDesde := 0
	if iter.desde != nil {
		cmpDesde = iter.diccionario.cmp(nodo.clave, *iter.desde)
	}

	cmpHasta := 0
	if iter.hasta != nil {
		cmpHasta = iter.diccionario.cmp(nodo.clave, *iter.hasta)
	}

	if (iter.desde == nil || cmpDesde >= 0) && (iter.hasta == nil || cmpHasta <= 0) {
		iter.pila.Apilar(nodo)
		iter.apilarIzquierdos(nodo.izquierdo)
	} else if iter.desde != nil && cmpDesde < 0 {
		iter.apilarIzquierdos(nodo.derecho)
	} else if iter.hasta != nil && cmpHasta > 0 {
		iter.apilarIzquierdos(nodo.izquierdo)
	}
}

func estaEnRango[K comparable](cmp funcCmp[K], clave K, desde *K, hasta *K) bool {
	return (desde == nil || cmp(clave, *desde) >= 0) && (hasta == nil || cmp(clave, *hasta) <= 0)
}

func (iter *iteradorAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(_MENSAJE_ERROR_ITERADOR)
	}

	desapilado := iter.pila.Desapilar()
	iter.apilarIzquierdos(desapilado.derecho)
}

func (a *abb[K, V]) Iterar(visitar func(K, V) bool) {
	a.IterarRango(nil, nil, visitar)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if a.raiz == nil {
		return
	}
	a.raiz.iterarRango(desde, hasta, visitar, a.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp funcCmp[K]) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || cmp(nodo.clave, *desde) > 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	if estaEnRango(cmp, nodo.clave, desde, hasta) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || cmp(nodo.clave, *hasta) < 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	return true
}
