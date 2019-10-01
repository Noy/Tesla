package tesla

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (a *AuthTesla) listVehiclesData() []TeslaStateResponse {
	resp, err := a.getVehicleState("GET", "https://owner-api.teslamotors.com/api/1/vehicles")
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

func (a *AuthTesla) stateRequest() *TeslaState {
	resp, err := a.getVehicleState("GET", "https://owner-api.teslamotors.com/api/1/vehicles/"+a.ID+"/vehicle_data")
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

func (a *AuthTesla) getVehicleState(method, url string) (*http.Response, error) {
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
