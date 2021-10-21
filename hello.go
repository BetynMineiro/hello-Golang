package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	showIntro()
	showMenu()
	command := getCommand()
	setMenu(command)
}

func showIntro() {
	nome := "Carlos"
	var sobreNome string = "Alberto"
	var idade int = 35
	version := 1.1

	fmt.Println("Ola", nome, sobreNome, idade, "anos")
	fmt.Println("Version:", version)

	fmt.Println("O tipo da Variavel version é - ", reflect.TypeOf(version))
	fmt.Println("O endereco da Variavel version é - ", &version)

}

func showMenu() {
	fmt.Println("Menu:")
	fmt.Println("-------------------------------")
	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir Logs.")
	fmt.Println("0 - Sair do Programa.")
}

func getCommand() int {
	var commandInput int
	fmt.Scan(&commandInput)

	fmt.Println("Comando Escolhido foi: ", commandInput)
	return commandInput
}

func setMenu(command int) {

	// if commandInput == 1 {
	// 	fmt.Println("Iniciando Monitoramento....")
	// } else if commandInput == 2 {
	// 	fmt.Println("Exibindo Logs....")
	// } else if commandInput == 0 {
	// 	fmt.Println("Saindo do Programa...")
	// } else {
	// 	fmt.Println("Comando desconhecido ")
	// }

	switch command {
	case 1:
		{
			fmt.Println("Iniciando Monitoramento....")
		}
	case 2:
		fmt.Println("Exibindo Logs....")
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido ")
		os.Exit(-1)
	}
}
