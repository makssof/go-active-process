package ActiveProcess

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/mitchellh/go-ps"
)

var (
	user32, errLoadDll                     = syscall.LoadDLL("user32.dll")
	GetForegroundWindow, errLoadProc1      = user32.FindProc("GetForegroundWindow")
	GetWindowThreadProcessId, errLoadProc2 = user32.FindProc("GetWindowThreadProcessId")
)

// Get returns current foreground window executable name
func Get() (string, error) {
	if errLoadDll != nil || errLoadProc1 != nil || errLoadProc2 != nil {
		return "", fmt.Errorf("can't get active window")
	}

	winHandle, _, err := syscall.SyscallN(GetForegroundWindow.Addr())

	if err != 0 {
		return "nil", fmt.Errorf("can't get active window: %v", err)
	}

	var procId int
	_, _, err = syscall.SyscallN(GetWindowThreadProcessId.Addr(), winHandle, uintptr(unsafe.Pointer(&procId)))

	if err != 0 {
		return "nil", fmt.Errorf("can't get active window id: %v", err)
	}

	proc, procErr := ps.FindProcess(procId)

	if procErr != nil {
		return "nil", fmt.Errorf("can't get active window process: %v", err)
	}

	if proc != nil {
		return proc.Executable(), nil
	}

	return "", fmt.Errorf("can't find active window process")
}
