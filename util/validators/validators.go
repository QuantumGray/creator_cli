package validators

import (
	"github.com/QuantumGray/fluttercreator_cli/typeserrors"
)

func CreateCommandArgsValidation(args []string) error {
	if len(args[0]) != 8 {
		return typeserrors.InvalidArgumentErr
	}
	if len(args) != 1 {
		return typeserrors.InvalidArgumentErr
	}
	return nil
}
