package tesla

func (a *AuthTesla) GetCarType() string {
	return a.stateRequest().TeslaStateResponse.VehicleState.CarType
}

func (a *AuthTesla) GetChargePortType() string {
	return a.stateRequest().TeslaStateResponse.VehicleState.ChargePortType
}

func (a *AuthTesla) IsEUVehicle() bool {
	return a.stateRequest().TeslaStateResponse.VehicleState.IsEUVehicle
}

func (a *AuthTesla) GetExteriorColor() string {
	return a.stateRequest().TeslaStateResponse.VehicleState.ExteriorColor
}

func (a *AuthTesla) GetOdometerReading() float64 {
	return a.stateRequest().TeslaStateResponse.VehicleState.Odometer
}

func (a *AuthTesla) GetCarVersion() string {
	return a.stateRequest().TeslaStateResponse.VehicleState.CarVersion
}

func (a *AuthTesla) GetChargingState() string {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargingState
}

func (a *AuthTesla) GetBatteryLevel() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.BatteryLevel
}

func (a *AuthTesla) GetChargePortLatch() string {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargePortLatch
}

func (a *AuthTesla) BatteryHeaterOn() bool {
	return a.stateRequest().TeslaStateResponse.ChargeState.BatteryHeaterOn
}

func (a *AuthTesla) GetChargeCurrentRequest() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeCurrentRequest
}

func (a *AuthTesla) GetChargeCurrentRequestMax() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeCurrentRequestMax
}

func (a *AuthTesla) ChargeEnableRequest() bool {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeEnableRequest
}

func (a *AuthTesla) GetChargeEnergyAdded() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeEnergyAdded
}

func (a *AuthTesla) GetChargeLimitSoc() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSoc
}

func (a *AuthTesla) GetChargeLimitSocMax() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocMax
}

func (a *AuthTesla) GetChargeLimitSocMin() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocMin
}

func (a *AuthTesla) GetChargeLimitSocStd() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeLimitSocStd
}

func (a *AuthTesla) GetChargeMilesAddedIdeal() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeMilesAddedIdeal
}

func (a *AuthTesla) GetChargeMilesAddedRated() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeMilesAddedRated
}

func (a *AuthTesla) GetChargePointDoorOpen() bool {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargePointDoorOpen
}

func (a *AuthTesla) GetChargeRate() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeRate
}

func (a *AuthTesla) chargeToMaxRange() bool {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargeToMaxRange
}

func (a *AuthTesla) GetChargerActualCurrent() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargerActualCurrent
}

func (a *AuthTesla) GetChargerPhases() string {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargerPhases
}

func (a *AuthTesla) GetChargerPilotCurrent() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargerPilotCurrent
}

func (a *AuthTesla) GetChargerPower() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargerPower
}

func (a *AuthTesla) GetEstimatedBatteryRange() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.EstimatedBatteryRange
}

func (a *AuthTesla) GetChargerVoltage() int {
	return a.stateRequest().TeslaStateResponse.ChargeState.ChargerVoltage
}

func (a *AuthTesla) GetTimeTilFullCharge() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.TimeTilFullCharge
}

func (a *AuthTesla) GetIdealBatteryRange() float64 {
	return a.stateRequest().TeslaStateResponse.ChargeState.IdealBatteryRange
}

func (a *AuthTesla) GetShiftState() string {
	return a.stateRequest().TeslaStateResponse.DriveState.ShiftState
}

func (a *AuthTesla) GetSpeed() string {
	return a.stateRequest().TeslaStateResponse.DriveState.Speed
}

func (a *AuthTesla) GetPower() int {
	return a.stateRequest().TeslaStateResponse.DriveState.Power
}

func (a *AuthTesla) GetLatitude() float64 {
	return a.stateRequest().TeslaStateResponse.DriveState.Latitude
}

func (a *AuthTesla) GetLongitude() float64 {
	return a.stateRequest().TeslaStateResponse.DriveState.Longitude
}

func (a *AuthTesla) GetHeading() int {
	return a.stateRequest().TeslaStateResponse.DriveState.Heading
}

func (a *AuthTesla) GetGpsAsOf() int64 {
	return a.stateRequest().TeslaStateResponse.GUISettings.GpsAsOf
}

func (a *AuthTesla) GetDistanceUnits() string {
	return a.stateRequest().TeslaStateResponse.GUISettings.GUIDistanceUnits
}

func (a *AuthTesla) GetTemperatureUnits() string {
	return a.stateRequest().TeslaStateResponse.GUISettings.GUITemperatureUnits
}

func (a *AuthTesla) GetChargeRateUnits() string {
	return a.stateRequest().TeslaStateResponse.GUISettings.GUIChargeRateUnits
}

func (a *AuthTesla) Is24HourTime() bool {
	return a.stateRequest().TeslaStateResponse.GUISettings.GUI24HourTime
}

func (a *AuthTesla) GetRangeDisplay() string {
	return a.stateRequest().TeslaStateResponse.GUISettings.GUIRangeDisplay
}

func (a *AuthTesla) GetCarID() int64 {
	return a.stateRequest().TeslaStateResponse.ID
}

func (a *AuthTesla) GetCarName() string {
	return a.stateRequest().TeslaStateResponse.Name
}

func (a *AuthTesla) GetCarState() string {
	return a.stateRequest().TeslaStateResponse.State
}

func (a *AuthTesla) GetCarVIN() string {
	return a.stateRequest().TeslaStateResponse.VIN
}

func (a *AuthTesla) GetCarUserID() int64 {
	return a.stateRequest().TeslaStateResponse.UserID
}

func (a *AuthTesla) GetCarVehicleID() int64 {
	return a.stateRequest().TeslaStateResponse.VehicleID
}

func (a *AuthTesla) IsInService() bool {
	return a.stateRequest().TeslaStateResponse.InService
}
