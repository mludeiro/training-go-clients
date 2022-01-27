package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pluralsight/webservice/pkg/models"
)

var NewClient models.Client

func CreateClient(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var client models.Client
	err = json.Unmarshal(bytes, &client)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b := client.CreateClient()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	clients := models.GetAllClients()
	res, _ := json.Marshal(clients)

	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["clientId"]
	ID, err := strconv.ParseInt(clientId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
		fmt.Println("Error:", err)
	}

	client, _ := models.GetBookById(ID)
	res, _ := json.Marshal(client)

	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var updateClient models.Client
	err = json.Unmarshal(bytes, &updateClient)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	vars := mux.Vars(r)
	clientId := vars["clientId"]
	ID, err := strconv.ParseInt(clientId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	clientDetails, db := models.GetBookById(ID)
	if updateClient.Name != "" {
		clientDetails.Name = updateClient.Name
	}

	db.Save(&clientDetails)
	res, _ := json.Marshal(clientDetails)

	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["clientId"]
	ID, err := strconv.ParseInt(clientId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	client := models.DeleteClient(ID)
	res, _ := json.Marshal(client)

	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
