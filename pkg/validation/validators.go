package validation

import "fmt"

func Required() ValidationRule {
	return ValidationRule{
		Name: "required",
		ErrorMsgFunc: func(rule ValidationRule) string {
			return fmt.Sprintf("%s field is required", rule.FieldName)
		},
		ValidateFunc: func(vr ValidationRule) bool {
			str, ok := vr.FieldVal.(string)
			if !ok {
				return false
			}
			return len(str) > 0
		},
	}
}

func EqualTo(value any) ValidationRuleFunc {
	return func() ValidationRule {
		return ValidationRule{
			Name: "equal",
			ErrorMsgFunc: func(vr ValidationRule) string {
				return fmt.Sprintf("%s and %s are not equal", vr.FieldVal, value)
			},
			ValidateFunc: func(vr ValidationRule) bool {
				return vr.FieldVal == value
			},
		}
	}
}
