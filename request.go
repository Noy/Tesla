// Copyright 2019 Noy H. All rights reserved.
// @author Noy Hillel
package tesla

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (a *AuthTesla) ListVehicles() []StateResponse {
	resp, err := a.baseVehicleRequest("GET", "https://owner-api.teslamotors.com/api/1/vehicles")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	var r Vehicle
	respBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &r); err != nil {
		log.Println(err.Error())
	}
	return r.StateResponse
}

func (a *AuthTesla) GetVehicleData() StateResponse {
	resp, err := a.baseVehicleRequest("GET", "https://owner-api.teslamotors.com/api/1/vehicles/"+a.ID+"/vehicle_data")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	var r State
	respBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &r); err != nil {
		log.Println(err.Error())
	}
	return r.StateResponse
}

func (a *AuthTesla) baseVehicleRequest(method, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
	req.Header.Set("user-agent", "007")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
