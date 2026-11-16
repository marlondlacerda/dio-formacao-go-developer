package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

var (
	clientes  = make(map[int]Cliente)
	proximoID = 1
	mutex     sync.RWMutex
)

func criarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente Cliente

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if cliente.Nome == "" || cliente.Telefone == "" || cliente.Email == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	cliente.ID = proximoID
	proximoID++
	clientes[cliente.ID] = cliente
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(cliente)
}

func listarClientes(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()

	resultado := make([]Cliente, 0, len(clientes))

	for _, cliente := range clientes {
		resultado = append(resultado, cliente)
	}

	mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resultado)
}

func buscarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mutex.RLock()
	cliente, existe := clientes[id]
	mutex.RUnlock()

	if !existe {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

func atualizarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var cliente Cliente

	err = json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if cliente.Nome == "" || cliente.Telefone == "" || cliente.Email == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, existe := clientes[id]

	if !existe {
		mutex.Unlock()
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	cliente.ID = id
	clientes[id] = cliente
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

func excluirCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, existe := clientes[id]

	if !existe {
		mutex.Unlock()
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	delete(clientes, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clientes", listarClientes).Methods("GET")
	router.HandleFunc("/clientes", criarCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", buscarCliente).Methods("GET")
	router.HandleFunc("/clientes/{id}", atualizarCliente).Methods("PUT")
	router.HandleFunc("/clientes/{id}", excluirCliente).Methods("DELETE")

	log.Println("Servidor iniciado em http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
