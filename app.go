package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simonaditiabbp/cron-backup-bornfit/repository"
	"github.com/simonaditiabbp/cron-backup-bornfit/usecase"
)

func appInitial(db_prod *gorm.DB, db_backup *gorm.DB) {
	start := time.Now()
	fmt.Println("at", start.Format(time.RFC3339))

	repository_prod := repository.NewProdConnection(db_prod)
	repository_backup := repository.NewBackupConnection(db_backup)

	uc := usecase.NewUsecaseConnection(repository_prod, repository_backup)

	err := uc.InitialBackup()
	if err != nil {
		fmt.Println("Error during initial backup:", err)
	}

	elapsed := time.Since(start).Seconds()

	fmt.Printf("\n\nProcess Time: %.2f seconds\n", elapsed)
}

func appIncremental(db_prod *gorm.DB, db_backup *gorm.DB) {
	start := time.Now()
	fmt.Println("at", start.Format(time.RFC3339))

	repository_prod := repository.NewProdConnection(db_prod)
	repository_backup := repository.NewBackupConnection(db_backup)

	uc := usecase.NewUsecaseConnection(repository_prod, repository_backup)

	err := uc.IncrementalBackup()
	if err != nil {
		fmt.Println("Error during incremental backup:", err)
	}

	elapsed := time.Since(start).Seconds()

	fmt.Printf("\n\nProcess Time: %.2f seconds\n", elapsed)
}
