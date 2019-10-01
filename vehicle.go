package main

func getCarType() string {
	return stateRequest().TeslaStateResponse.VehicleState.CarType
}

func getChargePortType() string {
	return stateRequest().TeslaStateResponse.VehicleState.ChargePortType
}
func isEUVehicle() bool {
	return stateRequest().TeslaStateResponse.VehicleState.IsEUVehicle
}
func getExteriorColor() string {
	return stateRequest().TeslaStateResponse.VehicleState.ExteriorColor
}

func getOdometerReading() float64 {
	return stateRequest().TeslaStateResponse.VehicleState.Odometer
}

func getCarVersion() string {
	return stateRequest().TeslaStateResponse.VehicleState.CarVersion
}

func getChargingState() string {
	return stateRequest().TeslaStateResponse.ChargeState.ChargingState
}

func getBatteryLevel() int {
	return stateRequest().TeslaStateResponse.ChargeState.BatteryLevel
}

func getChargePortLatch() string {
	return stateRequest().TeslaStateResponse.ChargeState.ChargePortLatch
}

func batteryHeaterOn() bool {
	return stateRequest().TeslaStateResponse.ChargeState.BatteryHeaterOn
}

func getChargeCurrentRequest() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeCurrentRequest
}

func getChargeCurrentRequestMax() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeCurrentRequestMax
}

func chargeEnableRequest() bool {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeEnableRequest
}

func getChargeEnergyAdded() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeEnergyAdded
}

func getChargeLimitSoc() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSoc
}

func getChargeLimitSocMax() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocMax
}

func getChargeLimitSocMin() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocMin
}

func getChargeLimitSocStd() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocStd
}

func getChargeMilesAddedIdeal() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeMilesAddedIdeal
}

func getChargeMilesAddedRated() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeMilesAddedRated
}

func getChargePointDoorOpen() bool {
	return stateRequest().TeslaStateResponse.ChargeState.ChargePointDoorOpen
}

func getChargeRate() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeRate
}

func chargeToMaxRange() bool {
	return stateRequest().TeslaStateResponse.ChargeState.ChargeToMaxRange
}

func getChargerActualCurrent() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargerActualCurrent
}

func getChargerPhases() string {
	return stateRequest().TeslaStateResponse.ChargeState.ChargerPhases
}

func getChargerPilotCurrent() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargerPilotCurrent
}

func getChargerPower() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargerPower
}

func getEstimatedBatteryRange() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.EstimatedBatteryRange
}

func getChargerVoltage() int {
	return stateRequest().TeslaStateResponse.ChargeState.ChargerVoltage
}

func getTimeTilFullCharge() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.TimeTilFullCharge
}

func getIdealBatteryRange() float64 {
	return stateRequest().TeslaStateResponse.ChargeState.IdealBatteryRange
}

func getShiftState() string {
	return stateRequest().TeslaStateResponse.DriveState.ShiftState
}

func getSpeed() string {
	return stateRequest().TeslaStateResponse.DriveState.Speed
}

func getPower() int {
	return stateRequest().TeslaStateResponse.DriveState.Power
}

func getLatitude() float64 {
	return stateRequest().TeslaStateResponse.DriveState.Latitude
}

func getLongitude() float64 {
	return stateRequest().TeslaStateResponse.DriveState.Longitude
}

func getHeading() int {
	return stateRequest().TeslaStateResponse.DriveState.Heading
}

func getGpsAsOf() int64 {
	return stateRequest().TeslaStateResponse.GUISettings.GpsAsOf
}

func getDistanceUnits() string {
	return stateRequest().TeslaStateResponse.GUISettings.GUIDistanceUnits
}

func getTemperatureUnits() string {
	return stateRequest().TeslaStateResponse.GUISettings.GUITemperatureUnits
}

func getChargeRateUnits() string {
	return stateRequest().TeslaStateResponse.GUISettings.GUIChargeRateUnits
}

func is24HourTime() bool {
	return stateRequest().TeslaStateResponse.GUISettings.GUI24HourTime
}

func getRangeDisplay() string {
	return stateRequest().TeslaStateResponse.GUISettings.GUIRangeDisplay
}

func getCarID() int64 {
	return stateRequest().TeslaStateResponse.ID
}

func getCarName() string {
	return stateRequest().TeslaStateResponse.Name
}

func getCarState() string {
	return stateRequest().TeslaStateResponse.State
}

func getCarVIN() string {
	return stateRequest().TeslaStateResponse.VIN
}

func getCarUserID() int64 {
	return stateRequest().TeslaStateResponse.UserID
}

func getCarVehicleID() int64 {
	return stateRequest().TeslaStateResponse.VehicleID
}

func isInService() bool {
	return stateRequest().TeslaStateResponse.InService
}
