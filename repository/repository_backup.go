package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

func NewBackupConnection(connection *gorm.DB) BackupRepository {
	return &BackupConnection{
		db_backup: connection,
	}
}

// InitializeAllTables creates all tables if they don't exist using GORM AutoMigrate
func (conn *BackupConnection) InitializeAllTables() error {
	fmt.Println("Creating tables if not exist...")

	// AutoMigrate will create tables if they don't exist
	err := conn.db_backup.AutoMigrate(
		&model.User{},
		&model.MembershipPlan{},
		&model.MembershipPlanSchedule{},
		&model.Membership{},
		&model.MembershipCheckinLog{},
		&model.Checkin{},
		&model.CheckinPTSession{},
		&model.PTSessionPlan{},
		&model.PersonalTrainerSession{},
		&model.PTSessionBooking{},
		&model.EventPlan{},
		&model.Class{},
		&model.ClassPurchase{},
		&model.ClassAttendance{},
		&model.StaffSchedule{},
		&model.MembershipFreeze{},
		&model.MembershipTransfer{},
		&model.Transaction{},
		&model.AuditLog{},
		&model.BackupMetadata{},
		&model.BackupLog{},
		&model.PrismaMigrations{},
	).Error

	if err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	fmt.Println("✓ All tables initialized successfully")
	return nil
}

func (conn *BackupConnection) TruncateAllTables() error {
	// Truncate dalam urutan yang benar untuk menghindari foreign key constraint
	tables := []string{
		"audit_log",
		"transaction",
		"membership_transfer",
		"membership_freeze",
		"staff_schedule",
		"class_attendance",
		"class_purchase",
		"class",
		"event_plan",
		"pt_session_booking",
		"personal_trainer_session",
		"pt_session_plan",
		"checkin_pt_session",
		"checkin",
		"membership_checkin_log",
		"membership",
		"membership_plan_schedule",
		"membership_plan",
		"user",
		"_prisma_migrations",
	}

	for _, table := range tables {
		// Use double quotes to escape table names (especially for reserved keywords like "user")
		err := conn.db_backup.Exec(fmt.Sprintf("TRUNCATE TABLE \"%s\" RESTART IDENTITY CASCADE", table)).Error
		if err != nil {
			return fmt.Errorf("failed to truncate %s: %v", table, err)
		}
	}

	return nil
}

func (conn *BackupConnection) SaveUsers(data []model.User) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		// Insert records one by one to avoid reflection issues with pointer fields
		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save user: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMembershipPlans(data []model.MembershipPlan) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership_plan: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMembershipPlanSchedules(data []model.MembershipPlanSchedule) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership_plan_schedule: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMemberships(data []model.Membership) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMembershipCheckinLogs(data []model.MembershipCheckinLog) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership_checkin_log: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveCheckins(data []model.Checkin) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save checkin: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveCheckinPTSessions(data []model.CheckinPTSession) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save checkin_pt_session: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SavePTSessionPlans(data []model.PTSessionPlan) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save pt_session_plan: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SavePersonalTrainerSessions(data []model.PersonalTrainerSession) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save personal_trainer_session: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SavePTSessionBookings(data []model.PTSessionBooking) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save pt_session_booking: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveEventPlans(data []model.EventPlan) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save event_plan: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveClasses(data []model.Class) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save class: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveClassPurchases(data []model.ClassPurchase) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save class_purchase: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveClassAttendances(data []model.ClassAttendance) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save class_attendance: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveStaffSchedules(data []model.StaffSchedule) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save staff_schedule: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMembershipFreezes(data []model.MembershipFreeze) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership_freeze: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveMembershipTransfers(data []model.MembershipTransfer) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save membership_transfer: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveTransactions(data []model.Transaction) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save transaction: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SaveAuditLogs(data []model.AuditLog) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save audit_log: %v", err)
			}
		}
	}
	return nil
}

func (conn *BackupConnection) SavePrismaMigrations(data []model.PrismaMigrations) error {
	if len(data) == 0 {
		return nil
	}

	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]

		for _, record := range batch {
			if err := conn.db_backup.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to save _prisma_migrations: %v", err)
			}
		}
	}
	return nil
}
