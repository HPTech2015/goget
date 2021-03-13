package tests

import (
	"testing"
	"goget/core"
	"os"
)

func TestGetFile(t *testing.T) {
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