package tests

import (
	"testing"
	"goget/core"
	"os"
)

func TestParseArgChar(t *testing.T) {
	assert := Assert{}
	argParser := core.ArgParser{}

	// Test for good.
	argStrs, _ := argParser.ParseArgChar("-vio")

	assert.AssertString(argStrs[0], "--version", t)
	assert.AssertString(argStrs[1], "--input-file", t)
	assert.AssertString(argStrs[2], "--output-file", t)

	assert.AssertSliceLen(argStrs, 3, t)

	// Test for bad.
	argStrs, _ = argParser.ParseArgChar("-iov")

	assert.AssertNotString(argStrs[0], "--version", t)
	assert.AssertNotString(argStrs[1], "--input-file", t)
	assert.AssertNotString(argStrs[2], "--output-file", t)

	argStrs, _ = argParser.ParseArgChar("-abc")

	assert.AssertSliceLen(argStrs, 0, t)
}

func TestArgInvoke(t *testing.T) {
	assert := Assert{}
	settings := core.Settings{}
	argParser := core.ArgParser{}

	// Test for good.
	pwd, _ := os.Getwd()

	argParser.ArgInvoke(1, "--input-file=http://fakedomain/fakepath", &settings)
	assert.AssertString(settings.LocalTarget, pwd + "/fakepath", t)

	argParser.ArgInvoke(1, "--output-file=./", &settings)
	assert.AssertString(settings.LocalTarget, pwd + "/fakepath", t)

	argParser.ArgInvoke(1, "--output-file=/fakepath", &settings)

	assert.AssertString(settings.RemoteTarget, "http://fakedomain/fakepath", t)
	assert.AssertString(settings.LocalTarget, "/fakepath", t)

	// Test for bad.
	err := argParser.ArgInvoke(1, "--fake-arg", &settings)
	assert.AssertString(err.Error(), "Argument --fake-arg does not exist.", t)
}

func TestExtractArgVal(t *testing.T) {
	assert := Assert{}
	argParser := core.ArgParser{}

	// Test for good.
	argVal, _ := argParser.ExtractArgVal(1, "--input-file", "--input-file=http://fakedomain/fakepath")

	assert.AssertString(argVal, "http://fakedomain/fakepath", t)

	// Test for bad.
	_, err := argParser.ExtractArgVal(20, "--input-file", "--input-file")

	assert.AssertString(err.Error(), "--input-file requires a value.", t)
}