package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var comando string

	for {

		pwd, _ := os.Getwd()
		fmt.Print("oper@", pwd, "%")
		//fmt.Scanf("%s", &comando)
		lector := bufio.NewReader(os.Stdin)
		comando, _ = lector.ReadString('\n')
		comando = strings.TrimRight(comando, "\r\n")

		if comando == "bye" {
			break
		}

		sliceComando := strings.Fields(comando)
		if sliceComando[0] == "cd" {
			if len(sliceComando) < 2 {
				fmt.Println("Error: faltó la ruta")
				continue
			}
			err := os.Chdir(sliceComando[1])
			if err != nil {
				fmt.Println("Error al cambiar de directorio:", err)
			}
			continue
		}

		//fmt.Println("Ejecutando ", comando, "...")
		//res := exec.Command(sliceComando[0], sliceComando[1:]...)
		res := exec.Command("bash", "-c", comando)
		salida, err := res.CombinedOutput()
		if err != nil {
			fmt.Println("Error al ejecutar el comando:", err)
		}
		fmt.Println(string(salida))

	}

}
