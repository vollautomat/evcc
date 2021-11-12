package configure

import (
	"fmt"

	"github.com/evcc-io/evcc/charger"
	"github.com/evcc-io/evcc/meter"
	"github.com/evcc-io/evcc/templates"
	"github.com/evcc-io/evcc/vehicle"
	"gopkg.in/yaml.v3"
)

func (c *CmdConfigure) configureDeviceSingleSetup() {
	var repeat bool = true
	var err error

	var values map[string]interface{}
	var deviceCategory string
	var supportedDeviceCategories []string
	var templateItem templates.Template

	deviceItem := device{}

	for ok := true; ok; ok = repeat {
		fmt.Println()

		templateItem, err = c.handleDeviceSelection(DeviceCategorySingleSetup)
		if err != nil {
			return
		}

		usageFound, usageChoices := c.paramChoiceValues(templateItem.Params, templates.ParamUsage)
		if !usageFound {
			fmt.Println("error")
			return
		}
		if len(usageChoices) == 0 {
			usageChoices = []string{UsageChoiceGrid, UsageChoicePV, UsageChoiceBattery}
		}

		supportedDeviceCategories = []string{}

		for _, usage := range usageChoices {
			switch usage {
			case UsageChoiceGrid:
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryGridMeter)
			case UsageChoicePV:
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryPVMeter)
			case UsageChoiceBattery:
				supportedDeviceCategories = append(supportedDeviceCategories, DeviceCategoryBatteryMeter)
			}
		}

		// we only ask for the configuration for the first usage
		deviceCategory = supportedDeviceCategories[0]

		values := c.processConfig(templateItem.Params, deviceCategory, false)

		deviceItem, err = c.processDeviceValues(values, templateItem, deviceItem, deviceCategory)
		if err != nil {
			if err != ErrDeviceNotValid {
				fmt.Println("Fehler: ", err)
			}
			fmt.Println()
			if !c.askConfigFailureNextStep() {
				fmt.Println()
				return
			}
			continue
		}

		repeat = false
	}

	c.addDeviceToConfiguration(deviceItem, deviceCategory)

	for _, deviceCategory = range supportedDeviceCategories[1:] {
		deviceItem, err := c.processDeviceValues(values, templateItem, deviceItem, deviceCategory)
		if err != nil {
			continue
		}

		c.addDeviceToConfiguration(deviceItem, deviceCategory)
	}

	fmt.Println("Erfolgreich hinzugefügt.")

	c.handleLinkedTypes(templateItem)
}

func (c *CmdConfigure) handleLinkedTypes(templateItem templates.Template) {
	var repeat bool = true

	linkedTemplates := c.paramUsageLinkedType(templateItem.Params)

	if linkedTemplates == nil {
		return
	}

	for _, linkedTemplate := range linkedTemplates {
		for ok := true; ok; ok = repeat {
			deviceItem := device{}

			linkedTemplateItem := templates.ByType(linkedTemplate.Type, DeviceClassMeter)
			if len(linkedTemplateItem.Params) == 0 || linkedTemplate.Usage == "" {
				return
			}

			if !c.askYesNo("Möchtest du " + DeviceCategories[linkedTemplate.Usage].article + " " + linkedTemplateItem.Description + " " + DeviceCategories[linkedTemplate.Usage].title + " hinzufügen") {
				repeat = false
				continue
			}

			values := c.processConfig(linkedTemplateItem.Params, linkedTemplate.Usage, false)
			deviceItem, err := c.processDeviceValues(values, linkedTemplateItem, deviceItem, linkedTemplate.Usage)
			if err != nil {
				if err != ErrDeviceNotValid {
					fmt.Println("Fehler: ", err)
				}
				fmt.Println()
				if c.askConfigFailureNextStep() {
					continue
				}

			} else {
				c.addDeviceToConfiguration(deviceItem, linkedTemplate.Usage)

				fmt.Println("Erfolgreich hinzugefügt.")
			}
			repeat = false
		}
		repeat = true
	}
}

func (c *CmdConfigure) configureDeviceCategory(deviceCategory string) (device, error) {
	fmt.Println()
	fmt.Printf("- %s konfigurieren\n", DeviceCategories[deviceCategory].title)

	var repeat bool = true

	device := device{
		Name:  DeviceCategories[deviceCategory].defaultName,
		Title: "",
		Yaml:  "",
	}

	for ok := true; ok; ok = repeat {
		fmt.Println()

		templateItem, err := c.handleDeviceSelection(deviceCategory)
		if err != nil {
			return device, ErrItemNotPresent
		}

		values := c.processConfig(templateItem.Params, deviceCategory, false)

		device, err := c.processDeviceValues(values, templateItem, device, deviceCategory)
		if err != nil {
			if err != ErrDeviceNotValid {
				fmt.Println("Fehler: ", err)
			}
			fmt.Println()
			if !c.askConfigFailureNextStep() {
				fmt.Println()
				return device, err
			}
			continue
		}

		repeat = false
	}

	c.addDeviceToConfiguration(device, deviceCategory)

	fmt.Println("Erfolgreich hinzugefügt.")

	return device, nil
}

// create a configured device from a template so we can test it
func (c *CmdConfigure) configureDevice(deviceCategory string, device templates.Template, values map[string]interface{}) (interface{}, error) {
	b, err := device.RenderResult(false, values)
	if err != nil {
		return nil, err
	}

	var instance struct {
		Type  string
		Other map[string]interface{} `yaml:",inline"`
	}

	if err := yaml.Unmarshal(b, &instance); err != nil {
		return nil, err
	}

	var v interface{}

	switch DeviceCategories[deviceCategory].class {
	case DeviceClassMeter:
		v, err = meter.NewFromConfig(instance.Type, instance.Other)
	case DeviceClassCharger:
		v, err = charger.NewFromConfig(instance.Type, instance.Other)
	case DeviceClassVehicle:
		v, err = vehicle.NewFromConfig(instance.Type, instance.Other)
	}
	if err != nil {
		return nil, err
	}

	return v, nil
}
