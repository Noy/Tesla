// Copyright 2019 Noy H. All rights reserved.
// @author Noy Hillel
package tesla

func (a *AuthTesla) GetCarType() string {
	return a.GetVehicleData().VehicleConfig.CarType
}

func (a *AuthTesla) GetChargePortType() string {
	return a.GetVehicleData().VehicleConfig.ChargePortType
}

func (a *AuthTesla) IsEUVehicle() bool {
	return a.GetVehicleData().VehicleState.IsEUVehicle
}

func (a *AuthTesla) GetExteriorColor() string {
	return a.GetVehicleData().VehicleConfig.ExteriorColor
}

func (a *AuthTesla) GetWheelType() string {
	return a.GetVehicleData().VehicleConfig.WheelType
}

func (a *AuthTesla) GetOdometerReading() float64 {
	return a.GetVehicleData().VehicleState.Odometer
}

func (a *AuthTesla) GetCarVersion() string {
	return a.GetVehicleData().VehicleState.CarVersion
}

func (a *AuthTesla) GetChargingState() string {
	return a.GetVehicleData().ChargeState.ChargingState
}

func (a *AuthTesla) GetBatteryLevel() int {
	return a.GetVehicleData().ChargeState.BatteryLevel
}

func (a *AuthTesla) GetChargePortLatch() string {
	return a.GetVehicleData().ChargeState.ChargePortLatch
}

func (a *AuthTesla) BatteryHeaterOn() bool {
	return a.GetVehicleData().ChargeState.BatteryHeaterOn
}

func (a *AuthTesla) GetChargeCurrentRequest() int {
	return a.GetVehicleData().ChargeState.ChargeCurrentRequest
}

func (a *AuthTesla) GetChargeCurrentRequestMax() int {
	return a.GetVehicleData().ChargeState.ChargeCurrentRequestMax
}

func (a *AuthTesla) ChargeEnableRequest() bool {
	return a.GetVehicleData().ChargeState.ChargeEnableRequest
}

func (a *AuthTesla) GetChargeEnergyAdded() float64 {
	return a.GetVehicleData().ChargeState.ChargeEnergyAdded
}

func (a *AuthTesla) GetChargeLimitSoc() int {
	return a.GetVehicleData().ChargeState.ChargeLimitSoc
}

func (a *AuthTesla) GetChargeLimitSocMax() int {
	return a.GetVehicleData().ChargeState.ChargeLimitSocMax
}

func (a *AuthTesla) GetChargeLimitSocMin() int {
	return a.GetVehicleData().ChargeState.ChargeLimitSocMin
}

func (a *AuthTesla) GetChargeLimitSocStd() int {
	return a.GetVehicleData().ChargeState.ChargeLimitSocStd
}

func (a *AuthTesla) GetChargeMilesAddedIdeal() float64 {
	return a.GetVehicleData().ChargeState.ChargeMilesAddedIdeal
}

func (a *AuthTesla) GetChargeMilesAddedRated() float64 {
	return a.GetVehicleData().ChargeState.ChargeMilesAddedRated
}

func (a *AuthTesla) GetChargePointDoorOpen() bool {
	return a.GetVehicleData().ChargeState.ChargePointDoorOpen
}

func (a *AuthTesla) GetChargeRate() float64 {
	return a.GetVehicleData().ChargeState.ChargeRate
}

func (a *AuthTesla) chargeToMaxRange() bool {
	return a.GetVehicleData().ChargeState.ChargeToMaxRange
}

func (a *AuthTesla) GetChargerActualCurrent() int {
	return a.GetVehicleData().ChargeState.ChargerActualCurrent
}

func (a *AuthTesla) GetChargerPhases() string {
	return a.GetVehicleData().ChargeState.ChargerPhases
}

func (a *AuthTesla) GetChargerPilotCurrent() int {
	return a.GetVehicleData().ChargeState.ChargerPilotCurrent
}

func (a *AuthTesla) GetChargerPower() int {
	return a.GetVehicleData().ChargeState.ChargerPower
}

func (a *AuthTesla) GetEstimatedBatteryRange() float64 {
	return a.GetVehicleData().ChargeState.EstimatedBatteryRange
}

func (a *AuthTesla) GetChargerVoltage() int {
	return a.GetVehicleData().ChargeState.ChargerVoltage
}

func (a *AuthTesla) GetTimeTilFullCharge() float64 {
	return a.GetVehicleData().ChargeState.TimeTilFullCharge
}

func (a *AuthTesla) GetIdealBatteryRange() float64 {
	return a.GetVehicleData().ChargeState.IdealBatteryRange
}

func (a *AuthTesla) GetShiftState() string {
	return a.GetVehicleData().DriveState.ShiftState
}

func (a *AuthTesla) GetSpeed() string {
	return a.GetVehicleData().DriveState.Speed
}

func (a *AuthTesla) GetPower() int {
	return a.GetVehicleData().DriveState.Power
}

func (a *AuthTesla) GetLatitude() float64 {
	return a.GetVehicleData().DriveState.Latitude
}

func (a *AuthTesla) GetLongitude() float64 {
	return a.GetVehicleData().DriveState.Longitude
}

func (a *AuthTesla) GetHeading() int {
	return a.GetVehicleData().DriveState.Heading
}

func (a *AuthTesla) GetGpsAsOf() int64 {
	return a.GetVehicleData().GUISettings.GpsAsOf
}

func (a *AuthTesla) GetDistanceUnits() string {
	return a.GetVehicleData().GUISettings.GUIDistanceUnits
}

func (a *AuthTesla) GetTemperatureUnits() string {
	return a.GetVehicleData().GUISettings.GUITemperatureUnits
}

func (a *AuthTesla) GetChargeRateUnits() string {
	return a.GetVehicleData().GUISettings.GUIChargeRateUnits
}

func (a *AuthTesla) Is24HourTime() bool {
	return a.GetVehicleData().GUISettings.GUI24HourTime
}

func (a *AuthTesla) GetRangeDisplay() string {
	return a.GetVehicleData().GUISettings.GUIRangeDisplay
}

func (a *AuthTesla) GetCarID() int64 {
	return a.GetVehicleData().ID
}

func (a *AuthTesla) GetCarName() string {
	return a.GetVehicleData().Name
}

func (a *AuthTesla) GetCarState() string {
	return a.GetVehicleData().State
}

func (a *AuthTesla) GetCarVIN() string {
	return a.GetVehicleData().VIN
}

func (a *AuthTesla) GetCarUserID() int64 {
	return a.GetVehicleData().UserID
}

func (a *AuthTesla) GetID() int64 {
	return a.GetVehicleData().UserID
}

func (a *AuthTesla) GetCarVehicleID() int64 {
	return a.GetVehicleData().VehicleID
}

func (a *AuthTesla) IsInService() bool {
	return a.GetVehicleData().InService
}
