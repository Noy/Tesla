package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func commandBaseRequest(cmd string) (string, error) {
	cfg := getConfig()
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://owner-api.teslamotors.com/api/1/vehicles/"+cfg.ID+"/command/"+cmd, nil)
	if err != nil {
		return "Error with URL Request", err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.AccessToken)
	req.Header.Set("user-agent", "007")
	resp, err := client.Do(req)
	if err != nil {
		return "error with doing request", err
	}
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Could not do it", err
	}
	return jsonPrettyPrint(string(respBody)), nil
}

func wakeUp() { doRequest("wake_up") }

func lock(setting bool) {
	lock := "lock"
	if !setting {
		lock = "unlock"
	}
	doRequest("door_" + lock)
}

func honkHorn() { doRequest("honk_horn") }

func flashLights() { doRequest("flash_lights") }

func autoConditioningStart() { doRequest("auto_conditioning_start") }

func autoConditioningStop() { doRequest("auto_conditioning_stop") }

func setTemps(driverTemp, passengerTemp string) {
	doRequest("set_temps?driver_temp=" + driverTemp + "passanger_temp=" + passengerTemp)
}

func setChargeLimit(percent string) { doRequest("set_charge_limit?percent=" + percent) }

func chargeMaxRange() { doRequest("charge_max_range") }

func chargeStandard() { doRequest("charge_standard") }

func acuteTrunk() { doRequest("actuate_trunk") }

func chargePort(openOrClose string) { doRequest("charge_port_door_" + openOrClose) }

func charge(startOrStop string) { doRequest("charge_" + startOrStop) }

func setValetMode(on string, pin string) { doRequest("set_valet_mode?on=" + on + "&password=" + pin) }

func doRequest(request string) {
	do, err := commandBaseRequest(request)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Print(do)
}

//func openTrunk(which string) {
//	id := getVehicleID()
//	idAsString := strconv.Itoa(int(id))
//	resp, err := baseRequest("POST", "vehicles/"+idAsString+"/command/trunk_open")
//	if err != nil {
//		log.Println(err.Error())
//	}
//	defer resp.Body.Close()
//	respBody, _ := ioutil.ReadAll(resp.Body)
//	log.Println(string(respBody))
//}
