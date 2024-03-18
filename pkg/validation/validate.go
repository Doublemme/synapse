package validation

import (
	"fmt"
	"unicode"
)

type ValidationRule struct {
	Name         string
	FieldName    any
	RuleVal      any
	FieldVal     any
	ErrorMsgFunc func(ValidationRule) string
	ValidateFunc func(ValidationRule) bool
}

type ValidationRuleFunc func() ValidationRule

type Fields map[string][]ValidationRule

func BindingRules(rules ...ValidationRuleFunc) []ValidationRule {
	bindingRules := make([]ValidationRule, len(rules))
	for i, rule := range rules {
		bindingRules[i] = rule()
	}
	return bindingRules
}

type Validator struct {
	data   any
	fields Fields
}

func NewValidator(data any, fields Fields) *Validator {
	return &Validator{
		data:   data,
		fields: fields,
	}
}

func (v *Validator) Validate(stderrTarget any) bool {
	isValid := true

	for fieldName, rules := range v.fields {
		if !unicode.IsUpper(rune(fieldName[0])) {
			continue
		}
		fmt.Println(fieldName, rules)
	}
	return isValid
}
