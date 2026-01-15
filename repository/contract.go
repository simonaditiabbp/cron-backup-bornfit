package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

type ProdRepository interface {
	// Get all data
	GetUsers() ([]model.User, error)
	GetMembershipPlans() ([]model.MembershipPlan, error)
	GetMembershipPlanSchedules() ([]model.MembershipPlanSchedule, error)
	GetMemberships() ([]model.Membership, error)
	GetMembershipCheckinLogs() ([]model.MembershipCheckinLog, error)
	GetCheckins() ([]model.Checkin, error)
	GetCheckinPTSessions() ([]model.CheckinPTSession, error)
	GetPTSessionPlans() ([]model.PTSessionPlan, error)
	GetPersonalTrainerSessions() ([]model.PersonalTrainerSession, error)
	GetPTSessionBookings() ([]model.PTSessionBooking, error)
	GetEventPlans() ([]model.EventPlan, error)
	GetClasses() ([]model.Class, error)
	GetClassPurchases() ([]model.ClassPurchase, error)
	GetClassAttendances() ([]model.ClassAttendance, error)
	GetStaffSchedules() ([]model.StaffSchedule, error)
	GetMembershipFreezes() ([]model.MembershipFreeze, error)
	GetMembershipTransfers() ([]model.MembershipTransfer, error)
	GetTransactions() ([]model.Transaction, error)
	GetAuditLogs() ([]model.AuditLog, error)

	// Get data modified after specific time (for incremental backup)
	GetUsersModifiedAfter(lastBackup time.Time) ([]model.User, error)
	GetMembershipPlansModifiedAfter(lastBackup time.Time) ([]model.MembershipPlan, error)
	GetMembershipPlanSchedulesModifiedAfter(lastBackup time.Time) ([]model.MembershipPlanSchedule, error)
	GetMembershipsModifiedAfter(lastBackup time.Time) ([]model.Membership, error)
	GetMembershipCheckinLogsModifiedAfter(lastBackup time.Time) ([]model.MembershipCheckinLog, error)
	GetCheckinsModifiedAfter(lastBackup time.Time) ([]model.Checkin, error)
	GetCheckinPTSessionsModifiedAfter(lastBackup time.Time) ([]model.CheckinPTSession, error)
	GetPTSessionPlansModifiedAfter(lastBackup time.Time) ([]model.PTSessionPlan, error)
	GetPersonalTrainerSessionsModifiedAfter(lastBackup time.Time) ([]model.PersonalTrainerSession, error)
	GetPTSessionBookingsModifiedAfter(lastBackup time.Time) ([]model.PTSessionBooking, error)
	GetEventPlansModifiedAfter(lastBackup time.Time) ([]model.EventPlan, error)
	GetClassesModifiedAfter(lastBackup time.Time) ([]model.Class, error)
	GetClassPurchasesModifiedAfter(lastBackup time.Time) ([]model.ClassPurchase, error)
	GetClassAttendancesModifiedAfter(lastBackup time.Time) ([]model.ClassAttendance, error)
	GetStaffSchedulesModifiedAfter(lastBackup time.Time) ([]model.StaffSchedule, error)
	GetMembershipFreezesModifiedAfter(lastBackup time.Time) ([]model.MembershipFreeze, error)
	GetMembershipTransfersModifiedAfter(lastBackup time.Time) ([]model.MembershipTransfer, error)
	GetTransactionsModifiedAfter(lastBackup time.Time) ([]model.Transaction, error)
	GetAuditLogsModifiedAfter(lastBackup time.Time) ([]model.AuditLog, error)
}

type ProdConnection struct {
	db_prod *gorm.DB
}

type BackupRepository interface {
	// Initial backup methods
	InitializeAllTables() error
	TruncateAllTables() error
	SaveUsers(data []model.User) error
	SaveMembershipPlans(data []model.MembershipPlan) error
	SaveMembershipPlanSchedules(data []model.MembershipPlanSchedule) error
	SaveMemberships(data []model.Membership) error
	SaveMembershipCheckinLogs(data []model.MembershipCheckinLog) error
	SaveCheckins(data []model.Checkin) error
	SaveCheckinPTSessions(data []model.CheckinPTSession) error
	SavePTSessionPlans(data []model.PTSessionPlan) error
	SavePersonalTrainerSessions(data []model.PersonalTrainerSession) error
	SavePTSessionBookings(data []model.PTSessionBooking) error
	SaveEventPlans(data []model.EventPlan) error
	SaveClasses(data []model.Class) error
	SaveClassPurchases(data []model.ClassPurchase) error
	SaveClassAttendances(data []model.ClassAttendance) error
	SaveStaffSchedules(data []model.StaffSchedule) error
	SaveMembershipFreezes(data []model.MembershipFreeze) error
	SaveMembershipTransfers(data []model.MembershipTransfer) error
	SaveTransactions(data []model.Transaction) error
	SaveAuditLogs(data []model.AuditLog) error

	// Incremental backup methods (upsert)
	UpsertUsers(data []model.User) (newCount, updateCount int, err error)
	UpsertMembershipPlans(data []model.MembershipPlan) (newCount, updateCount int, err error)
	UpsertMembershipPlanSchedules(data []model.MembershipPlanSchedule) (newCount, updateCount int, err error)
	UpsertMemberships(data []model.Membership) (newCount, updateCount int, err error)
	UpsertMembershipCheckinLogs(data []model.MembershipCheckinLog) (newCount, updateCount int, err error)
	UpsertCheckins(data []model.Checkin) (newCount, updateCount int, err error)
	UpsertCheckinPTSessions(data []model.CheckinPTSession) (newCount, updateCount int, err error)
	UpsertPTSessionPlans(data []model.PTSessionPlan) (newCount, updateCount int, err error)
	UpsertPersonalTrainerSessions(data []model.PersonalTrainerSession) (newCount, updateCount int, err error)
	UpsertPTSessionBookings(data []model.PTSessionBooking) (newCount, updateCount int, err error)
	UpsertEventPlans(data []model.EventPlan) (newCount, updateCount int, err error)
	UpsertClasses(data []model.Class) (newCount, updateCount int, err error)
	UpsertClassPurchases(data []model.ClassPurchase) (newCount, updateCount int, err error)
	UpsertClassAttendances(data []model.ClassAttendance) (newCount, updateCount int, err error)
	UpsertStaffSchedules(data []model.StaffSchedule) (newCount, updateCount int, err error)
	UpsertMembershipFreezes(data []model.MembershipFreeze) (newCount, updateCount int, err error)
	UpsertMembershipTransfers(data []model.MembershipTransfer) (newCount, updateCount int, err error)
	UpsertTransactions(data []model.Transaction) (newCount, updateCount int, err error)
	UpsertAuditLogs(data []model.AuditLog) (newCount, updateCount int, err error)

	// Backup metadata tracking
	GetLastBackupTime(tableName string) (time.Time, error)
	SaveBackupMetadata(metadata model.BackupMetadata) error
	InitializeBackupMetadataTable() error

	// Backup logs
	SaveBackupLog(log model.BackupLog) error
	InitializeBackupLogsTable() error

	// Get existing records for change tracking
	GetExistingRecordsByTable(tableName string, data interface{}) (map[int]interface{}, error)
}

type BackupConnection struct {
	db_backup *gorm.DB
}
