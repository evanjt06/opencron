package main

import (
	"fmt"
	scheduler "opencron/internal"
	"time"
)

func main() {
	s := scheduler.NewScheduler()

	// ✅ 1. One-time task: Notify admin
	s.ScheduleOnce(3*time.Second, func() {
		fmt.Println("[📢] Notify admin: server rebooted")
	})

	// ✅ 2. Repeating task: Clean up expired sessions
	s.ScheduleRepeat(5*time.Second, func() {
		fmt.Println("[🧹] Cleaned up expired sessions at", time.Now().Format("15:04:05"))
	})

	// ✅ 3. Repeating task: Backup database
	s.ScheduleRepeat(10*time.Second, func() {
		fmt.Println("[💾] Performed database backup at", time.Now().Format("15:04:05"))
	})

	fmt.Println("🚀 OpenScheduler started at", time.Now().Format("15:04:05"))
	s.Start()

	select {} // block forever so scheduler keeps running
}
