package usecase

import "github.com/simonaditiabbp/cron-backup-bornfit/repository"

type UsecaseFunction interface {
	InitialBackup() error
	IncrementalBackup() error
}

type UsecaseConnection struct {
	repo_prod   repository.ProdRepository
	repo_backup repository.BackupRepository
}
