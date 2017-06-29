package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server running")

	http.HandleFunc("/", index)
	http.HandleFunc("/groups/add", addNewGroup)
	http.HandleFunc("/groups/list", listAllGroups)
	http.HandleFunc("/groups/delete/{id}", deleteGroup)
	http.HandleFunc("/groups/edit/{id}", editGroup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey du")
}

func addNewGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "this action will add a new group")
	} else {
		fmt.Fprintf(w, "bad request %s", r.Method)
	}
}

func listAllGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "this action will lsit all groups")
	} else {
		fmt.Fprintf(w, "bad request %s", r.Method)
	}
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "this action will delete the group with id %s", r.PostFormValue(groupId))
	}
}

func editGroup(w http.ResponseWriter, r *http.Request) {

}
