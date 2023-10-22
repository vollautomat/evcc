package keys

const (
	// loadpoint settings
	Title = "title" // loadpoint title

	PhasesConfigured = "phasesConfigured" // configured phases (1/3, 0 for auto on 1p3p chargers, nil for plain chargers)
	PhasesEnabled    = "phasesEnabled"    // enabled phases (1/3)
	PhasesActive     = "phasesActive"     // active phases as used by vehicle (1/2/3)

	ChargerIcon    = "chargerIcon"    // charger icon for ui
	ChargerFeature = "chargerFeature" // charger feature

	// vehicle
	ClimaterActive         = "climaterActive"         // climater active
	VehicleIdentity        = "vehicleIdentity"        // vehicle identity
	VehicleCapacity        = "vehicleCapacity"        // vehicle battery capacity
	VehicleDetectionActive = "vehicleDetectionActive" // vehicle detection active
	VehicleIcon            = "vehicleIcon"            // vehicle icon for ui
	VehicleOdometer        = "vehicleOdometer"        // vehicle odometer
	VehiclePresent         = "vehiclePresent"         // vehicle detected
	VehicleRange           = "vehicleRange"           // vehicle range
	VehicleSoc             = "vehicleSoc"             // vehicle soc
	VehicleTargetSoc       = "vehicleTargetSoc"       // vehicle soc limit
	VehicleTitle           = "vehicleTitle"           // vehicle title

	// TODO rename value
	VehicleClimaterActive = "vehicleClimaterActive" // vehicle climater active

	// loadpoint status
	Mode        = "mode"        // charge mode
	Priority    = "priority"    // priority
	Enabled     = "enabled"     // loadpoint enabled
	Connected   = "connected"   // connected
	Charging    = "charging"    // charging
	MinCurrent  = "minCurrent"  // min current
	MaxCurrent  = "maxCurrent"  // max current
	MinSoc      = "minSoc"      // min soc
	LimitSoc    = "limitSoc"    // limit soc
	LimitEnergy = "limitEnergy" // limit energy

	// measurements
	ChargeCurrent     = "chargeCurrent"     // charge current
	ChargePower       = "chargePower"       // charge power
	ChargeCurrents    = "chargeCurrents"    // charge currents
	ChargeVoltages    = "chargeVoltages"    // charge voltages
	ChargedEnergy     = "chargedEnergy"     // charged energy
	ChargeDuration    = "chargeDuration"    // charge duration
	ChargeTotalImport = "chargeTotalImport" // charge meter total import

	// session
	ConnectedDuration       = "connectedDuration"       // connected duration
	ChargeRemainingDuration = "chargeRemainingDuration" // charge remaining duration
	ChargeRemainingEnergy   = "chargeRemainingEnergy"   // charge remaining energy

	// plan
	PlanTime           = "planTime"           // charge plan finish time goal
	PlanEnergy         = "planEnergy"         // charge plan energy goal
	PlanSoc            = "planSoc"            // charge plan soc goal
	PlanActive         = "planActive"         // charge plan has determined current slot to be an active slot
	PlanProjectedStart = "planProjectedStart" // charge plan start time (earliest slot)

	// remote control
	RemoteDisabled       = "remoteDisabled"       // remote disabled
	RemoteDisabledSource = "remoteDisabledSource" // remote disabled source
)
