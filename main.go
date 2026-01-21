package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"github.com/simonaditiabbp/cron-backup-bornfit/database"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	fmt.Println("===========================================")
	fmt.Println("Cron Backup Bornfit - Starting...")
	fmt.Println("===========================================")

	database.SetupProd()
	database.SetupBackup()

	// Get backup mode from environment (initial or incremental)
	backupMode := os.Getenv("BACKUP_MODE")
	if backupMode == "" {
		backupMode = "incremental" // Default to incremental
	}

	// Get cron schedule from environment
	cronSchedule := os.Getenv("CRON_SCHEDULE")
	if cronSchedule == "" {
		cronSchedule = "0 0 0,12 * * *" // Default: every day at 00:00 and 12:00
	}

	fmt.Printf("Backup Mode: %s\n", backupMode)
	fmt.Printf("Cron Schedule: %s\n", cronSchedule)
	fmt.Println("===========================================")

	// Option to run backup immediately on startup
	runOnStartup := os.Getenv("RUN_ON_STARTUP")
	if runOnStartup == "true" {
		fmt.Println("\nRunning backup immediately on startup...")
		if backupMode == "initial" {
			appInitial(database.GetDBProd(), database.GetDBBackup())
		} else {
			appIncremental(database.GetDBProd(), database.GetDBBackup())
		}
	}

	// Setup cron scheduler
	c := cron.New()

	if backupMode == "initial" {
		fmt.Println("\n\nStarting INITIAL backup job...")
		appInitial(database.GetDBProd(), database.GetDBBackup())
	} else {
		c.AddFunc(cronSchedule, func() {
			fmt.Println("\n\nCron triggered - Starting INCREMENTAL backup job...")
			appIncremental(database.GetDBProd(), database.GetDBBackup())
		})
	}

	go c.Start()

	fmt.Println("\nCron scheduler started. Waiting for scheduled backup jobs...")
	fmt.Println("Press Ctrl+C to stop the service")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("\n\nShutting down gracefully...")
}
