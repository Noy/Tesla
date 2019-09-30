package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	utils "github.com/TrafficLabel/Go-Utilities"
	"io/ioutil"
	"log"
	"net/http"
)

//https://akrion.docs.apiary.io/#reference/vehicles/vehicle-collection/list-all-vehicles?console=1
//https://owner-api.teslamotors.com/api/1/vehicles/:id/data_request/charge_state

func main() {
	fmt.Println(getVehicleState("CarVersion"))
}

type Config struct {
	AccessToken, ID string
}

func loadConfig() (c Config, err error) {
	_, err = toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Printf("Could not load config: %v", err.Error())
	}
	return
}

func getConfig() Config {
	cfg, err := loadConfig()
	if err != nil {
		log.Println("Could not open config")
	}
	return cfg
}

func commandBaseRequest(cmd string) (string, error) {
	cfg := getConfig()
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://owner-api.teslamotors.com/api/1/vehicles/"+cfg.ID+"/command/" + cmd, nil)
	if err != nil {
		return "Error with URL Request", err
	}
	req.Header.Set("Authorization", "Bearer " + cfg.AccessToken)
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
	return utils.JsonPrettyPrint(string(respBody)), nil
}

func getVehicleState(info string) interface{} {
	cfg := getConfig()
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://owner-api.teslamotors.com/api/1/vehicles/43330298813695250/vehicle_data", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer " + cfg.AccessToken)
	req.Header.Set("user-agent", "007")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if err != nil {
		log.Println(err.Error())
	}
	var r TeslaState
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &r); err != nil {
		log.Println(err.Error())
	}
	t := r.TeslaStateResponse
	tsr := r.TeslaStateResponse.ChargeState
	tsrv := r.TeslaStateResponse.VehicleState
	tsrd := r.TeslaStateResponse.DriveState
	tsrg := r.TeslaStateResponse.GUISettings
	switch info {
	case "CarType":
		return tsrv.CarType
	case "ChargePortType":
		return tsrv.ChargePortType
	case "IsEUVehicle":
		return tsrv.IsEUVehicle
	case "ExteriorColor":
		return tsrv.ExteriorColor
	case "Odometer":
		return tsrv.Odometer
	case "CarVersion":
		return tsrv.CarVersion
	case "ChargingState":
		return tsr.ChargingState
	case "BatteryRange":
		return tsr.BatteryRange
	case "BatteryLevel":
		return tsr.BatteryLevel
	case "ChargePortLatch":
		return tsr.ChargePortLatch
	case "BatteryHeaterOn":
		return tsr.BatteryHeaterOn
	case "ChargeCurrentRequest":
		return tsr.ChargeCurrentRequest
	case "ChargeCurrentRequestMax":
		return tsr.ChargeCurrentRequestMax
	case "ChargeEnableRequest":
		return tsr.ChargeEnableRequest
	case "ChargeEnergyAdded":
		return tsr.ChargeEnergyAdded
	case "ChargeLimitSoc":
		return tsr.ChargeLimitSoc
	case "ChargeLimitSocMax":
		return tsr.ChargeLimitSocMax
	case "ChargeLimitSocMin":
		return tsr.ChargeLimitSocMin
	case "ChargeLimitSocStd":
		return tsr.ChargeLimitSocStd
	case "ChargeMilesAddedIdeal":
		return tsr.ChargeMilesAddedIdeal
	case "ChargeMilesAddedRated":
		return tsr.ChargeMilesAddedRated
	case "ChargePointDoorOpen":
		return tsr.ChargePointDoorOpen
	case "ChargeRate":
		return tsr.ChargeRate
	case "ChargeToMaxRange":
		return tsr.ChargeToMaxRange
	case "ChargerActualCurrent":
		return tsr.ChargerActualCurrent
	case "ChargerPhases":
		return tsr.ChargerPhases
	case "ChargerPilotCurrent":
		return tsr.ChargerPilotCurrent
	case "ChargerPower":
		return tsr.ChargerPower
	case "EstimatedBatteryRange":
		return tsr.EstimatedBatteryRange
	case "ChargerVoltage":
		return tsr.ChargerVoltage
	case "TimeTilFullCharge":
		return tsr.TimeTilFullCharge
	case "IdealBatteryRange":
		return tsr.IdealBatteryRange
	case "ShiftState":
		return tsrd.ShiftState
	case "Speed":
		return tsrd.Speed
	case "Power":
		return tsrd.Power
	case "Latitude":
		return tsrd.Latitude
	case "Longitude":
		return tsrd.Longitude
	case "Heading":
		return tsrd.Heading
	case "GpsAsOf":
		return tsrg.GpsAsOf
	case "GUIDistanceUnits":
		return tsrg.GUIDistanceUnits
	case "GUITemperatureUnits":
		return tsrg.GUITemperatureUnits
	case "GUIChargeRateUnits":
		return tsrg.GUIChargeRateUnits
	case "GUI24HourTime":
		return tsrg.GUI24HourTime
	case "GUIRangeDisplay":
		return tsrg.GUIRangeDisplay
	case "ID":
		return t.ID
	case "Name":
		return t.Name
	case "State":
		return t.State
	case "VIN":
		return t.VIN
	case "UserID":
		return t.UserID
	case "VehicleID":
		return t.VehicleID
	case "InService":
		return t.InService
	}
	return "Setting not found"
}