package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"strconv"
	"io/ioutil"
	"os"
)

// Cliente representa a estrutura de um cliente
type Cliente struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Endereco  string `json:"endereco"`
	Telefone  string `json:"telefone"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
}

var clientes []Cliente
var idCounter int

// Middleware para habilitar CORS
func enableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            log.Println("Requisição OPTIONS recebida, enviando cabeçalhos CORS")
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Caminho para o arquivo JSON onde os dados serão armazenados
const filePath = "clientes.json"

// Função para salvar os clientes no arquivo JSON
func saveClientesToFile() {
    file, _ := json.MarshalIndent(clientes, "", "  ")
    _ = ioutil.WriteFile(filePath, file, 0644)
}

// Função para carregar os clientes do arquivo JSON
func loadClientesFromFile() {
    if _, err := os.Stat(filePath); err == nil {
        file, _ := ioutil.ReadFile(filePath)
        _ = json.Unmarshal(file, &clientes)
        if len(clientes) > 0 {
            idCounter = clientes[len(clientes)-1].ID
        }
    }
}

func init() {
    loadClientesFromFile()
    if len(clientes) == 0 {
        clientes = append(clientes, Cliente{
			ID:        1, 
			Nome:      "Usuário", 
			Sobrenome: "Exemplo", 
			Email:     "user@example.com", 
			Endereco:  "Rua Exemplo, 123",
			Telefone:  "99999-9999",
			Cidade:    "Cidade Exemplo",
			Estado:    "EX",
		})
        idCounter = 1
        saveClientesToFile()
    }
}

// Função para buscar todos os clientes
func getClientes(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição GET para /clientes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// Função para adicionar um novo cliente
func addCliente(w http.ResponseWriter, r *http.Request) {
    var cliente Cliente
    _ = json.NewDecoder(r.Body).Decode(&cliente)
    idCounter++
    cliente.ID = idCounter
    clientes = append(clientes, cliente)
    saveClientesToFile()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cliente)
}

// Função para atualizar um cliente existente
func updateCliente(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i, item := range clientes {
        if item.ID == id {
            clientes = append(clientes[:i], clientes[i+1:]...)
            var cliente Cliente
            _ = json.NewDecoder(r.Body).Decode(&cliente)
            cliente.ID = id
            clientes = append(clientes, cliente)
            saveClientesToFile()
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(cliente)
            return
        }
    }
    http.NotFound(w, r)
}

// Função para buscar um cliente específico pelo ID
func getClientePorID(w http.ResponseWriter, r *http.Request) {
	log.Println("Recebendo requisição GET para /clientes/{id}")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, cliente := range clientes {
		if cliente.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}
	http.NotFound(w, r)
}

// Função para deletar um cliente existente
func deleteCliente(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for i, item := range clientes {
        if item.ID == id {
            clientes = append(clientes[:i], clientes[i+1:]...)
            saveClientesToFile()
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]string{"message": "Cliente deletado"})
            return
        }
    }
    http.NotFound(w, r)
}

func main() {
	r := mux.NewRouter()

	// Aplicar o middleware de CORS
	r.Use(enableCors)

	// Adicionar uma rota genérica para capturar todas as requisições OPTIONS
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requisição OPTIONS recebida na rota geral")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})

	// Definir as rotas
	r.HandleFunc("/clientes", getClientes).Methods("GET")
	r.HandleFunc("/clientes", addCliente).Methods("POST")
	r.HandleFunc("/clientes/{id}", getClientePorID).Methods("GET")
	r.HandleFunc("/clientes/{id}", updateCliente).Methods("PUT")
	r.HandleFunc("/clientes/{id}", deleteCliente).Methods("DELETE")

	log.Println("API rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
