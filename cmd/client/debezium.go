package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ConnectorHealthCheck() {
	response, err := http.Get("http://localhost:8083/connectors/cdc_main_postgres")
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		RegisterConnector()
	}
}

func RegisterConnector() *http.Response {
	plan, _ := ioutil.ReadFile("./connector/debezium-connector.json")
	response, err := http.Post("http://localhost:8083/connectors/", "application/json", bytes.NewBuffer(plan))
	fmt.Println(bytes.NewBuffer(plan))

	if err != nil {
		panic(err)
	}

	return response
}
