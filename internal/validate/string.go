package validate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// stringLenBetweenValidator validates that a string Attribute's length is in a range.
type stringLenBetweenValidator struct {
	tfsdk.AttributeValidator

	minLength, maxLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be between %d and %d", validator.minLength, validator.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	l = len(s.Value)

	if l < validator.minLength || l > validator.maxLength {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be in the range [%d, %d], got %d", validator.minLength, validator.maxLength, l),
		)

		return
	}
}

// StringLenBetween returns a new string length between validator.
func StringLenBetween(minLength, maxLength int) tfsdk.AttributeValidator {
	if minLength < 0 || maxLength < 0 || minLength > maxLength {
		return nil
	}

	return stringLenBetweenValidator{
		minLength: minLength,
		maxLength: maxLength,
	}
}

// stringLenAtLeastValidator validates that a string Attribute's length is at least a certain value.
type stringLenAtLeastValidator struct {
	tfsdk.AttributeValidator

	minLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be at least %d", validator.minLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	l = len(s.Value)

	if l < validator.minLength {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be at least %d, got %d", validator.minLength, l),
		)

		return
	}
}

// StringLenAtLeast returns a new string length at least validator.
func StringLenAtLeast(minLength int) tfsdk.AttributeValidator {
	if minLength < 0 {
		return nil
	}

	return stringLenAtLeastValidator{
		minLength: minLength,
	}
}

// stringLenAtMostValidator validates that a string Attribute's length is at most a certain value.
type stringLenAtMostValidator struct {
	tfsdk.AttributeValidator

	maxLength int
}

// Description describes the validation in plain text formatting.
func (validator stringLenAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("string length must be at most %d", validator.maxLength)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringLenAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringLenAtMostValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	l = len(s.Value)

	if l > validator.maxLength {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be at most %d, got %d", validator.maxLength, l),
		)

		return
	}
}

// StringLenAtMost returns a new string length at least validator.
func StringLenAtMost(maxLength int) tfsdk.AttributeValidator {
	if maxLength < 0 {
		return nil
	}

	return stringLenAtMostValidator{
		maxLength: maxLength,
	}
}

// stringInSliceValidator validates that a string Attribute's value matches the value of an element in the valid slice.
type stringInSliceValidator struct {
	tfsdk.AttributeValidator

	valid []string
}

// Description describes the validation in plain text formatting.
func (validator stringInSliceValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be one of %v", validator.valid)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringInSliceValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringInSliceValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	for _, val := range validator.valid {
		if s.Value == val {
			return
		}
	}

	response.Diagnostics.AddAttributeError(
		request.AttributePath,
		"Invalid value",
		fmt.Sprintf("expected value to be one of %v, got %s", validator.valid, s.Value),
	)
}

// StringInSlice returns a new string in slice validator.
func StringInSlice(valid []string) tfsdk.AttributeValidator {
	return stringInSliceValidator{
		valid: valid,
	}
}

// stringIsJsonObjectValidator validates that a string Attribute's value is a valid JSON object.
type stringIsJsonObjectValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator stringIsJsonObjectValidator) Description(_ context.Context) string {
	return "value must be a valid JSON object"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator stringIsJsonObjectValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator stringIsJsonObjectValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	// A JSON object starts with a '{'
	if s.Value[:1] != "{" {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"expected value to be a valid JSON object",
		)
		return
	}

	// Use json.Unmarshal() instead of just json.Valid() to get the parsing error
	var i interface{}
	err := json.Unmarshal([]byte(s.Value), &i)
	if err != nil {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			fmt.Sprintf("expected value to be valid JSON: %s", err),
		)
	}

}

// StringIsJsonObject returns a new string is JSON validator.
func StringIsJsonObject() tfsdk.AttributeValidator {
	return stringIsJsonObjectValidator{}
}
