package test

import (
	"Alagappan/file"
	"errors"
	"testing"
)





func TestReadFromTextFile(t *testing.T) {
	expected:=[]string{
		"https://reqres.in/api/users/3",
		"https://reqres.in/api/users/1",
		"https://reqres.in/api/users/2",
		"https://reqres.in/api/users?page=2",
		"https://reqres.in/api/users?page=1",
	}
	actual,_,_:=file.ParseFile("../files/fileWithProperURL.txt")

	if len(expected) != len(actual) {
		t.Errorf("All values in file not read properly")
	}

	for i,v:= range actual {
		if v != expected[i] {
			t.Errorf("Values not read in order")
		}
	}
}

func TestFilePathError(t *testing.T) {
	fileName:="../files/Nofile.txt"
	expectedErr:=errors.New("open "+fileName+": no such file or directory")
	expectedMessage:="Not able to open the file "+fileName
	_,message,err:=file.ParseFile(fileName)

	if expectedErr.Error()!=err.Error(){
		t.Errorf("Error message not proper")
	}

	if expectedMessage != message {
		t.Errorf("Custom message not proper")
	}
}

func TestReadFromCSVFile(t *testing.T) {
	expected:=[]string{
		"https://reqres.in/api/users/3",
		"https://reqres.in/api/users/1",
		"https://reqres.in/api/users/2",
		"https://reqres.in/api/users?page=2",
		"https://reqres.in/api/users?page=1",
	}
	actual,_,_:=file.ParseFile("../files/fileWithProperURL.csv")

	if len(expected) != len(actual) {
		t.Errorf("All values in file not read properly")
	}

	for i,v:= range actual {
		if v != expected[i] {
			t.Errorf("Values not read in order")
		}
	}
}

func TestEmptyLineReadFromCSVFile(t *testing.T) {
	expected:=[]string{
		"https://reqres.in/api/users/2",
		"",
		"https://reqres.in/api/users/1",
		"",
	}
	actual,_,_:=file.ParseFile("../files/emptyLines.csv")

	if len(expected) != len(actual) {
		t.Errorf("All values in file not read properly")
	}

	for i,v:= range actual {
		if v != expected[i] {
			t.Errorf("Values not read in order")
		}
	}
}

func TestFilePermissionError(t *testing.T) {
	fileName:="../files/NoPermissionFile.csv"
	expectedErr:=errors.New("open "+fileName+": permission denied")
	expectedMessage:="Not able to open the file "+fileName
	_,message,err:=file.ParseFile(fileName)

	if expectedErr.Error()!=err.Error(){
		t.Errorf("Error message not proper")
	}

	if expectedMessage != message {
		t.Errorf("Custom message not proper")
	}
}