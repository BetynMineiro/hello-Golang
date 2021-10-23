package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const numeroMonitoramentos = 5
const delay = 10

func main() {
	showIntro()
	exibeNomes()
	for {
		showMenu()
		command := getCommand()
		setMenu(command)
	}

}

func showIntro() {
	nome, sobreNome, idade, _ := returnUserInfo()
	version := 1.1

	fmt.Println("Ola", nome, sobreNome, idade, "anos")
	fmt.Println("Version:", version)

	fmt.Println("O tipo da Variavel version é - ", reflect.TypeOf(version))
	fmt.Println("O endereco da Variavel version é - ", &version)

}
func returnUserInfo() (string, string, int, bool) {
	nome := "Carlos"
	var sobreNome string = "Alberto"
	var idade int = 35
	possuiFilhos := true

	return nome, sobreNome, idade, possuiFilhos
}
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
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

	switch command {
	case 1:
		{
			startMonitor()
		}
	case 2:
		fmt.Println("Exibindo Logs....")
		showLogs()
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido ")
		os.Exit(-1)
	}
}

func startMonitor() {
	fmt.Println("Iniciando Monitoramento....")
	fmt.Println("")
	for i := 1; i <= numeroMonitoramentos; i++ {
		sites := []string{"https://random-status-code.herokuapp.com/", "https://github.com/", "https://valor.globo.com/valor-data/carteira-valor/"}
		fmt.Println("Round:", i)
		fmt.Println("")
		fmt.Println("....... Utilizando For ...............")
		for i := 0; i < len(sites); i++ {
			verifyStatusSite(i, sites[i])
		}

		fmt.Println("")
		fmt.Println("....... Utilizando Range ...............")
		fmt.Println("")
		for i, site := range sites {
			verifyStatusSite(i, site)
		}

		fmt.Println("")
		fmt.Println("....... Utilizando leitura de arquivo ...............")
		fmt.Println("")

		sitesArquivo := readFile()
		for i, site := range sitesArquivo {
			verifyStatusSite(i, site)
		}
		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
}
func readFile() []string {
	var list []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	leitor := bufio.NewReader(file)
	for {
		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)
		list = append(list, line)
		if err == io.EOF {
			break
		}

	}
	file.Close()
	return list
}

func verifyStatusSite(position int, site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "na posição:", position, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "na posição:", position, "está com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}
func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online - " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))

}
