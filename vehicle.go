package tesla

func (a *AuthTesla) GetCarType() string {
	return a.stateRequest().StateResponse.VehicleState.CarType
}

func (a *AuthTesla) GetChargePortType() string {
	return a.stateRequest().StateResponse.VehicleState.ChargePortType
}

func (a *AuthTesla) IsEUVehicle() bool {
	return a.stateRequest().StateResponse.VehicleState.IsEUVehicle
}

func (a *AuthTesla) GetExteriorColor() string {
	return a.stateRequest().StateResponse.VehicleState.ExteriorColor
}

func (a *AuthTesla) GetOdometerReading() float64 {
	return a.stateRequest().StateResponse.VehicleState.Odometer
}

func (a *AuthTesla) GetCarVersion() string {
	return a.stateRequest().StateResponse.VehicleState.CarVersion
}

func (a *AuthTesla) GetChargingState() string {
	return a.stateRequest().StateResponse.ChargeState.ChargingState
}

func (a *AuthTesla) GetBatteryLevel() int {
	return a.stateRequest().StateResponse.ChargeState.BatteryLevel
}

func (a *AuthTesla) GetChargePortLatch() string {
	return a.stateRequest().StateResponse.ChargeState.ChargePortLatch
}

func (a *AuthTesla) BatteryHeaterOn() bool {
	return a.stateRequest().StateResponse.ChargeState.BatteryHeaterOn
}

func (a *AuthTesla) GetChargeCurrentRequest() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeCurrentRequest
}

func (a *AuthTesla) GetChargeCurrentRequestMax() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeCurrentRequestMax
}

func (a *AuthTesla) ChargeEnableRequest() bool {
	return a.stateRequest().StateResponse.ChargeState.ChargeEnableRequest
}

func (a *AuthTesla) GetChargeEnergyAdded() float64 {
	return a.stateRequest().StateResponse.ChargeState.ChargeEnergyAdded
}

func (a *AuthTesla) GetChargeLimitSoc() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeLimitSoc
}

func (a *AuthTesla) GetChargeLimitSocMax() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeLimitSocMax
}

func (a *AuthTesla) GetChargeLimitSocMin() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeLimitSocMin
}

func (a *AuthTesla) GetChargeLimitSocStd() int {
	return a.stateRequest().StateResponse.ChargeState.ChargeLimitSocStd
}

func (a *AuthTesla) GetChargeMilesAddedIdeal() float64 {
	return a.stateRequest().StateResponse.ChargeState.ChargeMilesAddedIdeal
}

func (a *AuthTesla) GetChargeMilesAddedRated() float64 {
	return a.stateRequest().StateResponse.ChargeState.ChargeMilesAddedRated
}

func (a *AuthTesla) GetChargePointDoorOpen() bool {
	return a.stateRequest().StateResponse.ChargeState.ChargePointDoorOpen
}

func (a *AuthTesla) GetChargeRate() float64 {
	return a.stateRequest().StateResponse.ChargeState.ChargeRate
}

func (a *AuthTesla) chargeToMaxRange() bool {
	return a.stateRequest().StateResponse.ChargeState.ChargeToMaxRange
}

func (a *AuthTesla) GetChargerActualCurrent() int {
	return a.stateRequest().StateResponse.ChargeState.ChargerActualCurrent
}

func (a *AuthTesla) GetChargerPhases() string {
	return a.stateRequest().StateResponse.ChargeState.ChargerPhases
}

func (a *AuthTesla) GetChargerPilotCurrent() int {
	return a.stateRequest().StateResponse.ChargeState.ChargerPilotCurrent
}

func (a *AuthTesla) GetChargerPower() int {
	return a.stateRequest().StateResponse.ChargeState.ChargerPower
}

func (a *AuthTesla) GetEstimatedBatteryRange() float64 {
	return a.stateRequest().StateResponse.ChargeState.EstimatedBatteryRange
}

func (a *AuthTesla) GetChargerVoltage() int {
	return a.stateRequest().StateResponse.ChargeState.ChargerVoltage
}

func (a *AuthTesla) GetTimeTilFullCharge() float64 {
	return a.stateRequest().StateResponse.ChargeState.TimeTilFullCharge
}

func (a *AuthTesla) GetIdealBatteryRange() float64 {
	return a.stateRequest().StateResponse.ChargeState.IdealBatteryRange
}

func (a *AuthTesla) GetShiftState() string {
	return a.stateRequest().StateResponse.DriveState.ShiftState
}

func (a *AuthTesla) GetSpeed() string {
	return a.stateRequest().StateResponse.DriveState.Speed
}

func (a *AuthTesla) GetPower() int {
	return a.stateRequest().StateResponse.DriveState.Power
}

func (a *AuthTesla) GetLatitude() float64 {
	return a.stateRequest().StateResponse.DriveState.Latitude
}

func (a *AuthTesla) GetLongitude() float64 {
	return a.stateRequest().StateResponse.DriveState.Longitude
}

func (a *AuthTesla) GetHeading() int {
	return a.stateRequest().StateResponse.DriveState.Heading
}

func (a *AuthTesla) GetGpsAsOf() int64 {
	return a.stateRequest().StateResponse.GUISettings.GpsAsOf
}

func (a *AuthTesla) GetDistanceUnits() string {
	return a.stateRequest().StateResponse.GUISettings.GUIDistanceUnits
}

func (a *AuthTesla) GetTemperatureUnits() string {
	return a.stateRequest().StateResponse.GUISettings.GUITemperatureUnits
}

func (a *AuthTesla) GetChargeRateUnits() string {
	return a.stateRequest().StateResponse.GUISettings.GUIChargeRateUnits
}

func (a *AuthTesla) Is24HourTime() bool {
	return a.stateRequest().StateResponse.GUISettings.GUI24HourTime
}

func (a *AuthTesla) GetRangeDisplay() string {
	return a.stateRequest().StateResponse.GUISettings.GUIRangeDisplay
}

func (a *AuthTesla) GetCarID() int64 {
	return a.stateRequest().StateResponse.ID
}

func (a *AuthTesla) GetCarName() string {
	return a.stateRequest().StateResponse.Name
}

func (a *AuthTesla) GetCarState() string {
	return a.stateRequest().StateResponse.State
}

func (a *AuthTesla) GetCarVIN() string {
	return a.stateRequest().StateResponse.VIN
}

func (a *AuthTesla) GetCarUserID() int64 {
	return a.stateRequest().StateResponse.UserID
}

func (a *AuthTesla) GetCarVehicleID() int64 {
	return a.stateRequest().StateResponse.VehicleID
}

func (a *AuthTesla) IsInService() bool {
	return a.stateRequest().StateResponse.InService
}
