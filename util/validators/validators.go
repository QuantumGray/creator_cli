package validators

import (
	typeserrors "fluttercreator/util/typeserrors"
)

// CreateCommandArgsValidation : validates arguments from create command
func CreateCommandArgsValidation(args []string) error {
	if len(args[0]) != 8 {
		return typeserrors.ErrInvalidArgument
	}
	if len(args) != 1 {
		return typeserrors.ErrInvalidArgument
	}
	return nil
}

// AppNameValidation : validates app name to be a valid argument
func AppNameValidation(arg string) {

}
