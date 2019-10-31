package utils

import "strings"

// SnakeCase converts a string to snake case string.
func SnakeCase(original string) string {
	if strings.ToLower(original) != original {
		original = CamelCaseToSnakeCase(original)
	}
	return SnakeCaseToSnakeCase(original, true)
}

// SnakeCaseToSnakeCase converts SnakeCase to SnakeCase.
func SnakeCaseToSnakeCase(original string, clean ...bool) string {
	converted := original

	if len(clean) == 1 && clean[0] {
		converted = strings.Replace(converted, " ", "_", -1)
		converted = strings.Replace(converted, "-", "_", -1)
		converted = strings.ToLower(converted)
	}

	convertedParts := strings.Split(converted, "_")

	for index, part := range convertedParts {
		if part != "" {
			if word := lowercaseToLowercaseWordMap[part]; word != "" {
				part = word
			}
		}
		convertedParts[index] = part
	}

	return strings.Join(convertedParts, "_")
}

// SnakeCaseToCamelCase converts SnakeCase to CamelCase.
func SnakeCaseToCamelCase(original string) string {
	converted := SnakeCaseToSnakeCase(original, true)
	convertedParts := strings.Split(converted, "_")

	for index, part := range convertedParts {
		if part != "" {
			if word := lowercaseToCapitalizedWordMap[part]; word != "" {
				part = word
			} else {
				part = strings.ToUpper(string(part[0])) + part[1:]
			}
		}
		convertedParts[index] = part
	}

	return strings.Join(convertedParts, "")
}

// SnakeCaseToDashConnected converts SnakeCase to DashConnected.
func SnakeCaseToDashConnected(original string) string {
	return strings.Replace(SnakeCaseToSnakeCase(original, true), "_", "-", -1)
}
