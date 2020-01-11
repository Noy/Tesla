// Copyright 2019 Noy H. All rights reserved.
// @author Noy Hillel
package tesla

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (a *AuthTesla) commandBaseRequest(cmd string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://owner-api.teslamotors.com/api/1/vehicles/"+a.ID+"/command/"+cmd, nil)
	if err != nil {
		return "Error with URL Request", err
	}
	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
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

func (a *AuthTesla) WakeUp() { a.doRequest("wake_up") }

func (a *AuthTesla) Lock(setting bool) {
	lock := "lock"
	if !setting {
		lock = "unlock"
	}
	a.doRequest("door_" + lock)
}

func (a *AuthTesla) HonkHorn() {
	a.doRequest("honk_horn")
}

func (a *AuthTesla) FlashLights() { a.doRequest("flash_lights") }

func (a *AuthTesla) AutoConditioning(startOrStop string) {
	a.doRequest("auto_conditioning_" + startOrStop)
}

func (a *AuthTesla) SetTemps(driverTemp, passengerTemp string) {
	a.doRequest("set_temps?driver_temp=" + driverTemp + "passanger_temp=" + passengerTemp)
}

func (a *AuthTesla) SetChargeLimit(percent string) {
	a.doRequest("set_charge_limit?percent=" + percent)
}

func (a *AuthTesla) ChargeMaxRange() {
	a.doRequest("charge_max_range")
}

func (a *AuthTesla) ChargeStandard() {
	a.doRequest("charge_standard")
}

func (a *AuthTesla) AcuteTrunk() {
	a.doRequest("actuate_trunk")
}

func (a *AuthTesla) ChargePort(openOrClose string) {
	a.doRequest("charge_port_door_" + openOrClose)
}

func (a *AuthTesla) Charge(startOrStop string) {
	a.doRequest("charge_" + startOrStop)
}

func (a *AuthTesla) SetValetMode(on string, pin string) {
	a.doRequest("set_valet_mode?on=" + on + "&password=" + pin)
}

func (a *AuthTesla) SetHeatedSeatTo(heater, level string) {
	a.doRequest("remote_seat_heater_request?heater=" + heater + "&level=" + level)
}

func (a *AuthTesla) OpenTrunk(which string) {
	a.doRequest("actuate_trunk?which_trunk=" + which)
}

func (a *AuthTesla) SetSentryMode(val string) {
	a.doRequest("set_sentry_mode?on=" + val)
}

func (a *AuthTesla) doRequest(request string) {
	do, err := a.commandBaseRequest(request)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Print(do)
}
