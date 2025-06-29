package diccionario

type DiccionarioOrdenado[K comparable, V any] interface {
	Diccionario[K, V]

	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
