package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

func NewProdConnection(connection *gorm.DB) ProdRepository {
	return &ProdConnection{
		db_prod: connection,
	}
}

func (conn *ProdConnection) GetUsers() ([]model.User, error) {
	var data []model.User
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipPlans() ([]model.MembershipPlan, error) {
	var data []model.MembershipPlan
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipPlanSchedules() ([]model.MembershipPlanSchedule, error) {
	var data []model.MembershipPlanSchedule
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMemberships() ([]model.Membership, error) {
	var data []model.Membership
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipCheckinLogs() ([]model.MembershipCheckinLog, error) {
	var data []model.MembershipCheckinLog
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetCheckins() ([]model.Checkin, error) {
	var data []model.Checkin
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetCheckinPTSessions() ([]model.CheckinPTSession, error) {
	var data []model.CheckinPTSession
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPTSessionPlans() ([]model.PTSessionPlan, error) {
	var data []model.PTSessionPlan
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPersonalTrainerSessions() ([]model.PersonalTrainerSession, error) {
	var data []model.PersonalTrainerSession
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPTSessionBookings() ([]model.PTSessionBooking, error) {
	var data []model.PTSessionBooking
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetEventPlans() ([]model.EventPlan, error) {
	var data []model.EventPlan
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClasses() ([]model.Class, error) {
	var data []model.Class
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClassPurchases() ([]model.ClassPurchase, error) {
	var data []model.ClassPurchase
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetClassAttendances() ([]model.ClassAttendance, error) {
	var data []model.ClassAttendance
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetStaffSchedules() ([]model.StaffSchedule, error) {
	var data []model.StaffSchedule
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipFreezes() ([]model.MembershipFreeze, error) {
	var data []model.MembershipFreeze
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetMembershipTransfers() ([]model.MembershipTransfer, error) {
	var data []model.MembershipTransfer
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetTransactions() ([]model.Transaction, error) {
	var data []model.Transaction
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetAuditLogs() ([]model.AuditLog, error) {
	var data []model.AuditLog
	err := conn.db_prod.Unscoped().Find(&data).Error
	return data, err
}

func (conn *ProdConnection) GetPrismaMigrations() ([]model.PrismaMigrations, error) {
	var data []model.PrismaMigrations
	err := conn.db_prod.Table("_prisma_migrations").Find(&data).Error
	return data, err
}
