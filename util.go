package kuropi

import "errors"

func stringSliceContains(arr []string, s string) bool {
	for _, elt := range arr {
		if s == elt {
			return true
		}
	}
	return false
}

func applyDefinitionToContext(ctx Context, defs []Definition) error {
	for _, def := range defs {
		if err := ctx.AddDefinition(def); err != nil {
			return err
		}
	}
	return nil
}

func validateDefinition(def Definition) error {
	if def.Name == "" {
		return errors.New("Definition must have a name")
	}
	if def.Build == nil {
		return errors.New("Definition must provide build method")
	}
	return nil
}
