package repository

import (
	"time"

	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

// Get data modified after specific time
func (conn *ProdConnection) GetUsersModifiedAfter(lastBackup time.Time) ([]model.User, error) {
	var data []model.User
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipPlansModifiedAfter(lastBackup time.Time) ([]model.MembershipPlan, error) {
	var data []model.MembershipPlan
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipPlanSchedulesModifiedAfter(lastBackup time.Time) ([]model.MembershipPlanSchedule, error) {
	var data []model.MembershipPlanSchedule
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipsModifiedAfter(lastBackup time.Time) ([]model.Membership, error) {
	var data []model.Membership
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipCheckinLogsModifiedAfter(lastBackup time.Time) ([]model.MembershipCheckinLog, error) {
	var data []model.MembershipCheckinLog
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetCheckinsModifiedAfter(lastBackup time.Time) ([]model.Checkin, error) {
	var data []model.Checkin
	err := conn.db_prod.Where("created_at > ?", lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetCheckinPTSessionsModifiedAfter(lastBackup time.Time) ([]model.CheckinPTSession, error) {
	var data []model.CheckinPTSession
	err := conn.db_prod.Where("created_at > ?", lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPTSessionPlansModifiedAfter(lastBackup time.Time) ([]model.PTSessionPlan, error) {
	var data []model.PTSessionPlan
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPersonalTrainerSessionsModifiedAfter(lastBackup time.Time) ([]model.PersonalTrainerSession, error) {
	var data []model.PersonalTrainerSession
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPTSessionBookingsModifiedAfter(lastBackup time.Time) ([]model.PTSessionBooking, error) {
	var data []model.PTSessionBooking
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetEventPlansModifiedAfter(lastBackup time.Time) ([]model.EventPlan, error) {
	var data []model.EventPlan
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClassesModifiedAfter(lastBackup time.Time) ([]model.Class, error) {
	var data []model.Class
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClassPurchasesModifiedAfter(lastBackup time.Time) ([]model.ClassPurchase, error) {
	var data []model.ClassPurchase
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClassAttendancesModifiedAfter(lastBackup time.Time) ([]model.ClassAttendance, error) {
	var data []model.ClassAttendance
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetStaffSchedulesModifiedAfter(lastBackup time.Time) ([]model.StaffSchedule, error) {
	var data []model.StaffSchedule
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipFreezesModifiedAfter(lastBackup time.Time) ([]model.MembershipFreeze, error) {
	var data []model.MembershipFreeze
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipTransfersModifiedAfter(lastBackup time.Time) ([]model.MembershipTransfer, error) {
	var data []model.MembershipTransfer
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetTransactionsModifiedAfter(lastBackup time.Time) ([]model.Transaction, error) {
	var data []model.Transaction
	err := conn.db_prod.Where("updated_at > ? OR created_at > ?", lastBackup, lastBackup).Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetAuditLogsModifiedAfter(lastBackup time.Time) ([]model.AuditLog, error) {
	var data []model.AuditLog
	err := conn.db_prod.Where("timestamp > ?", lastBackup).Find(&data).Error
	return data, err
}
