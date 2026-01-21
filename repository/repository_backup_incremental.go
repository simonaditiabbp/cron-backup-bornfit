package repository

import (
	"fmt"
	"time"

	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

// InitializeBackupMetadataTable creates backup_metadata table if not exists
func (conn *BackupConnection) InitializeBackupMetadataTable() error {
	return conn.db_backup.AutoMigrate(&model.BackupMetadata{}).Error
}

// InitializeBackupLogsTable creates backup_logs table if not exists
func (conn *BackupConnection) InitializeBackupLogsTable() error {
	return conn.db_backup.AutoMigrate(&model.BackupLog{}).Error
}

// SaveBackupLog saves backup log entry
func (conn *BackupConnection) SaveBackupLog(log model.BackupLog) error {
	log.CreatedAt = time.Now()
	return conn.db_backup.Create(&log).Error
}

// GetLastBackupTime gets the last successful backup timestamp for a table
func (conn *BackupConnection) GetLastBackupTime(tableName string) (time.Time, error) {
	var metadata model.BackupMetadata
	err := conn.db_backup.Where("\"table_name\" = ? AND status = ?", tableName, "success").
		Order("last_backup_at DESC").
		First(&metadata).Error

	if err != nil {
		// If no record found, return zero time (will backup all data)
		return time.Time{}, nil
	}

	return metadata.LastBackupAt, nil
}

// SaveBackupMetadata saves backup metadata (creates new record for history tracking)
func (conn *BackupConnection) SaveBackupMetadata(metadata model.BackupMetadata) error {
	metadata.CreatedAt = time.Now()
	metadata.UpdatedAt = time.Now()

	// Delete old unique constraint if exists and create new record
	// This allows multiple backup history records per table
	err := conn.db_backup.Exec("ALTER TABLE backup_metadata DROP CONSTRAINT IF EXISTS backup_metadata_table_name_key").Error
	if err != nil {
		fmt.Printf("Note: Could not drop unique constraint (may not exist): %v\n", err)
	}

	return conn.db_backup.Create(&metadata).Error
}

// Helper function to upsert records
func upsertRecords[T any](conn *BackupConnection, data []T, tableName string) (newCount, updateCount int, err error) {
	if len(data) == 0 {
		return 0, 0, nil
	}

	for _, record := range data {
		// Try to find existing record by ID using reflection
		var existing T
		result := conn.db_backup.Table(tableName).Where("id = ?", getID(record)).First(&existing)

		if result.Error != nil {
			// Record doesn't exist, insert it
			if err := conn.db_backup.Table(tableName).Create(&record).Error; err != nil {
				return newCount, updateCount, fmt.Errorf("failed to insert record: %v", err)
			}
			newCount++
		} else {
			// Record exists, update it
			if err := conn.db_backup.Table(tableName).Where("id = ?", getID(record)).Updates(&record).Error; err != nil {
				return newCount, updateCount, fmt.Errorf("failed to update record: %v", err)
			}
			updateCount++
		}
	}

	return newCount, updateCount, nil
}

// Helper to get ID from struct using type assertion
func getID(record interface{}) interface{} {
	switch v := record.(type) {
	case model.User:
		return v.ID
	case model.MembershipPlan:
		return v.ID
	case model.MembershipPlanSchedule:
		return v.ID
	case model.Membership:
		return v.ID
	case model.MembershipCheckinLog:
		return v.ID
	case model.Checkin:
		return v.ID
	case model.CheckinPTSession:
		return v.ID
	case model.PTSessionPlan:
		return v.ID
	case model.PersonalTrainerSession:
		return v.ID
	case model.PTSessionBooking:
		return v.ID
	case model.EventPlan:
		return v.ID
	case model.Class:
		return v.ID
	case model.ClassPurchase:
		return v.ID
	case model.ClassAttendance:
		return v.ID
	case model.StaffSchedule:
		return v.ID
	case model.MembershipFreeze:
		return v.ID
	case model.MembershipTransfer:
		return v.ID
	case model.Transaction:
		return v.ID
	case model.AuditLog:
		return v.ID
	case model.PrismaMigrations:
		return v.ID
	default:
		return 0
	}
}

// Upsert methods for each table
func (conn *BackupConnection) UpsertUsers(data []model.User) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "user")
}

func (conn *BackupConnection) UpsertMembershipPlans(data []model.MembershipPlan) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership_plan")
}

func (conn *BackupConnection) UpsertMembershipPlanSchedules(data []model.MembershipPlanSchedule) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership_plan_schedule")
}

func (conn *BackupConnection) UpsertMemberships(data []model.Membership) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership")
}

func (conn *BackupConnection) UpsertMembershipCheckinLogs(data []model.MembershipCheckinLog) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership_checkin_log")
}

func (conn *BackupConnection) UpsertCheckins(data []model.Checkin) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "checkin")
}

func (conn *BackupConnection) UpsertCheckinPTSessions(data []model.CheckinPTSession) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "checkin_pt_session")
}

func (conn *BackupConnection) UpsertPTSessionPlans(data []model.PTSessionPlan) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "pt_session_plan")
}

func (conn *BackupConnection) UpsertPersonalTrainerSessions(data []model.PersonalTrainerSession) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "personal_trainer_session")
}

func (conn *BackupConnection) UpsertPTSessionBookings(data []model.PTSessionBooking) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "pt_session_booking")
}

func (conn *BackupConnection) UpsertEventPlans(data []model.EventPlan) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "event_plan")
}

func (conn *BackupConnection) UpsertClasses(data []model.Class) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "class")
}

func (conn *BackupConnection) UpsertClassPurchases(data []model.ClassPurchase) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "class_purchase")
}

func (conn *BackupConnection) UpsertClassAttendances(data []model.ClassAttendance) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "class_attendance")
}

func (conn *BackupConnection) UpsertStaffSchedules(data []model.StaffSchedule) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "staff_schedule")
}

func (conn *BackupConnection) UpsertMembershipFreezes(data []model.MembershipFreeze) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership_freeze")
}

func (conn *BackupConnection) UpsertMembershipTransfers(data []model.MembershipTransfer) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "membership_transfer")
}

func (conn *BackupConnection) UpsertTransactions(data []model.Transaction) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "transaction")
}

func (conn *BackupConnection) UpsertAuditLogs(data []model.AuditLog) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "audit_log")
}

func (conn *BackupConnection) UpsertPrismaMigrations(data []model.PrismaMigrations) (newCount, updateCount int, err error) {
	return upsertRecords(conn, data, "_prisma_migrations")
}

// GetExistingRecordsByTable retrieves existing records from backup DB by their IDs
// Returns a map of ID -> record for records that exist in backup
func (conn *BackupConnection) GetExistingRecordsByTable(tableName string, data interface{}) (map[int]interface{}, error) {
	existingMap := make(map[int]interface{})

	// Extract IDs from the data
	var ids []int
	switch v := data.(type) {
	case []model.User:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.MembershipPlan:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.MembershipPlanSchedule:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.Membership:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.MembershipCheckinLog:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.Checkin:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.CheckinPTSession:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.PTSessionPlan:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.PersonalTrainerSession:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.PTSessionBooking:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.EventPlan:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.Class:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.ClassPurchase:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.ClassAttendance:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.StaffSchedule:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.MembershipFreeze:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.MembershipTransfer:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.Transaction:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.AuditLog:
		for _, r := range v {
			ids = append(ids, r.ID)
		}
	case []model.PrismaMigrations:
		// PrismaMigrations uses string UUID as ID, cannot append to []int slice
		// Skip extraction and return empty map - upsertRecords will handle existence check per record
		return existingMap, nil
	default:
		return existingMap, nil
	}

	if len(ids) == 0 {
		return existingMap, nil
	}

	// Query backup DB for existing records with these IDs
	switch data.(type) {
	case []model.User:
		var existing []model.User
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.MembershipPlan:
		var existing []model.MembershipPlan
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.MembershipPlanSchedule:
		var existing []model.MembershipPlanSchedule
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.Membership:
		var existing []model.Membership
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.MembershipCheckinLog:
		var existing []model.MembershipCheckinLog
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.Checkin:
		var existing []model.Checkin
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.CheckinPTSession:
		var existing []model.CheckinPTSession
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.PTSessionPlan:
		var existing []model.PTSessionPlan
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.PersonalTrainerSession:
		var existing []model.PersonalTrainerSession
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.PTSessionBooking:
		var existing []model.PTSessionBooking
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.EventPlan:
		var existing []model.EventPlan
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.Class:
		var existing []model.Class
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.ClassPurchase:
		var existing []model.ClassPurchase
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.ClassAttendance:
		var existing []model.ClassAttendance
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.StaffSchedule:
		var existing []model.StaffSchedule
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.MembershipFreeze:
		var existing []model.MembershipFreeze
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.MembershipTransfer:
		var existing []model.MembershipTransfer
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.Transaction:
		var existing []model.Transaction
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	case []model.AuditLog:
		var existing []model.AuditLog
		if err := conn.db_backup.Unscoped().Where("id IN (?)", ids).Find(&existing).Error; err != nil {
			return existingMap, err
		}
		for _, record := range existing {
			existingMap[record.ID] = record
		}
	}

	return existingMap, nil
}
