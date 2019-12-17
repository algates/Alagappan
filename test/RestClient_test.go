package test

import (
	"Alagappan/request"
	"errors"
	"testing"
)

func TestGetRequest(t *testing.T) {

	expectedResp:=`{
  "data": {
    "id": 2,
    "email": "janet.weaver@reqres.in",
    "first_name": "Janet",
    "last_name": "Weaver",
    "avatar": "https://s3.amazonaws.com/uifaces/faces/twitter/josephstein/128.jpg"
  }
}`

	actualRespBytes,actualErr:=request.MakeHTTPRequest("https://reqres.in/api/users/2")

	if actualErr!= nil && string(actualRespBytes) != expectedResp {
		t.Fail()
	}
}

func TestMalformedRequest(t *testing.T) {

	expectedResp:=""
	expectedError := errors.New("dial tcp :80: connect: connection refused")
	actualRespBytes,actualErr:=request.MakeHTTPRequest("reqres.in/api/users/2")
	if actualErr.Error()!= expectedError.Error() || string(actualRespBytes) != expectedResp {
		t.Fail()
	}
}

func TestResponseCodeNot200(t *testing.T) {

	expectedResp:=`{}`
	expectedError:=errors.New("Response code not 200 404")
	actualRespBytes,actualErr:=request.MakeHTTPRequest("https://reqres.in/api/v1/users1")
	if actualErr.Error()!= expectedError.Error() || string(actualRespBytes) != expectedResp {
		t.Fail()
	}
}


