package goerrorhandler

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

var (
	want     = errors.New("This is error message")
	errPlace = "Test Palce"
)

func TestCheckErr(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErr(want)
	t.Log(buf.String())
	t.Log(want.Error())
	if strings.Contains(want.Error(), buf.String()) {
		t.Errorf("Function checkErr() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErr() contain output %q", want)
	}
}

func TestCheckErrPlace(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErrPlace(want, errPlace)
	t.Log(buf.String())
	t.Log(want.Error())
	if strings.Contains(want.Error(), buf.String()) && strings.Contains(errPlace, buf.String()) {
		t.Errorf("Function checkErrPlace() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErrPlace() contain output %q, %q", want, errPlace)
	}
}

func TestCheckErrDebugTrue(t *testing.T) {
	var buf bytes.Buffer
	var debugmode bool = true
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErrDebug(want, debugmode)
	t.Log(buf.String())
	t.Log(want.Error())
	if strings.Contains(want.Error(), buf.String()) {
		t.Errorf("Function checkErrDebug() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErrDebug() contain output %q", want)
	}
}

func TestCheckErrDebugFalse(t *testing.T) {
	var buf bytes.Buffer
	var debugmode bool = false
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErrDebug(want, debugmode)
	t.Log(buf.String())
	t.Log(want.Error())
	if !(strings.Contains(want.Error(), buf.String())) {
		t.Errorf("Function checkErrDebug() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErrDebug() no output %q", want)
	}
}

func TestCheckErrPlaceDebugTrue(t *testing.T) {
	var buf bytes.Buffer
	var debugmode bool = true
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErrPlaceDebug(want, errPlace, debugmode)
	t.Log(buf.String())
	t.Log(want.Error())
	if strings.Contains(want.Error(), buf.String()) || strings.Contains(errPlace, buf.String()) {
		t.Errorf("Function checkErrPlaceDebug() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErrPlaceDebug() contain output %q and/or %q", want, errPlace)
	}
}

func TestCheckErrPlaceDebugFalse(t *testing.T) {
	var buf bytes.Buffer
	var debugmode bool = false
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	checkErrPlaceDebug(want, errPlace, debugmode)
	t.Log(buf.String())
	t.Log(want.Error())
	if !(strings.Contains(want.Error(), buf.String()) || strings.Contains(errPlace, buf.String())) {
		t.Errorf("Function checkErrPlaceDebug() = %q, want %q", buf, want)
	} else {
		t.Logf("Function checkErrPlaceDebug() no output %q and/or %q", want, errPlace)
	}
}
