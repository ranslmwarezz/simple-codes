package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CNPJ struct {
	// Como o campo possui underscore lá na struct da api
	// preciso especificar como o json chegará aqui para que o unmarshal funcione
	RazaoSocial string `json:"razao_social"`
	NomeFantasia string `json:"nome_fantasia"`
	Cnpj string
	DescricaoSituacaoCadastral string `json:"descricao_situacao_cadastral"`
	DataInicioAtividade string `json:"data_inicio_atividade"`
	Logradouro string
	Numero string
	Bairro string
	Municipio string
	Uf string
	Cep string
}

type ErrorResponse struct {
    Name    string 
    Message string 
    Type    string
}

func main() {

	var cnpj CNPJ
	var erroResponse ErrorResponse
	var cnpjTyped string
	fmt.Print("Digite o CNPJ: ")
	fmt.Scan(&cnpjTyped)
	url := fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%s", cnpjTyped)

	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("falha na requisição: %v", err)
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		log.Fatalf("erro ao realizar a leitura do body: %v", err)
	}

	if res.StatusCode > 299 {
		fmt.Printf("{ falha na requisição - StatusCode: %d }\n", res.StatusCode)
		if err := json.Unmarshal(body, &erroResponse); err != nil {
			log.Fatalf("erro ao fazer o unmarshal do json: %v", err)
		}
		fmt.Printf("Messagem: %s", erroResponse.Message)
		fmt.Printf("\nTipo de erro: %s", erroResponse.Type)
		fmt.Printf("\nNome: %s\n", erroResponse.Name)
		return
	}

	if err := json.Unmarshal(body, &cnpj); err != nil {
		log.Fatalf("erro ao fazer o unmarshal do json: %v", err)
	}


	fmt.Printf("{ Dados do CNPJ }\n")
	fmt.Printf("Razão Social: %s", cnpj.RazaoSocial)
	fmt.Printf("\nNome Fantasia: %s", cnpj.NomeFantasia)
	fmt.Printf("\nCPNJ: %s", cnpj.Cnpj)
	fmt.Printf("\nSituação Cadastral: %s", cnpj.DescricaoSituacaoCadastral)
	fmt.Printf("\nInicio das Atividades: %s", cnpj.DataInicioAtividade)
	fmt.Printf("\nLogradouro: %s", cnpj.Logradouro)
	fmt.Printf("\nNúmero: %s", cnpj.Numero)
	fmt.Printf("\nBairro: %s", cnpj.Bairro)
	fmt.Printf("\nMunicípio: %s", cnpj.Municipio)
	fmt.Printf("\nUf: %s", cnpj.Uf)
	fmt.Printf("\nCEP: %s\n", cnpj.Cep)
}