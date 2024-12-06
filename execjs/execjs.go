package execjs

import (
	_ "embed"
	"fmt"

	"github.com/dop251/goja"
	"github.com/robertkrimen/otto"
	v8 "rogchap.com/v8go"
)

//go:embed jsonpath.js
var jsonPathLib []byte

//go:embed example.json
var versionFile []byte

// NO! fails due to "Fatal javascript OOM in CALL_AND_RETRY_LAST" when executed in loop more then 1000 times
func ExecJSONPathWithV8(jsonPath string, json string) (string, error) {
	ctx := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	ctx.RunScript(string(jsonPathLib), "init-lib.js")

	script := fmt.Sprintf(`
const result = this.JSONPath({path: '%s', json: %s});
if (result && result.length > 0) result[0];
`, jsonPath, json)

	// fmt.Printf("script: %s", script)
	val, err := ctx.RunScript(script, "run-jsonpath.js")
	if err != nil {
		fmt.Printf("error running script: %s", err)
		return "", err
	}
	return val.String(), nil
}

// NO! doesn't support keywords like 'class'
func ExecJSONPathWithOtto(jsonPath string, json string) (string, error) {
	script := fmt.Sprintf(`
	%s
	var result = JSONPath({path: '%s', json: %s});
	var returned
	if (result && result.length > 0){
		returned = result[0];
	}
	`, string(jsonPathLib), jsonPath, json)

	// fmt.Printf("script: %s", script)

	vm := otto.New()
	_, err := vm.Run(script)
	if err != nil {
		return "", err
	}

	var returned string

	if value, err := vm.Get("returned"); err == nil {
		if value_str, err := value.ToString(); err == nil {
			returned = value_str
		} else {
			return "", err
		}
	} else {
		return "", err
	}

	return returned, nil
}

// YES! works fine
func ExecJSONPathWithGoJa(jsonPath string, json string) (string, error) {
	script := fmt.Sprintf(`
%s
const result = this.JSONPath({path: '%s', json: %s});
if (result && result.length > 0) result[0];
`, jsonPathLib, jsonPath, json)

	// fmt.Printf("script: %s", script)

	vm := goja.New()
	v, err := vm.RunString(script)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", v.Export()), nil
}

func GetJSONFileContent() []byte {
	return versionFile
}
