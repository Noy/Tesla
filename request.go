package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func listVehiclesData() []TeslaStateResponse {
	resp, err := getVehicleState("GET", "https://owner-api.teslamotors.com/api/1/vehicles")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	var r TeslaVehicle
	respBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &r); err != nil {
		log.Println(err.Error())
	}
	return r.TeslaStateResponse
}

func stateRequest() *TeslaState {
	cfg := getConfig()
	resp, err := getVehicleState("GET", "https://owner-api.teslamotors.com/api/1/vehicles/"+cfg.ID+"/vehicle_data")
	if err != nil {
		log.Println(err.Error())
	}
	var r TeslaState
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &r); err != nil {
		log.Println(err.Error())
	}
	return &r
}

func getVehicleState(method, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+getConfig().AccessToken)
	req.Header.Set("user-agent", "007")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
