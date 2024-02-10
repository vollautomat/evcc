package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/evcc-io/evcc/charger"
	"github.com/evcc-io/evcc/meter"
	"github.com/evcc-io/evcc/util/config"
	"github.com/evcc-io/evcc/util/templates"
	"github.com/evcc-io/evcc/vehicle"
	"github.com/gorilla/mux"
)

// loadpointsHandler returns a device configurations by class
func loadpointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	class, err := templates.ClassString(vars["class"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	var res []map[string]any

	switch class {
	case templates.Meter:
		res, err = devicesConfig(class, config.Meters())

	case templates.Charger:
		res, err = devicesConfig(class, config.Chargers())

	case templates.Vehicle:
		res, err = devicesConfig(class, config.Vehicles())
	}

	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	jsonResult(w, res)
}

type loadpointConfig struct {
	ID         string
	Title      string
	ChargerRef string
	MeterRef   string
}

// newLoadpointHandler creates a new device by class
func newLoadpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var req loadpointConfig
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	delete(req, "type")

	conf, err := newDevice(class, req, loadpoint.NewFromConfig, config.Loadpoints())
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	setConfigDirty()

	res := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   conf.ID,
		Name: config.NameForID(conf.ID),
	}

	jsonResult(w, res)
}

// updateLoadpointHandler updates database device's configuration by class
func updateLoadpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	class, err := templates.ClassString(vars["class"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	var req map[string]any
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}
	delete(req, "type")

	switch class {
	case templates.Charger:
		err = updateDevice(id, class, req, charger.NewFromConfig, config.Chargers())

	case templates.Meter:
		err = updateDevice(id, class, req, meter.NewFromConfig, config.Meters())

	case templates.Vehicle:
		err = updateDevice(id, class, req, vehicle.NewFromConfig, config.Vehicles())
	}

	setConfigDirty()

	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	res := struct {
		ID int `json:"id"`
	}{
		ID: id,
	}

	jsonResult(w, res)
}

// deleteLoadpointHandler deletes a device from database by class
func deleteLoadpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	class, err := templates.ClassString(vars["class"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	switch class {
	case templates.Charger:
		err = deleteDevice(id, config.Chargers())

	case templates.Meter:
		err = deleteDevice(id, config.Meters())

	case templates.Vehicle:
		err = deleteDevice(id, config.Vehicles())
	}

	setConfigDirty()

	if err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	res := struct {
		ID int `json:"id"`
	}{
		ID: id,
	}

	jsonResult(w, res)
}
