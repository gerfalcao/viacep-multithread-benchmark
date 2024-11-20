package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getCep(url string, ch chan string) {
	req, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	ch <- string(res)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println(`Favor rodar o comando $ go run main.go <cep>`)
		os.Exit(1)
	}
	
	if len(os.Args) > 3 {
		fmt.Println("Informe somente um cep")
		os.Exit(1)
	}

	c1 := make(chan string)
	c2 := make(chan string)

	cep := os.Args[1]

	urlBrasilApi := "https://brasilapi.com.br/api/cep/v1/" + cep
	urlViaCep := "https://viacep.com.br/ws/" + cep + "/json/"

	go getCep(urlBrasilApi, c1)
	go getCep(urlViaCep, c2)

	select {
		case msg1 := <- c1:
			fmt.Println("Recebeu primeiro de Brasil Api. Dados: \n", msg1)
		
		case msg2 := <- c2:	
			fmt.Println("Recebeu primeiro de ViaCEP. Dados: \n", msg2)
		
		case <- time.After(time.Millisecond * 1000):
			println("Erro: timeout 1000ms")
	}

}