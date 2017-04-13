package server

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"encoding/json"
	"os"
)

const CURRENT_SERVICE = "service-one"
const SERVICE_TO_CALL = "service-two"
var CONSUL = os.Getenv("CONSUL_HOST")

type ServiceResponse struct {
	ID      string        `json:"id"`
	Node    string        `json:"node"`
	Address string        `json:"address"`
	TaggedAddresses struct {
		lan string `json:"lan"`
		wan string `json:"lan"`
	}      `json:"TaggedAddresses"`
	NodeMeta struct {
	}      `json:"NodeMeta"`
	ServiceID                string        `json:"ServiceID"`
	ServiceName              string        `json:"ServiceName"`
	ServiceTags              []string      `json:"ServiceTags"`
	ServiceAddress           string        `json:"ServiceAddress"`
	ServicePort              int           `json:"ServicePort"`
	ServiceEnableTagOverride bool          `json:"ServiceEnableTagOverride"`
	CreateIndex              int           `json:"CreateIndex"`
	ModifyIndex              int           `json:"ModifyIndex"`
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	//err := registerService(port)

	router.HandleFunc("/"+CURRENT_SERVICE+"/other-service/", callService).Methods(http.MethodGet)
	router.HandleFunc("/"+CURRENT_SERVICE+"/", HandleRequests).Methods(http.MethodGet)

	httpServer := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), router)

	log.Fatal(httpServer)
}

func HandleRequests(writer http.ResponseWriter, reader *http.Request) {
	println("handleRequest...")
	io.WriteString(writer, "Hello World, I'm " + CURRENT_SERVICE)
}

func callService(w http.ResponseWriter, reader *http.Request) {
	println("handleRequest...")
	serviceOne := getOtherService(SERVICE_TO_CALL)
	result := callApi(serviceOne)
	io.WriteString(w, result)
}
func callApi(serviceUri string) string {
	client2 := &http.Client{}
	request2, err2 := http.NewRequest("GET", serviceUri, nil)
	if err2 != nil {
		panic(err2)
	}
	response2, err2 := client2.Do(request2)
	bodyBytes2, _ := ioutil.ReadAll(response2.Body)
	bodyString2 := string(bodyBytes2)
	defer response2.Body.Close()
	println(bodyString2)

	return "Called " + serviceUri + "\n=>" + bodyString2
}
func getOtherService(serviceName string) (string) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", CONSUL + "/" + serviceName, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	println(string(bodyBytes))
	var serviceresp []ServiceResponse
	if err := json.Unmarshal(bodyBytes, &serviceresp); err != nil {
		panic(err)
	}

	uri := fmt.Sprintf("http://%s:%d/%s", serviceresp[0].ServiceAddress, serviceresp[0].ServicePort, SERVICE_TO_CALL)

	return uri
}
