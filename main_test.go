package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRegexMatch1(t *testing.T){
	want:= NewCoeff(1,1,3)
	got := EmptyNewCoeff()

	got.ValidateEquation("1x+1y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch2(t *testing.T){
	want:= NewCoeff(1,-1,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("1x-1y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
	
}

func TestRegexMatch3(t *testing.T){
	want:= NewCoeff(1,1,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("x+y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch4(t *testing.T){
	want:= NewCoeff(1,-1,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("x-y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch5(t *testing.T){
	want:= NewCoeff(1,0,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("x=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch6(t *testing.T){
	want:= NewCoeff(0,1,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch7(t *testing.T){
	want:= NewCoeff(-1,0,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("-x=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestRegexMatch8(t *testing.T){
	want:= NewCoeff(-21,0,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("-21x=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}
func TestRegexMatch9(t *testing.T){
	want:= NewCoeff(0,-1,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("-y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}
func TestRegexMatch10(t *testing.T){
	want:= NewCoeff(0,-21,3)
	got := EmptyNewCoeff()
	got.ValidateEquation("-21y=3")
	assert.Equal(t, got.a, want.a, "Coefficants don't match.")
	assert.Equal(t, got.b, want.b, "Coefficants don't match.")
	assert.Equal(t, got.c, want.c, "Coefficants don't match.")
}

func TestExtractCoeffFromString(t *testing.T){
	got := EmptyNewCoeff()
	got_arr := got.ExtractCoeffFromString("21x+16y=3")
	want:= []float64{21,16,3}
	assert.Equal(t, len(got_arr), len(want), "lengths Coefficants don't match.")
	assert.Equal(t, got_arr[0], want[0], "Coefficants don't match.")
	assert.Equal(t, got_arr[1], want[1], "Coefficants don't match.")
	assert.Equal(t, got_arr[2], want[2], "Coefficants don't match.")
}