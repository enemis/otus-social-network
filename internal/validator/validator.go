package validator

import (
	"otus-social-network/internal/model"

	"github.com/golodash/galidator"
)

func BuildValidator(input any) galidator.Validator {
	customValidators := make(map[string]func(interface{}) bool, 10)
	customValidators["post_status"] = postStatusValidate

	g := galidator.G().CustomMessages(galidator.Messages{
		"post_status":      "$value is not valid post's status",
		"required":         "$value is required",
		"required_without": "$field should be filled if other field empty",
	}).CustomValidators(customValidators)
	validator := g.Validator(input)

	return validator
}

func postStatusValidate(value interface{}) bool {
	_, ok := value.(model.PostStatus)
	return ok
}
