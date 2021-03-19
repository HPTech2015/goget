package tests

import (
	"testing"
	"goget/core"
	"os"
)

func TestPull(t *testing.T) {
	assert := Assert{}
	get := core.GetFile{}
	pwd, _ := os.Getwd() 

	// Test for good.
	localFilePath := pwd + "/OIP.-GK_c_18AXchHTMajvT1rgHaEM?pid=ImgDet&rs=1"
	err := get.Pull(localFilePath, "https://th.bing.com/th/id/OIP.-GK_c_18AXchHTMajvT1rgHaEM?pid=ImgDet&rs=1")
	if  err != nil {
		t.Errorf("Unexpected Error: " + err.Error())
	}

	file, err := os.Stat(localFilePath)
	if err != nil {
		t.Errorf("Unexpected Error: " + err.Error())
	}
	assert.AssertString(file.Name(), "OIP.-GK_c_18AXchHTMajvT1rgHaEM?pid=ImgDet&rs=1", t)

	os.Remove(localFilePath)

	// Test for bad.
	err = get.Pull(localFilePath, "https://fakedomain/fakefile")
	assert.AssertString(err.Error(), "Get \"https://fakedomain/fakefile\": dial tcp: lookup fakedomain: no such host", t)
}

func TestPullAndCheck(t *testing.T) {
	assert := Assert{}
	get := core.GetFile{}
	pwd, _ := os.Getwd() 

	// Test for good.
	localFilePath := pwd + "/test.txt"
	testSig, err := get.PullAndCheck(localFilePath, "http://localhost/test.txt")
	if  err != nil {
		t.Errorf("Unexpected Error: " + err.Error())
	}

	assert.AssertTrue(testSig, t)

	os.Remove(localFilePath)
	os.Remove(localFilePath + ".sha256")

	// Test for bad.
	_, err = get.PullAndCheck(localFilePath, "https://fakedomain/fakefile")
	assert.AssertString(err.Error(), "Get \"https://fakedomain/fakefile\": dial tcp: lookup fakedomain: no such host", t)

	localFilePath = pwd + "/test_bad.txt"
	testSig, err = get.PullAndCheck(localFilePath, "http://localhost/test_bad.txt")
	if  err != nil {
		t.Errorf("Unexpected Error: " + err.Error())
	}

	assert.AssertFalse(testSig, t)

	os.Remove(localFilePath)
	os.Remove(localFilePath + ".sha256")
}