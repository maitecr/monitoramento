package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
	//	"reflect"
)

const monitoramentos = 3
const delay = 10

func main() { // função principal do programa, é por ela que o programa inicia

	lerArquivo()

	introducao()
	for {
		menu()

		comando := lerComando()
		//	fmt.Println("Tipo: ", reflect.TypeOf(comando))

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida")
			os.Exit(-1)
		}
	}
}

func introducao() {
	nome := "Gopher"
	//	var idade = 14
	var versao = 1.1
	fmt.Println("Olá, ", nome) //vírgula para concatenar
	fmt.Println("Esta é a versão", versao)
	//	fmt.Println("Tipo: ", reflect.TypeOf(nome))
}

func menu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando) //endereço da variável
	fmt.Println("Comando informado: ", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//sites := []string{"https://www.alura.com.br/", "https://www.google.com.br/"} //slice
	//sites = append(sites, "https://www.pucrs.br/")

	sites := lerArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Posição:", i, " / Site: ", site)

			testarSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
		fmt.Println(" ")
	}

	fmt.Println(" ")
}

func testarSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Println("Site: ", site, " carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, " está com problemas. Status Code: ", resposta.StatusCode)
	}
}

func lerArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(linha)

	return sites
}
