# 🖥️ Abort Shutdown

**Abort Shutdown** is a lightweight Windows system tray utility written in Go that automatically detects and aborts any pending system shutdowns, restarts, or log-offs.  
It’s designed to quietly run in the background and prevent unwanted shutdowns — whether triggered manually, by software updates, or by automated processes such as **Group Policy (GPO)** shutdown commands.

This tool is especially useful in managed or corporate environments where administrators may schedule system restarts through Group Policy, but you need to temporarily keep your session active.

---

## 🔧 Features

### 🛡️ Automatic Shutdown Prevention  
Continuously monitors your system and cancels any pending shutdowns every few seconds using the Windows `shutdown /a` command — including those triggered by Windows Update or Group Policy.

### 🧰 System Tray Integration  
Runs silently in the background with an icon in the system tray for easy access and control.

### 📋 Activity Logging  
Logs all shutdown abort events to a file named `abort_shutdown.log`, including timestamps.

### ⚡ Quick Access Controls  
- **View Log:** Opens the log file in Notepad to review activity.  
- **Exit:** Gracefully quits the application.

### 🔇 Silent Operation  
Executes background commands without opening command windows or interrupting your workflow.

---

## ⚙️ How It Works

1. When launched, the app adds a small tray icon labeled **"abort_shutdown"**.  
2. Every 5 seconds, it runs the `shutdown /a` command silently to cancel any pending system shutdowns.  
3. If a shutdown is detected and aborted (including GPO-initiated ones), the event is logged in `abort_shutdown.log`.  
4. You can open the log or exit the program anytime through the tray menu.

---

## 🪟 System Requirements

- **OS:** Windows 10 or later  
- **Permissions:** Standard user permissions are typically sufficient  
- **Dependencies:** None (self-contained executable built with Go)

---

## 🧠 Technical Notes

- Implemented using Go’s [`github.com/getlantern/systray`](https://github.com/getlantern/systray) package for tray integration.  
- Uses Go’s `embed` package to bundle the tray icon directly into the binary.  
- Log handling ensures persistent tracking of abort events between sessions.  
- Helpful for environments where **Group Policy (GPO)** or other automated management tools enforce system shutdowns or restarts.
