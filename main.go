package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

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

const filePath = "clientes.json"

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

func saveClientesToFile() {
	file, err := json.MarshalIndent(clientes, "", "  ")
	if err != nil {
		log.Println("Erro ao codificar os dados para JSON:", err)
		return
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		log.Println("Erro ao salvar os dados no arquivo:", err)
	}
}

func loadClientesFromFile() {
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Println("Erro ao ler o arquivo clientes.json:", err)
			return
		}
		err = json.Unmarshal(file, &clientes)
		if err != nil {
			log.Println("Erro ao decodificar o JSON:", err)
			return
		}
		if len(clientes) > 0 {
			idCounter = clientes[len(clientes)-1].ID
		}
	} else {
		log.Println("Arquivo clientes.json não encontrado, iniciando com lista vazia.")
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

func getClientes(w http.ResponseWriter, r *http.Request) {
	log.Println("Função getClientes foi chamada")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

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

	// Definir as rotas antes dos arquivos estáticos
	r.HandleFunc("/clientes", getClientes).Methods("GET")
	r.HandleFunc("/clientes", addCliente).Methods("POST")
	r.HandleFunc("/clientes/{id}", getClientePorID).Methods("GET")
	r.HandleFunc("/clientes/{id}", updateCliente).Methods("PUT")
	r.HandleFunc("/clientes/{id}", deleteCliente).Methods("DELETE")

	// Aplicar o middleware de CORS
	r.Use(enableCors)

	// Servir arquivos estáticos
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./"))))

	log.Println("API rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
