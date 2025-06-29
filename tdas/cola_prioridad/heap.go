package cola_prioridad

const _CAP_INICIAL_ = 16
const _FACTOR_REDIMENSION_ = 2
const _MINIMO_REDUCCION_ = 4
const _MENSAJE_COLA_VACIA_ = "La cola esta vacia"

type funcCmp[T any] func(T, T) int

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func indiceHijoIzq(i int) int {
	return 2*i + 1
}

func indiceHijoDer(i int) int {
	return 2*i + 2
}

func indicePadre(i int) int {
	return (i - 1) / 2
}

func maximo[T any](cmp funcCmp[T], arr []T, indiceA, indiceB int) int {
	if cmp(arr[indiceA], arr[indiceB]) > 0 {
		return indiceA
	}
	return indiceB
}

func CrearHeap[T any](funcion_cmp funcCmp[T]) ColaPrioridad[T] {
	datos := make([]T, _CAP_INICIAL_)
	return &colaConPrioridad[T]{datos: datos, cant: 0, cmp: funcion_cmp}
}

func upheap[T any](arreglo []T, i, tamanio int, funcion_cmp funcCmp[T]) {
	if i == 0 {
		return
	}
	padre := indicePadre(i)
	if maximo(funcion_cmp, arreglo, i, padre) == padre {
		return
	}
	arreglo[padre], arreglo[i] = arreglo[i], arreglo[padre]
	upheap[T](arreglo, padre, tamanio, funcion_cmp)
}

func downheap[T any](arreglo []T, i, tamanio int, funcion_cmp funcCmp[T]) {
	hIzq := indiceHijoIzq(i)
	hDer := indiceHijoDer(i)
	if hIzq >= tamanio {
		return
	}

	hijoMayor := hIzq
	if hDer < tamanio {
		hijoMayor = maximo(funcion_cmp, arreglo, hIzq, hDer)
	}

	if maximo(funcion_cmp, arreglo, i, hijoMayor) == i {
		return
	}

	arreglo[i], arreglo[hijoMayor] = arreglo[hijoMayor], arreglo[i]
	downheap[T](arreglo, hijoMayor, tamanio, funcion_cmp)
}
func heapify[T any](arreglo []T, funcion_cmp funcCmp[T]) {
	if len(arreglo) == 0 {
		return
	}
	tamanio := len(arreglo)
	for i := tamanio/2 - 1; i >= 0; i-- {
		downheap(arreglo, i, tamanio, funcion_cmp)
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp funcCmp[T]) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(funcion_cmp)
	}
	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	heapify(datos, funcion_cmp)
	return &colaConPrioridad[T]{datos: datos, cant: len(arreglo), cmp: funcion_cmp}
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) redimensionar(nuevaCap int) {
	vectorDatos := make([]T, nuevaCap)
	copy(vectorDatos, cola.datos)
	cola.datos = vectorDatos
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	if cap(cola.datos) == cola.cant {
		cola.redimensionar(cap(cola.datos) * _FACTOR_REDIMENSION_)
	}
	cola.datos[cola.cant] = dato
	cola.cant++
	upheap(cola.datos, cola.cant-1, cola.cant, cola.cmp)
}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_COLA_VACIA_)
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(_MENSAJE_COLA_VACIA_)
	}
	borrado := cola.datos[0]
	cola.datos[0], cola.datos[cola.cant-1] = cola.datos[cola.cant-1], cola.datos[0]
	cola.cant--
	downheap(cola.datos, 0, cola.cant, cola.cmp)

	if cap(cola.datos)/_MINIMO_REDUCCION_ >= cola.cant && cap(cola.datos)/_FACTOR_REDIMENSION_ > _CAP_INICIAL_ {
		cola.redimensionar(cap(cola.datos) / _FACTOR_REDIMENSION_)
	}
	return borrado
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func HeapSort[T any](elementos []T, funcion_cmp funcCmp[T]) {
	tamanio := len(elementos)
	heapify(elementos, funcion_cmp)

	for i := tamanio - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, 0, i, funcion_cmp)
	}
}
