package pila

const (
	MENSAJE_PILA_VACIA  = "La pila esta vacia"
	_FACTOR_CRECIMIENTO = 2
	_FACTOR_REDUCCION   = 2
	_CAPACIDAD_MINIMA   = 10
	_MINIMO_REDUCCION   = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _CAPACIDAD_MINIMA), cantidad: 0}
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	if nuevaCapacidad < _CAPACIDAD_MINIMA {
		if cap(pila.datos) == _CAPACIDAD_MINIMA {
			return
		}
		nuevaCapacidad = _CAPACIDAD_MINIMA
	}
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, pila.datos[:pila.cantidad])
	pila.datos = nuevosDatos
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(MENSAJE_PILA_VACIA)
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * _FACTOR_CRECIMIENTO)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(MENSAJE_PILA_VACIA)
	}
	pila.cantidad--
	elemento := pila.datos[pila.cantidad]

	if cap(pila.datos)/_MINIMO_REDUCCION >= pila.cantidad && cap(pila.datos)/_FACTOR_REDUCCION > _CAPACIDAD_MINIMA {
		pila.redimensionar(cap(pila.datos) / _FACTOR_REDUCCION)
	}
	return elemento
}
