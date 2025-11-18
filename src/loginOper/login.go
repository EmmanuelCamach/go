package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func main() {

	const hPasswdBD string = "caf90169eefa5f807d577486b9f795ab86ae2983c5c20806cff959117e90af18"
	intentos := 1
	for {
		fmt.Println("\n*************************************************")
		fmt.Println("*         Login Sistemas Operativos             *")
		fmt.Println("*************************************************")

		fmt.Print("Intento No ", intentos, ".Digite Password:")
		//var passwd string
		//fmt.Scanf("%s", &passwd)
		passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("\nError al leer la contraseña:", err)
			return
		}
		fmt.Println()

		hashTemp := sha256.Sum256(passwd)
		hPasswdUser := fmt.Sprintf("%x", hashTemp)
		//fmt.Println("Hash de ", string(passwd), " es ", hPasswdUser)

		if hPasswdBD == hPasswdUser {
			fmt.Println("\t\t\tLogin Correcto! - OK")
			break
		} else {
			fmt.Println("\t\t\tLogin INCorrecto! X X X X X")
			LimpiarPantalla()
		}
		intentos++

	}
}

func LimpiarPantalla() {
	var cmd *exec.Cmd
	if os.PathSeparator == '\\' {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Linux/Mac
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
