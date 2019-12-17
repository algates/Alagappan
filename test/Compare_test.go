package test

import (
	"Alagappan/compare"
	"testing"
)

func TestBoolDiff(t *testing.T) {
	json1:= `{
 "isEnabled": true 
}`
	json2:= `{
 "isEnabled": false 
}`
	expected:=false
actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

if actual!=expected {
	t.Fail()
}

}

func TestBoolCheck(t *testing.T) {
	json1:= `{
 "isEnabled": true 
}`
	json2:= `{
 "isEnabled": true 
}`
	expected:=true
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}

func TestNumberDiff(t *testing.T) {
	json1:= `{
 "id": 1 
}`
	json2:= `{
 "id": 2 
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}

}

func TestNumberCheck(t *testing.T) {
	json1:= `{
 "id": 1 
}`
	json2:= `{
 "id": 1 
}`
	expected:=true
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}

func TestStringDiff(t *testing.T) {
	json1:= `{
 "id": "1" 
}`
	json2:= `{
 "id": "2" 
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}

}

func TestStringCheck(t *testing.T) {
	json1:= `{
 "id": "1" 
}`
	json2:= `{
 "id": "1" 
}`
	expected:=true
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}

func TestArrayLengthDiff(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":1
},
{
	"id":2
}
] 
}`
	json2:= `{
 "id": [
{
	"id":1
}
] 
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}

}

func TestArrayValueDiff(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":1
}
] 
}`
	json2:= `{
 "id": [
{
	"id":3
}
] 
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}

}

func TestTypeDiff(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":1
}
] 
}`
	json2:= `{
 "id":{
	"id":3
}
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}

}

func TestArrayOrderCheck(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":2
},
{
	"id":1
},
] 
}`
	json2:= `{
 "id": [
{
	"id":1
},
{
	"id":2
}
] 
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":1
},
{
	"id":2
}
] 
}`
	json2:= `{
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}

func TestInvalidJson(t *testing.T) {
	json1:= `{
 "id": [
{
	"id":1
},
{
	"id":2,
}
] 
}`
	json2:= `{
}`
	expected:=false
	actual:=compare.CompareJSON([]byte(json1),[]byte(json2))

	if actual!=expected {
		t.Fail()
	}
}