package tests

import (
	"testing"
	"os"
)

/*
	A struct that contains methods to help
	make unit and functional tests legible.
*/
type Assert struct {
}

/*
	Test the expected string is equal to the test string.
*/
func (assert *Assert) AssertString(str1, str2 string, t *testing.T) bool {
	if str1 != str2 {
		t.Errorf("Expected %v, but got %v.", str2, str1)
		return false
	}

	return true
}

/*
	Test the expected string is not equal to the test string.
*/
func (assert *Assert) AssertNotString(str1, str2 string, t *testing.T) bool {
	if str1 == str2 {
		t.Errorf("Did not expect %v, but got %v.", str2, str1)
		return false
	}

	return true
}

/*
	Test the expected integer is equal to the test integer.
*/
func (assert *Assert) AssertInteger(i1, i2 int, t *testing.T) bool {
	if i1 != i2 {
		t.Errorf("Expected %v, but got %v.", i2, i1)
		return false
	}

	return true
}

/*
	Test the expected integer is not equal to the test integer.
*/
func (assert *Assert) AssertNotInteger(i1, i2 int, t *testing.T) bool {
	if i1 == i2 {
		t.Errorf("Did not expect %v, but got %v.", i2, i1)
		return false
	}

	return true
}

/*
	Test the expected length of a string slice is not equal to the test length.
*/
func (assert *Assert) AssertSliceLen(strs []string, l int, t *testing.T) bool {
	if len(strs) != l {
		t.Errorf("Expected %v, but got %v.", l, len(strs))
		return false
	}

	return true
}

/*
	Test if boolean is true.
*/
func (assert *Assert) AssertTrue(testBool bool, t *testing.T) bool {
	if !testBool {
		t.Error("Expected true, but got false.")
		return testBool
	}

	return testBool
}

/*
	Test if boolean is false.
*/
func (assert *Assert) AssertFalse(testBool bool, t *testing.T) bool {
	if testBool {
		t.Error("Expected false, but got true.")
		return !testBool
	}

	return !testBool
}

/*
	Test if file is exists.
*/
func (assert *Assert) AssertFileExists(filePath string, t *testing.T) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Errorf("The %v file does not exist.", filePath)
		return false
	}

	return true
}

/*
	Test if file exists.
*/
func (assert *Assert) AssertNotFileExists(filePath string, t *testing.T) bool {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("The %v file exists.", filePath)
		return false
	}

	return true
}