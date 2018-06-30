package main

import (
	"github.com/golang-training/ch02/ex01/tempconv"
	"testing"
)

func TestFahrenheit(t *testing.T) {
	f := tempconv.Fahrenheit(10)
	if f.String() != "10℉" {
		t.Errorf("expected: 10℉ but was actual; %s", f)
	}
}

func TestFToC(t *testing.T) {
	f := tempconv.Fahrenheit(10)
	c := tempconv.FToC(f)
	if c.String() != "-12.222222222222221℃" {
		t.Errorf("expected: -12.222222222222221℃ but was actual; %s", c)
	}
}

func TestCelsius(t *testing.T) {
	c := tempconv.Celsius(10)
	if c.String() != "10℃" {
		t.Errorf("expected: 10℃ but was actual; %s", c)
	}
}

func TestCToF(t *testing.T) {
	c := tempconv.Celsius(10)
	f := tempconv.CToF(c)
	if f.String() != "50℉" {
		t.Errorf("expected: 50℉ but was actual; %s", f)
	}
}

func TestKelvin(t *testing.T) {
	k := tempconv.Kelvin(10)
	if k.String() != "10K" {
		t.Errorf("expected: 10K but was actual; %s", k)
	}
}

func TestKToC(t *testing.T) {
	k := tempconv.Kelvin(10)
	c := tempconv.KToC(k)
	if c.String() != "-263.15℃" {
		t.Errorf("expected: -263.15℃ but was actual; %s", c)
	}
}

func TestCToK(t *testing.T) {
	c := tempconv.Celsius(10)
	k := tempconv.CToK(c)
	if k.String() != "283.15K" {
		t.Errorf("expected: 283.15K but was actual; %s", k)
	}
}
