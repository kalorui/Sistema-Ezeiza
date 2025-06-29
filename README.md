# Trabajo práctico de la materia Algoritmos y Estructuras de Datos, facultad de ingenieria de la UBA

## Descripción

El aeropuerto de Algueiza opera la entrada y salida de aviones; los operarios de dicho aeropuerto nos han pedido que implementemos un sistema de consulta de vuelos que les permita: ordenar, filtrar, analizar y obtener información de los distintos vuelos.

## Archivos

El sistema que actualmente tiene Algueiza genera continuamente archivos en formato .csv que contienen la información de cada vuelo. Nuestro sistema deberá procesar estos archivos y responder a las consultas que los operarios necesiten realizar. El sistema tiene que ser capaz de recibir en cualquier momento un nuevo archivo conteniendo el detalle de vuelos (nuevos o viejos) y actualizar sus datos.

## Interfaz

  agregar_archivo <nombre_archivo>: procesa de forma completa un archivo de .csv que contiene datos de vuelos.

  ver_tablero <K cantidad vuelos> <modo: asc/desc> <desde> <hasta>: muestra los K vuelos ordenados por fecha de forma ascendente (asc) o descendente (desc), cuya fecha de despegue esté dentro de el intervalo <desde> <hasta> (inclusive).

  info_vuelo <código vuelo>: muestra toda la información posible en sobre el vuelo que tiene el código pasado por parámetro.

  prioridad_vuelos <K cantidad vuelos>: muestra los códigos de los K vuelos que tienen mayor prioridad.

  siguiente_vuelo <aeropuerto origen> <aeropuerto destino> <fecha>: muestra la información del vuelo (tal cual en info_vuelo) del próximo vuelo directo que conecte los aeropuertos de origen y destino, a partir de la fecha indicada (inclusive). Si no hay un siguiente vuelo cargado, imprimir No hay vuelo registrado desde <aeropuerto origen> hacia <aeropuerto destino> desde <fecha> (con los valores que correspondan).

  borrar <desde> <hasta>: borra todos los vuelos cuya fecha de despegue estén dentro del intervalo <desde> <hasta> (inclusive).
