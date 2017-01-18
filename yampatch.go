package yampatch

import (
	"github.com/cppforlife/go-patch/patch"

	yaml "gopkg.in/yaml.v2"
)

func ApplyOps(unpatchedYAML string, operationsYAML string) (string, error) {
	var unmarshaledYAML interface{}
	var unmarshaledPatchedYAML interface{}
	var opDefs []patch.OpDefinition

	err := yaml.Unmarshal([]byte(unpatchedYAML), &unmarshaledYAML)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal([]byte(operationsYAML), &opDefs)
	if err != nil {
		return "", err
	}

	ops, err := patch.NewOpsFromDefinitions(opDefs)
	if err != nil {
		return "", err
	}

	unmarshaledPatchedYAML, err = ops.Apply(unmarshaledYAML)
	if err != nil {
		return "", err
	}

	outBytes, err := yaml.Marshal(unmarshaledPatchedYAML)
	if err != nil {
		return "", err
	}

	return string(outBytes), nil
}
