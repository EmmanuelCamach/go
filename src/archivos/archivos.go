package main

import (
	"fmt"
	"os"
	"strings"
)

var nomArchivo string = "proyOper.users"

//password debian - windows10 - redhat - ubuntu

func main() {

	for {
		var opcion int
		fmt.Println("****** OPERACIONES CON ARCHIVOS-OPERATIVOS ********")
		fmt.Println("1. Crear Archivo:", nomArchivo)
		fmt.Println("2. leer Archivo:", nomArchivo)
		fmt.Println("3. Adicionar usuario")
		fmt.Println("4. Salir")
		fmt.Print("\n\nDigite Opcion (1-4): ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			crearArchivo()
		case 2:
			leerArchivo()
		case 3:
			modifArchivo()
		case 4:
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func crearArchivo() {
	datos := []byte("luis:81d93757457f988523814ae0009837ae893f38d3fe123f2c37896f118b4c7804\n")
	err := os.WriteFile(nomArchivo, datos, 0666)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	fmt.Println("Archivo creado correctamente.")
}

func leerArchivo() {
	datos, err := os.ReadFile(nomArchivo)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	credenciales := strings.Split(strings.TrimSpace(string(datos)), "\n")
	for x, dato := range credenciales {
		fmt.Println("Usuario", x+1, "→", dato)
	}
}

func modifArchivo() {
	datos := "\njose:7d3b5c83009fadf734c06eeecd7fbe256c69f71c8ba0429e4d7ad5f54b2e4097"
	archivo, err := os.OpenFile(nomArchivo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(datos)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Usuario agregado correctamente.")
}
