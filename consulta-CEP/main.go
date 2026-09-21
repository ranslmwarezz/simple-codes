package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Address struct {
	Cep         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	Uf          string
	Estado      string
	Regiao      string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
	Erro        bool
}

func main() {

	var endereco Address
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scan(&cep)

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		log.Fatalf("erro ao realizar a leitura do body: %v", err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("falha ao fazer a requisição - StatusCode: %d", res.StatusCode)
	}

	if err := json.Unmarshal(body, &endereco); err != nil {
		log.Fatalf("erro ao fazer o unmarshal do json: %v", err)
	}

	if endereco.Erro {
		fmt.Printf("O CEP %s não segue o formato correto ou não existe", cep)
		return
	}

	fmt.Printf("{ Dados do CEP }\n")
	fmt.Printf("Logadouro: %s", endereco.Logradouro)
	fmt.Printf("\nBairro: %s", endereco.Bairro)
	fmt.Printf("\nLocalidade: %s", endereco.Localidade)
	fmt.Printf("\nEstado: %s", endereco.Estado)
	fmt.Printf("\nRegião: %s\n", endereco.Regiao)
}
