package main

type TeslaState struct {
	TeslaStateResponse TeslaStateResponse `json:"response"`
}

type TeslaVehicle struct {
	TeslaStateResponse []TeslaStateResponse `json:"response"`
}

type ChargeState struct {
	ChargingState               string  `json:"charging_state"`
	FastChargerType             string  `json:"fast_charger_type"`
	FastChargerBrand            string  `json:"fast_charger_brand"`
	ChargeLimitSoc              int     `json:"charge_limit_soc"`
	ChargeLimitSocStd           int     `json:"charge_limit_soc_std"`
	ChargeLimitSocMin           int     `json:"charge_limit_soc_min"`
	ChargeLimitSocMax           int     `json:"charge_limit_soc_max"`
	ChargeToMaxRange            bool    `json:"charge_to_max_range"`
	FastChargerPresent          bool    `json:"fast_charger_present"`
	BatteryRange                float64 `json:"battery_range"`
	EstimatedBatteryRange       float64 `json:"est_battery_range"`
	IdealBatteryRange           float64 `json:"ideal_battery_range"`
	BatteryLevel                int     `json:"battery_level"`
	UsableBatteryLevel          int     `json:"usable_battery_level"`
	ChargeEnergyAdded           float64 `json:"charge_energy_added"`
	ChargeMilesAddedRated       float64 `json:"charge_miles_added_rated"`
	ChargeMilesAddedIdeal       float64 `json:"charge_miles_added_ideal"`
	ChargerVoltage              int     `json:"charger_voltage"`
	ChargerPilotCurrent         int     `json:"charger_pilot_current"`
	ChargerActualCurrent        int     `json:"charger_actual_current"`
	ChargerPower                int     `json:"charger_power"`
	TimeTilFullCharge           float64 `json:"minutes_to_full_charge"`
	TripCharging                bool    `json:"trip_charging"`
	ChargeRate                  float64 `json:"charge_rate"`
	ChargePointDoorOpen         bool    `json:"charge_port_door_open"`
	ConnChargeCable             string  `json:"conn_charge_cable"`
	ScheduledChargingStartTime  string  `json:"scheduled_charging_start_time"`
	ScheduledChargingPending    bool    `json:"scheduled_charging_pending"`
	UserChargeEnableRequest     string  `json:"user_charge_enable_request"`
	ChargeEnableRequest         bool    `json:"charge_enable_request"`
	ChargerPhases               string  `json:"charger_phases"`
	ChargePortLatch             string  `json:"charge_port_latch"`
	ChargeCurrentRequest        int     `json:"charge_current_request"`
	ChargeCurrentRequestMax     int     `json:"charge_current_request_max"`
	ManagedChargingActive       bool    `json:"managed_charging_active"`
	ManagedChargingUserCanceled bool    `json:"managed_charging_user_canceled"`
	ManagedChargingStartTime    string  `json:"managed_charging_start_time"`
	BatteryHeaterOn             bool    `json:"battery_heater_on"`
	NotEnoughPowerToHeat        bool    `json:"not_enough_power_to_heat"`
}

type VehicleState struct {
	CarType        string  `json:"car_type"`
	ChargePortType string  `json:"charge_port_type"`
	IsEUVehicle    bool    `json:"eu_vehicle"`
	ExteriorColor  string  `json:"exterior_color"`
	CarVersion     string  `json:"car_version"`
	SentryMode     bool    `json:"sentry_mode"`
	Odometer       float64 `json:"Odometer"`
}

type DriveState struct {
	ShiftState string  `json:"shift_state"`
	Speed      string  `json:"speed"`
	Power      int     `json:"power"`
	Heading    int     `json:"heading"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type GUISettings struct {
	GpsAsOf             int64  `json:"gps_as_of"`
	GUIDistanceUnits    string `json:"gui_distance_units"`
	GUITemperatureUnits string `json:"gui_temperature_units"`
	GUIChargeRateUnits  string `json:"gui_charge_rate_units"`
	GUI24HourTime       bool   `json:"gui_24_hour_time"`
	GUIRangeDisplay     string `json:"gui_range_display"`
}

type TeslaStateResponse struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	VehicleID   int64  `json:"vehicle_id"`
	VIN         string `json:"vin"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	State       string `json:"state"`
	InService   bool   `json:"in_service"`
	//ClimateState ClimateState `json:"climate_state"`
	DriveState   DriveState   `json:"drive_state"`
	ChargeState  ChargeState  `json:"charge_state"`
	GUISettings  GUISettings  `json:"gui_settings"`
	VehicleState VehicleState `json:"vehicle_state"`
}
