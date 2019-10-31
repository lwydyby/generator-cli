package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "latest_version", SnakeCase("lastest_version"))
	assert.Equal(t, "x_qs_date", SnakeCase("x-qs-date"))
	assert.Equal(t, "x_qs_date", SnakeCase("x-qs-date"))
	assert.Equal(t, "hello_world", SnakeCase("Hello World"))
	assert.Equal(t, "qingstor", SnakeCase("QingStor"))
	assert.Equal(t, "qingstor", SnakeCase("Qing Stor"))
}

func TestSnakeCaseToSnakeCase(t *testing.T) {
	assert.Equal(t, "latest_version", SnakeCaseToSnakeCase("lastest_version"))
	assert.Equal(t, "x_qs_date", SnakeCaseToSnakeCase("X-QS-Date", true))
	assert.Equal(t, "hello_world!", SnakeCaseToSnakeCase("Hello World!", true))
}

func TestSnakeCaseToCamelCase(t *testing.T) {
	assert.Equal(t, "CamelCase", SnakeCaseToCamelCase("camel_case"))
	assert.Equal(t, "BucketACL", SnakeCaseToCamelCase("bucket_acl"))
	assert.Equal(t, "AllowedOrigin", SnakeCaseToCamelCase("allowed_origin"))
	assert.Equal(t, "PartNumberMarker", SnakeCaseToCamelCase("part_number_marker"))
	assert.Equal(t, "CamelCase", SnakeCaseToCamelCase("camel_case"))
	assert.Equal(t, "CamelCase", SnakeCaseToCamelCase("camel_case"))
	assert.Equal(t, "CamelCase", SnakeCaseToCamelCase("camel_case"))
	assert.Equal(t, "XQSDate", SnakeCaseToCamelCase("X-QS-Date"))
}

func TestSnakeCaseToDashConnected(t *testing.T) {
	assert.Equal(t, "camel-case", SnakeCaseToDashConnected("camel_case"))
	assert.Equal(t, "bucket-acl", SnakeCaseToDashConnected("bucket_acl"))
}
