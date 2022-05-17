package logger

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_Debug(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(DebugLevel).
		SetOutput(buf)

	underTest.Debug("test debug")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if !strings.Contains(string(result), "test debug") {
		t.Errorf("Expected '%s' got '%s'", "test debug", string(result))
	}
}

func Test_Debugf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(DebugLevel).
		SetOutput(buf)

	underTest.Debugf("format %s", "debugf")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if !strings.Contains(string(result), "format debugf") {
		t.Errorf("Expected '%s' got '%s'", "format debugf", string(result))
	}
}

func Test_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(InfoLevel).
		SetOutput(buf)

	underTest.Debug("test debug")
	underTest.Info("test info")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if strings.Contains(string(result), "test debug") {
		t.Errorf("Unexpected '%s' in '%s'", "test debug", string(result))
	}
	if !strings.Contains(string(result), "test info") {
		t.Errorf("Expected '%s' got '%s'", "test info", string(result))
	}
}

func Test_Infof(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(InfoLevel).
		SetOutput(buf)

	underTest.Debugf("format %s", "debugf")
	underTest.Infof("format %s", "infof")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if strings.Contains(string(result), "format debugf") {
		t.Errorf("Unexpected '%s' in '%s'", "format debugf", string(result))
	}
	if !strings.Contains(string(result), "format infof") {
		t.Errorf("Expected '%s' got '%s'", "format infof", string(result))
	}
}

func Test_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(ErrorLevel).
		SetOutput(buf)

	underTest.Debug("test debug")
	underTest.Info("test info")
	underTest.Error("test error")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if strings.Contains(string(result), "test debug") {
		t.Errorf("Unexpected '%s' in '%s'", "test debug", string(result))
	}
	if strings.Contains(string(result), "test info") {
		t.Errorf("Unexpected '%s' in '%s'", "test info", string(result))
	}
	if !strings.Contains(string(result), "test error") {
		t.Errorf("Expected '%s' got '%s'", "test error", string(result))
	}
}

func Test_Errorf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(ErrorLevel).
		SetOutput(buf)

	underTest.Debugf("format %s", "debugf")
	underTest.Infof("format %s", "infof")
	underTest.Errorf("format %s", "errorf")
	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if !strings.Contains(string(result), "my-logger") {
		t.Errorf("Expected '%s' got '%s'", "my-logger", string(result))
	}
	if strings.Contains(string(result), "format debugf") {
		t.Errorf("Unexpected '%s' in '%s'", "format debugf", string(result))
	}
	if strings.Contains(string(result), "format infof") {
		t.Errorf("Unexpected '%s' in '%s'", "format infof", string(result))
	}
	if !strings.Contains(string(result), "format errorf") {
		t.Errorf("Expected '%s' got '%s'", "format errorf", string(result))
	}
}

func Test_Fatal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(FatalLevel).
		SetOutput(buf)

	underTest.Debug("test debug")
	underTest.Info("test info")
	underTest.Error("test error")

	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if len(result) > 0 {
		t.Errorf("Unexpected '%s' in result", string(result))
	}

	//testing os.Exit()
	if os.Getenv("FATAL") == "1" {
		underTest.Fatal("test fatal")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func Test_Fatalf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buf := bytes.NewBuffer([]byte{})
	underTest := NewLoggerService().
		SetServiceName("my-logger").
		SetLevel(FatalLevel).
		SetOutput(buf)

	underTest.Debug("test debug")
	underTest.Info("test info")
	underTest.Error("test error")

	// wait for async call
	time.Sleep(10 * time.Millisecond)

	result := buf.Bytes()
	if len(result) > 0 {
		t.Errorf("Unexpected '%s' in result", string(result))
	}

	//testing os.Exit()
	if os.Getenv("FATAL") == "1" {
		underTest.SetOutput(os.Stdout)
		underTest.Fatalf("format %s", "fatalf")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}
