package main

import (
	"fmt"
	"syscall/js"

	"github.com/GoogleContainerTools/kpt-functions-catalog/functions/go/set-namespace/transformer"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

func main() {
	fmt.Println("Hello WASM")
	js.Global().Set("transformNs", transformNamespaceWrapper())
	<-make(chan bool)
}

func transformNamespace(input string) (string, error) {
	output, err := fn.Run(fn.ResourceListProcessorFunc(transformer.SetNamespace), []byte(input))
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func transformNamespaceWrapper() js.Func {  
        jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
                if len(args) != 1 {
                        return "Invalid number of arguments passed"
                }
                input := args[0].String()
                fmt.Printf("input %s\n", input)
                transformed, err := transformNamespace(input)
                if err != nil {
                        fmt.Printf("unable to transform: %v\n", err)
                        return err.Error()
                }
                return transformed
        })
        return jsonFunc
}
