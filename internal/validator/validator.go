package validator

import "github.com/golodash/galidator"

func BuildValidator(input any) galidator.Validator {
	g := galidator.G().CustomMessages(galidator.Messages{
		"required": "$value is not string",
	})
	validator := g.Validator(input)

	return validator
}
