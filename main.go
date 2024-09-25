package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users  = []User{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	http.HandleFunc("/users", handleUsers)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(users)
	case "POST":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		mu.Lock()
		user.ID = nextID
		nextID++
		users = append(users, user)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	case "PUT":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		mu.Lock()
		for i, u := range users {
			if u.ID == user.ID {
				users[i] = user
				break
			}
		}
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	case "DELETE":
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		mu.Lock()
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
