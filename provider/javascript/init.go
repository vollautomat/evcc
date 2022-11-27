package javascript

import (
	"fmt"
	"log"
	"strings"

	"github.com/evcc-io/evcc/util"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/samber/lo"
)

// Configure initializes JS VMs
func Configure(other map[string]interface{}) error {
	cc := []struct {
		VM     string
		Script string
	}{}

	if err := util.DecodeOther(other, &cc); err != nil {
		return err
	}

	// init all VMs that require it
	for _, conf := range cc {
		if conf.Script == "" {
			continue
		}

		if _, ok := registry[conf.VM]; !ok {
			vm := otto.New()

			_, err := vm.Run(conf.Script)
			if err != nil {
				return err
			}

			registry[conf.VM] = vm
		}
	}

	return nil
}

var registry = make(map[string]*otto.Otto)

// RegisteredVM returns a JS VM. If name is not empty, it will return a shared instance.
func RegisteredVM(name string) *otto.Otto {
	vm, ok := registry[name]

	// create new VM
	if !ok {
		vm = otto.New()
		setConsole(vm, name)

		if name != "" {
			registry[name] = vm
		}
	}

	return vm
}

func setConsole(vm *otto.Otto, name string) {
	if name == "" {
		name = "javascript"
	}

	log := util.NewLogger(name)

	console := map[string]any{
		"trace": printer(log.TRACE),
		"log":   printer(log.DEBUG),
		"info":  printer(log.INFO),
		"error": printer(log.ERROR),
	}

	if err := vm.Set("console", console); err != nil {
		panic(err)
	}

	// vm.Run(`console.log("Hello, World.");`)
	// os.Exit(1)
}

func printer(log *log.Logger) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		output := lo.Map(call.ArgumentList, func(a otto.Value, _ int) string {
			return fmt.Sprintf("%v", a)
		})

		log.Println(strings.Join(output, " "))

		return otto.UndefinedValue()
	}
}
