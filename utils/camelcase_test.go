package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	assert.Equal(t, "EIP", CamelCase("Eip"))
	assert.Equal(t, "DNSAlias", CamelCase("DNS Alias"))
	assert.Equal(t, "DNSAlias", CamelCase("dns_alias"))
	assert.Equal(t, "HelloWorld", CamelCase("hello_world"))
	assert.Equal(t, "QingStor", CamelCase("qingstor"))
}

func TestCamelCaseToCamelCase(t *testing.T) {
	assert.Equal(t, "EIP", CamelCaseToCamelCase("Eip"))
	assert.Equal(t, "DNS Alias", CamelCaseToCamelCase("DNS Alias"))
	assert.Equal(t, "DNSAlias", CamelCaseToCamelCase("DNS Alias", true))
	assert.Equal(t, "Describe-VxNets", CamelCaseToCamelCase("Describe-Vxnets"))
	assert.Equal(t, "DescribeVxNets", CamelCaseToCamelCase("Describe-Vxnets", true))
}

func TestCamelCaseToSnakeCase(t *testing.T) {
	assert.Equal(t, "snake_case", CamelCaseToSnakeCase("SnakeCase"))
	assert.Equal(t, "x_qs_date", CamelCaseToSnakeCase("XQSDate"))
	assert.Equal(t, "bucket_acl", CamelCaseToSnakeCase("BucketACL"))
	assert.Equal(t, "cpu", CamelCaseToSnakeCase("CPU"))
	assert.Equal(t, "vxnets", CamelCaseToSnakeCase("VxNets"))
}

func TestConvertCamelCaseToDashConnected(t *testing.T) {
	a := "/<bucket-name>?stats"
	assert.Equal(t, a, CamelCaseToDashConnected("/<bucketName>?stats"))

	b := "/<bucket-name>/<object-key>"
	assert.Equal(t, b, CamelCaseToDashConnected("/<bucketName>/<objectKey>"))

	c := "/create-key-pair"
	assert.Equal(t, c, CamelCaseToDashConnected("/CreateKeyPair"))
}

func TestLowerFirstCharacter(t *testing.T) {
	assert.Equal(t, LowerFirstCharacter("ABC"), "aBC")
}

func TestUpperFirstWord(t *testing.T) {
	assert.Equal(t, LowerFirstWord("CPUModel"), "cpuModel")
}

func TestUpperFirstCharacter(t *testing.T) {
	assert.Equal(t, UpperFirstCharacter("aBC"), "ABC")
}
