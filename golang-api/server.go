package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Account struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	EmployeeCode int    `json:"employeeCode"`
}

func getFileData(filename string) ([]Account, error) {
	file, err := ioutil.ReadFile("base.json")
	if err != nil {
		return []Account{}, err
	}
	accounts := []Account{}
	err = json.Unmarshal(file, &accounts)
	if err != nil {

		return []Account{}, err
	}
	return accounts, nil
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	account := Account{}
	err := json.Unmarshal(reqBody, &account)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}

	accounts, err := getFileData("base.json")
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}
	accounts = append(accounts, account)

	file, _ := json.MarshalIndent(accounts, "", " ")
	_ = ioutil.WriteFile("base.json", file, 0644)
	fmt.Fprintf(w, "Created successful!")
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	account := Account{}
	err := json.Unmarshal(reqBody, &account)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}

	accounts, err := getFileData("base.json")
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}
	for i := 0; i < len(accounts); i++ {
		if accounts[i].EmployeeCode == account.EmployeeCode {
			accounts = append(accounts[:i], accounts[i+1:]...)
			file, _ := json.MarshalIndent(accounts, "", " ")
			_ = ioutil.WriteFile("base.json", file, 0644)
			fmt.Fprintf(w, "Deleted successful!")
			return
		}
	}
	fmt.Fprintf(w, "No Records found with this employee id to delete!")
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	account := Account{}
	err := json.Unmarshal(reqBody, &account)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}

	accounts, err := getFileData("base.json")
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}
	for i := 0; i < len(accounts); i++ {
		if accounts[i].EmployeeCode == account.EmployeeCode {
			accounts[i] = account
			file, _ := json.MarshalIndent(accounts, "", " ")
			_ = ioutil.WriteFile("base.json", file, 0644)
			fmt.Fprintf(w, "Updated successful!")
			return
		}
	}
	fmt.Fprintf(w, "No Records found with this employee id to update!")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	account := Account{}
	err := json.Unmarshal(reqBody, &account)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}

	file, err := ioutil.ReadFile("base.json")
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(file))
	return
}

func getByEmployeeCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	codeS := vars["code"]

	code, err := strconv.Atoi(codeS)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}

	accounts, err := getFileData("base.json")
	if err != nil {
		fmt.Fprintf(w, "Error: ", err.Error())
		return
	}
	for i := 0; i < len(accounts); i++ {
		if accounts[i].EmployeeCode == code {
			w.Header().Add("Content-Type", "application/json")
			response, err := json.Marshal(accounts[i])
			if err != nil {
				fmt.Fprintf(w, "Error: ", err.Error())
				return
			}
			fmt.Fprintf(w, string(response))
			return
		}
	}
	fmt.Fprintf(w, "No Records found with this employee id to return!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/account", createAccount).Methods("POST")
	myRouter.HandleFunc("/account", deleteAccount).Methods("DELETE")
	myRouter.HandleFunc("/account", updateAccount).Methods("PUT")
	myRouter.HandleFunc("/account", getAll).Methods("GET")
	myRouter.HandleFunc("/account/{code}", getByEmployeeCode).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter)) //":8081"
}

func main() {
	handleRequests()
}
