package validator

import "github.com/golodash/galidator"

func BuildValidator(input any) galidator.Validator {
	g := galidator.G().CustomMessages(galidator.Messages{
		"required":         "$value is required",
		"required_without": "$field should be filled if other field empty",
	})
	validator := g.Validator(input)

	return validator
}
