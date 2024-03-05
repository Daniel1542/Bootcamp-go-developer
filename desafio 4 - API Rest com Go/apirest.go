package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "Person type" (tipo um objeto)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func loadPeopleData() error {
	fileContent, err := ioutil.ReadFile("pessoas.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &people)
	if err != nil {
		return err
	}

	return nil
}

func savePeopleData() error {
	fileContent, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("pessoas.json", fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}


func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = generateID()
	people = append(people, person)
	if err := savePeopleData(); err != nil {
		http.Error(w, "Erro ao salvar dados", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			var updatedPerson Person
			_ = json.NewDecoder(r.Body).Decode(&updatedPerson)
			updatedPerson.ID = item.ID
			people[index] = updatedPerson
			if err := savePeopleData(); err != nil {
				http.Error(w, "Erro ao salvar dados", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			if err := savePeopleData(); err != nil {
				http.Error(w, "Erro ao salvar dados", http.StatusInternalServerError)
				return
			}
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func generateID() string {
	return "IDGeradoDinamicamente"
}

// função principal para executar a api
func main() {
	if err := loadPeopleData(); err != nil {
		log.Fatal("Erro ao carregar dados do arquivo JSON:", err)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
