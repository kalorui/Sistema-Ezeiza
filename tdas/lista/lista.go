package lista

type Lista[T any] interface {

	// EstaVacia devuelte verdadero si la lista no contiene elementos, en caso contrario devuelve falso.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al comienzo de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista. Si la lista tiene elementos, elimina el primero y lo devuelve.
	// Si esta vacia, entra en panico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el valor del primer elemento en la lista. Si la lista esta vacia entra en panico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el valor del ultimo elemento en la lista. Si la lista esta vacia entra en panico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recorre todos los elementos de la lista y ejecuta la función visitar en cada uno de ellos.
	// La función visitar recibe un elemento de tipo T y debe devolver un valor booleano.
	// Si la función devuelve verdadero, la iteración continúa; si devuelve falso, la iteración se detiene.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un nuevo iterador para recorrer la lista.
	// El iterador permite acceder a los elementos de la lista uno por uno mediante sus métodos
	// como `Siguiente` y `Borrar`.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento en la posicion actual del iterador.
	// Si se invoca sobre un iterador que ya haya iterado todos los elementos entra en panico con un mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve verdadero si el iterador tiene un elemento actual disponible, en caso contrario devuelve falso.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemeneto, en caso que ya se hayan iterado todos los elementos entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta un elemento nuevo en la posicion actual del iterador.
	// El nuevo elemento queda como actual después de la inserción.
	Insertar(T)

	// Borrar borra el elemento actual del iterador y devuelve su valor.
	// Si se invoca sobre un iterador que ya haya iterado todos los elementos entra en panico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
