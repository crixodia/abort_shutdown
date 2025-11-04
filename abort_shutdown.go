package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/getlantern/systray"
)

//go:embed abort_shutdown.ico
var iconData []byte

var logger *log.Logger
var logFile *os.File

func main() {
	runtime.LockOSThread() // required for GUI apps
	systray.Run(onReady, onExit)
}

func onReady() {
	// Tray icon & tooltip
	systray.SetIcon(iconData)
	systray.SetTitle("Shutdown Monitor")
	systray.SetTooltip("Monitoring and aborting scheduled shutdowns...")

	// Menu options
	mViewLog := systray.AddMenuItem("View Log", "Open the log file")
	mQuit := systray.AddMenuItem("Exit", "Quit the program")

	// Start the background monitor
	go startMonitor()

	// Menu actions
	go func() {
		for {
			select {
			case <-mViewLog.ClickedCh:
				openLogFile()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func startMonitor() {
	var err error
	logFile, err = os.OpenFile("abort_shutdown.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger = log.New(logFile, "", log.LstdFlags)
	logger.Println("Shutdown Monitor started — will abort any pending shutdowns.")

	for {
		cmd := exec.Command("shutdown", "/a")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // silent execution
		err := cmd.Run()
		if err == nil {
			logger.Println("Detected and aborted a pending shutdown.")
		}
		time.Sleep(5 * time.Second)
	}
}

func openLogFile() {
	cmd := exec.Command("notepad.exe", "abort_shutdown.log")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}

func onExit() {
	if logFile != nil {
		logFile.Close()
	}
}
