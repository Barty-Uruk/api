package main

import (
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "io/ioutil"
    "reflect"
)

// The person Type (more like an object)
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
type Hello struct {
    name  string `json:"name"`
}


var people []Person


// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
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

func HelloRaf (w http.ResponseWriter, r *http.Request) {
    var dat map[string]interface{}
    // var hello Hello
    // byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    //data, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    body, _ := ioutil.ReadAll(r.Body)
    // decoder := json.NewDecoder(r.Body)
    // decoder.Decode(&hello)
    // _ = json.NewDecoder(r.Body)

  //  _ = dec.Decode(&hello)
    json.Unmarshal(body, &dat)
    // r.ParseForm()
    // x := r.Form.Get("name")
    // helloName2 := json.RawMessage(helloName)
    // json.Unmarshal(body, &hello)
    // fmt.Println(string(hello.Name))
    // decoder := json.NewDecoder(r.Body)
      fmt.Println(string(body))

      i := dat["name"]
      fmt.Println(reflect.TypeOf(body))
      // hello.name = string(dat["name"])

      json.NewEncoder(w).Encode(i)
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    router.HandleFunc("/hello", HelloRaf).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
