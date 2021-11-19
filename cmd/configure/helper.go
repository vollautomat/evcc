package configure

import (
	"fmt"
	"sort"
	"strings"

	"github.com/evcc-io/evcc/templates"
	"gopkg.in/yaml.v3"
)

func (c *CmdConfigure) processDeviceSelection(deviceCategory DeviceCategory) (templates.Template, error) {
	templateItem := c.selectItem(deviceCategory)

	if templateItem.Description == c.localizedString("ItemNotPresent", nil) {
		return templateItem, c.errItemNotPresent
	}

	err := c.processDeviceRequirements(templateItem)
	if err != nil {
		return templateItem, err
	}

	return templateItem, nil
}

func (c *CmdConfigure) processDeviceValues(values map[string]interface{}, templateItem templates.Template, device device, deviceCategory DeviceCategory) (device, error) {
	c.addedDeviceIndex++

	device.Name = fmt.Sprintf("%s%d", DeviceCategories[deviceCategory].defaultName, c.addedDeviceIndex)
	device.Title = templateItem.Description
	for _, param := range templateItem.Params {
		if param.Name != "title" {
			continue
		}
		if len(param.Value) > 0 {
			device.Title = param.Value
		}
	}

	deviceIsValid := false
	v, err := c.configureDevice(deviceCategory, templateItem, values)
	if err == nil {
		fmt.Println()
		if deviceCategory == DeviceCategoryPVMeter || deviceCategory == DeviceCategoryBatteryMeter || deviceCategory == DeviceCategoryGridMeter {
			fmt.Println(c.localizedString("TestingDevice_TitleUsage", localizeMap{"Device": templateItem.Description, "Usage": deviceCategory.String()}))
		} else {
			fmt.Println(c.localizedString("TestingDevice_Title", localizeMap{"Device": templateItem.Description}))
		}

		deviceIsValid, err = c.testDevice(deviceCategory, v)
		if deviceCategory == DeviceCategoryCharger {
			if deviceIsValid && err == nil {
				device.ChargerHasMeter = true
			}
		}
	}

	if !deviceIsValid {
		c.addedDeviceIndex--
		return device, c.errDeviceNotValid
	}

	templateItem.Params = append(templateItem.Params, templates.Param{Name: "name", Value: device.Name})
	b, err := templateItem.RenderProxyWithValues(values, false)
	if err != nil {
		c.addedDeviceIndex--
		return device, err
	}

	device.Yaml = string(b)

	return device, nil
}

// handle device requirements
func (c *CmdConfigure) processDeviceRequirements(templateItem templates.Template) error {
	if len(templateItem.Requirements.Description.String(c.lang)) > 0 {
		fmt.Println()
		fmt.Println(c.localizedString("Requirements_Title", nil))
		fmt.Println("  ", templateItem.Requirements.Description.String(c.lang))
		if len(templateItem.Requirements.URI) > 0 {
			fmt.Println("  " + c.localizedString("Requirements_More", nil) + " " + templateItem.Requirements.URI)
		}
	}

	// check if sponsorship is required
	if templateItem.Requirements.Sponsorship && c.configuration.SponsorToken() == "" {
		fmt.Println()
		fmt.Println(c.localizedString("Requirements_Sponsorship_Title", nil))
		fmt.Println()
		if !c.askYesNo(c.localizedString("Requirements_Sponsorship_Token", nil)) {
			return c.errItemNotPresent
		}
		sponsortoken := c.askValue(question{
			label:    c.localizedString("Requirements_Sponsorship_Token_Input", nil),
			help:     "",
			required: true})
		c.configuration.SetSponsorToken(sponsortoken)
	}

	// check if we need to setup an EEBUS HEMS
	if templateItem.Requirements.Eebus {
		if c.configuration.EEBUS() == "" {
			eebusConfig, err := c.eebusCertificate()

			if err != nil {
				return fmt.Errorf("%s: %s", c.localizedString("Requirements_EEBUS_Cert_Error", nil), err)
			}

			err = c.configureEEBus(eebusConfig)
			if err != nil {
				return err
			}

			eebusYaml, err := yaml.Marshal(eebusConfig)
			if err != nil {
				return err
			}
			c.configuration.SetEEBUS(string(eebusYaml))
		}

		fmt.Println()
		fmt.Println(c.localizedString("Requirements_EEBUS_Pairing", nil))
		fmt.Scanln()
	}

	return nil
}

// return template items of a given class
func (c *CmdConfigure) fetchElements(deviceCategory DeviceCategory) []templates.Template {
	var items []templates.Template

	for _, tmpl := range templates.ByClass(DeviceCategories[deviceCategory].class.String()) {
		if len(tmpl.Params) == 0 || len(tmpl.Description) == 0 {
			continue
		}

		if deviceCategory == DeviceCategoryGuidedSetup {
			if tmpl.GuidedSetup.Enable {
				items = append(items, tmpl)
			}
		} else {
			if len(DeviceCategories[deviceCategory].categoryFilter) == 0 ||
				c.paramChoiceContains(tmpl.Params, templates.ParamUsage, DeviceCategories[deviceCategory].categoryFilter.String()) {
				items = append(items, tmpl)
			}
		}
	}

	sort.Slice(items[:], func(i, j int) bool {
		// sort generic templates to the bottom
		if items[i].Generic && !items[j].Generic {
			return false
		}
		if !items[i].Generic && items[j].Generic {
			return true
		}
		return strings.ToLower(items[i].Description) < strings.ToLower(items[j].Description)
	})

	return items
}

// helper function to check if a param choice contains a given value
func (c *CmdConfigure) paramChoiceContains(params []templates.Param, name, filter string) bool {
	nameFound, choices := c.paramChoiceValues(params, name)

	if !nameFound {
		return false
	}

	for _, choice := range choices {
		if string(choice) == filter {
			return true
		}
	}

	return false
}

func (c *CmdConfigure) paramChoiceValues(params []templates.Param, name string) (bool, []DeviceCategory) {
	nameFound := false

	choices := []DeviceCategory{}

	for _, item := range params {
		if item.Name != name {
			continue
		}

		nameFound = true

		for _, choice := range item.Choice {
			choices = append(choices, DeviceCategory(choice))
		}
	}

	return nameFound, choices
}

// Process an EVCC configuration item
// Returns
//   a map with param name and values
func (c *CmdConfigure) processConfig(paramItems []templates.Param, deviceCategory DeviceCategory, includeAdvanced bool) map[string]interface{} {
	usageFilter := DeviceCategories[deviceCategory].categoryFilter

	additionalConfig := make(map[string]interface{})
	selectedModbusKey := ""

	fmt.Println()
	fmt.Println(c.localizedString("Config_Title", nil))
	fmt.Println()

	for _, param := range paramItems {
		if param.Name == "modbus" {
			choices := []string{}
			choiceKeys := []string{}
			for _, choice := range param.Choice {
				switch choice {
				case ModbusChoiceRS485:
					choices = append(choices, "Serial (USB-RS485 Adapter)")
					choiceKeys = append(choiceKeys, ModbusKeyRS485Serial)
					choices = append(choices, "Serial (Ethernet-RS485 Adapter)")
					choiceKeys = append(choiceKeys, ModbusKeyRS485TCPIP)
				case ModbusChoiceTCPIP:
					choices = append(choices, "TCP/IP")
					choiceKeys = append(choiceKeys, ModbusKeyTCPIP)
				}
			}

			if len(choices) > 0 {
				// ask for modbus address
				id := c.askValue(question{
					label:        "ID",
					help:         "Modbus ID",
					defaultValue: 1,
					valueType:    templates.ParamValueTypeNumber,
					required:     true})
				additionalConfig[ModbusParamNameId] = id

				// ask for modbus interface type
				index := 0
				if len(choices) > 1 {
					index, _ = c.askChoice(c.localizedString("Config_ModbusInterface", nil), choices)
				}
				selectedModbusKey = choiceKeys[index]
				switch selectedModbusKey {
				case ModbusKeyRS485Serial:
					device := c.askValue(question{
						label:        "Device",
						help:         "USB-RS485 Adapter Adresse",
						exampleValue: ModbusParamValueDevice,
						required:     true})
					additionalConfig[ModbusParamNameDevice] = device

					baudrate := c.askValue(question{
						label:        "Baudrate",
						defaultValue: ModbusParamValueBaudrate,
						valueType:    templates.ParamValueTypeNumber,
						required:     true})
					additionalConfig[ModbusParamNameBaudrate] = baudrate

					comset := c.askValue(question{
						label:        "ComSet",
						defaultValue: ModbusParamValueComset,
						required:     true})
					additionalConfig[ModbusParamNameComset] = comset

				case ModbusKeyRS485TCPIP, ModbusKeyTCPIP:
					if selectedModbusKey == ModbusKeyRS485TCPIP {
						additionalConfig[ModbusParamNameRTU] = "true"
					}
					host := c.askValue(question{
						label:        "Host",
						exampleValue: ModbusParamValueHost,
						required:     true})
					additionalConfig[ModbusParamNameHost] = host

					port := c.askValue(question{
						label:        "Port",
						defaultValue: ModbusParamValuePort,
						valueType:    templates.ParamValueTypeNumber,
						required:     true})
					additionalConfig[ModbusParamNamePort] = port
				}
			}
		} else if param.Name != templates.ParamUsage {
			if !includeAdvanced && param.Advanced {
				continue
			}

			userFriendly := c.userFriendlyTexts(param)
			value := c.askValue(question{
				label:        userFriendly.Name,
				defaultValue: userFriendly.Default,
				exampleValue: userFriendly.Example,
				help:         userFriendly.Help.String(c.lang),
				valueType:    userFriendly.ValueType,
				mask:         userFriendly.Mask,
				required:     userFriendly.Required})
			additionalConfig[param.Name] = value
		} else if param.Name == templates.ParamUsage {
			if usageFilter != "" {
				additionalConfig[param.Name] = usageFilter.String()
			}
		}
	}

	return additionalConfig
}
