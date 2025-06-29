package main

import (
	"bufio"
	"fmt"
	"os"
	comandos "tp2/comandos"
	sistema "tp2/sistema"
)

func main() {
	sistema := sistema.CrearGestorVuelos()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		err := comandos.EjecutarComando(sistema, linea)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error en comando", err)
		} else {
			fmt.Println("OK")
		}
	}
}
