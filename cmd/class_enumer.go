// Code generated by "enumer -type Class -trimprefix Class -transform=lower -text"; DO NOT EDIT.

package cmd

import (
	"fmt"
	"strings"
)

const _ClassName = "configfilemeterchargervehicletariffcircuitsitemqttdatabasemodbusproxyeebusjavascriptgohemsinfluxmessengersponsorshiploadpoint"

var _ClassIndex = [...]uint8{0, 10, 15, 22, 29, 35, 42, 46, 50, 58, 69, 74, 84, 86, 90, 96, 105, 116, 125}

const _ClassLowerName = "configfilemeterchargervehicletariffcircuitsitemqttdatabasemodbusproxyeebusjavascriptgohemsinfluxmessengersponsorshiploadpoint"

func (i Class) String() string {
	i -= 1
	if i < 0 || i >= Class(len(_ClassIndex)-1) {
		return fmt.Sprintf("Class(%d)", i+1)
	}
	return _ClassName[_ClassIndex[i]:_ClassIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ClassNoOp() {
	var x [1]struct{}
	_ = x[ClassConfigFile-(1)]
	_ = x[ClassMeter-(2)]
	_ = x[ClassCharger-(3)]
	_ = x[ClassVehicle-(4)]
	_ = x[ClassTariff-(5)]
	_ = x[ClassCircuit-(6)]
	_ = x[ClassSite-(7)]
	_ = x[ClassMqtt-(8)]
	_ = x[ClassDatabase-(9)]
	_ = x[ClassModbusProxy-(10)]
	_ = x[ClassEEBus-(11)]
	_ = x[ClassJavascript-(12)]
	_ = x[ClassGo-(13)]
	_ = x[ClassHEMS-(14)]
	_ = x[ClassInflux-(15)]
	_ = x[ClassMessenger-(16)]
	_ = x[ClassSponsorship-(17)]
	_ = x[ClassLoadpoint-(18)]
}

var _ClassValues = []Class{ClassConfigFile, ClassMeter, ClassCharger, ClassVehicle, ClassTariff, ClassCircuit, ClassSite, ClassMqtt, ClassDatabase, ClassModbusProxy, ClassEEBus, ClassJavascript, ClassGo, ClassHEMS, ClassInflux, ClassMessenger, ClassSponsorship, ClassLoadpoint}

var _ClassNameToValueMap = map[string]Class{
	_ClassName[0:10]:         ClassConfigFile,
	_ClassLowerName[0:10]:    ClassConfigFile,
	_ClassName[10:15]:        ClassMeter,
	_ClassLowerName[10:15]:   ClassMeter,
	_ClassName[15:22]:        ClassCharger,
	_ClassLowerName[15:22]:   ClassCharger,
	_ClassName[22:29]:        ClassVehicle,
	_ClassLowerName[22:29]:   ClassVehicle,
	_ClassName[29:35]:        ClassTariff,
	_ClassLowerName[29:35]:   ClassTariff,
	_ClassName[35:42]:        ClassCircuit,
	_ClassLowerName[35:42]:   ClassCircuit,
	_ClassName[42:46]:        ClassSite,
	_ClassLowerName[42:46]:   ClassSite,
	_ClassName[46:50]:        ClassMqtt,
	_ClassLowerName[46:50]:   ClassMqtt,
	_ClassName[50:58]:        ClassDatabase,
	_ClassLowerName[50:58]:   ClassDatabase,
	_ClassName[58:69]:        ClassModbusProxy,
	_ClassLowerName[58:69]:   ClassModbusProxy,
	_ClassName[69:74]:        ClassEEBus,
	_ClassLowerName[69:74]:   ClassEEBus,
	_ClassName[74:84]:        ClassJavascript,
	_ClassLowerName[74:84]:   ClassJavascript,
	_ClassName[84:86]:        ClassGo,
	_ClassLowerName[84:86]:   ClassGo,
	_ClassName[86:90]:        ClassHEMS,
	_ClassLowerName[86:90]:   ClassHEMS,
	_ClassName[90:96]:        ClassInflux,
	_ClassLowerName[90:96]:   ClassInflux,
	_ClassName[96:105]:       ClassMessenger,
	_ClassLowerName[96:105]:  ClassMessenger,
	_ClassName[105:116]:      ClassSponsorship,
	_ClassLowerName[105:116]: ClassSponsorship,
	_ClassName[116:125]:      ClassLoadpoint,
	_ClassLowerName[116:125]: ClassLoadpoint,
}

var _ClassNames = []string{
	_ClassName[0:10],
	_ClassName[10:15],
	_ClassName[15:22],
	_ClassName[22:29],
	_ClassName[29:35],
	_ClassName[35:42],
	_ClassName[42:46],
	_ClassName[46:50],
	_ClassName[50:58],
	_ClassName[58:69],
	_ClassName[69:74],
	_ClassName[74:84],
	_ClassName[84:86],
	_ClassName[86:90],
	_ClassName[90:96],
	_ClassName[96:105],
	_ClassName[105:116],
	_ClassName[116:125],
}

// ClassString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ClassString(s string) (Class, error) {
	if val, ok := _ClassNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ClassNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Class values", s)
}

// ClassValues returns all values of the enum
func ClassValues() []Class {
	return _ClassValues
}

// ClassStrings returns a slice of all String values of the enum
func ClassStrings() []string {
	strs := make([]string, len(_ClassNames))
	copy(strs, _ClassNames)
	return strs
}

// IsAClass returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Class) IsAClass() bool {
	for _, v := range _ClassValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Class
func (i Class) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Class
func (i *Class) UnmarshalText(text []byte) error {
	var err error
	*i, err = ClassString(string(text))
	return err
}
