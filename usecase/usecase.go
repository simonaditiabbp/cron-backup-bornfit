package usecase

import (
	"fmt"
	"time"

	"github.com/simonaditiabbp/cron-backup-bornfit/model"
	"github.com/simonaditiabbp/cron-backup-bornfit/repository"
)

func NewUsecaseConnection(db_prod repository.ProdRepository, db_backup repository.BackupRepository) UsecaseFunction {
	return &UsecaseConnection{
		repo_prod:   db_prod,
		repo_backup: db_backup,
	}
}

func (uc *UsecaseConnection) InitialBackup() error {
	startTime := time.Now()
	logBuilder := ""

	header := "===========================================\n"
	header += "Starting INITIAL BACKUP process...\n"
	header += "===========================================\n"
	header += fmt.Sprintf("Started at: %s\n", startTime.Format("2006-01-02 15:04:05"))

	fmt.Print(header)
	logBuilder += header

	// Step 1: Initialize all tables (create if not exist)
	msg := "\nStep 1: Initializing database schema...\n"
	fmt.Print(msg)
	logBuilder += msg
	if err := uc.repo_backup.InitializeAllTables(); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to initialize tables: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to initialize tables: %v", err)
	}

	// Step 2: Initialize backup metadata and logs table
	msg = "\nStep 2: Initializing backup metadata and logs table...\n"
	fmt.Print(msg)
	logBuilder += msg
	if err := uc.repo_backup.InitializeBackupMetadataTable(); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to initialize backup metadata table: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to initialize backup metadata table: %v", err)
	}
	if err := uc.repo_backup.InitializeBackupLogsTable(); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to initialize backup logs table: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to initialize backup logs table: %v", err)
	}

	// Step 3: Truncate all tables in backup database
	msg = "\nStep 3: Truncating backup tables...\n"
	fmt.Print(msg)
	logBuilder += msg
	if err := uc.repo_backup.TruncateAllTables(); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to truncate tables: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to truncate tables: %v", err)
	}

	// Step 4: Backup Users
	msg = "\nStep 4: Backing up users...\n"
	fmt.Print(msg)
	logBuilder += msg
	users, err := uc.repo_prod.GetUsers()
	if err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to get users: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to get users: %v", err)
	}
	msg = fmt.Sprintf("Found %d users\n", len(users))
	fmt.Print(msg)
	logBuilder += msg
	if err := uc.repo_backup.SaveUsers(users); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to save users: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to save users: %v", err)
	}

	// Step 5: Backup Membership Plans
	msg = "Backing up membership plans...\n"
	fmt.Print(msg)
	logBuilder += msg
	plans, err := uc.repo_prod.GetMembershipPlans()
	if err != nil {
		return fmt.Errorf("failed to get membership plans: %v", err)
	}
	fmt.Printf("Found %d membership plans\n", len(plans))
	if err := uc.repo_backup.SaveMembershipPlans(plans); err != nil {
		return fmt.Errorf("failed to save membership plans: %v", err)
	}

	// Step 4: Backup Membership Plan Schedules
	fmt.Println("Backing up membership plan schedules...")
	schedules, err := uc.repo_prod.GetMembershipPlanSchedules()
	if err != nil {
		return fmt.Errorf("failed to get membership plan schedules: %v", err)
	}
	fmt.Printf("Found %d membership plan schedules\n", len(schedules))
	if err := uc.repo_backup.SaveMembershipPlanSchedules(schedules); err != nil {
		return fmt.Errorf("failed to save membership plan schedules: %v", err)
	}

	// Step 5: Backup Memberships
	fmt.Println("Backing up memberships...")
	memberships, err := uc.repo_prod.GetMemberships()
	if err != nil {
		return fmt.Errorf("failed to get memberships: %v", err)
	}
	fmt.Printf("Found %d memberships\n", len(memberships))
	if err := uc.repo_backup.SaveMemberships(memberships); err != nil {
		return fmt.Errorf("failed to save memberships: %v", err)
	}

	// Step 6: Backup Membership Checkin Logs
	fmt.Println("Backing up membership checkin logs...")
	checkinLogs, err := uc.repo_prod.GetMembershipCheckinLogs()
	if err != nil {
		return fmt.Errorf("failed to get membership checkin logs: %v", err)
	}
	fmt.Printf("Found %d membership checkin logs\n", len(checkinLogs))
	if err := uc.repo_backup.SaveMembershipCheckinLogs(checkinLogs); err != nil {
		return fmt.Errorf("failed to save membership checkin logs: %v", err)
	}

	// Step 7: Backup Checkins
	fmt.Println("Backing up checkins...")
	checkins, err := uc.repo_prod.GetCheckins()
	if err != nil {
		return fmt.Errorf("failed to get checkins: %v", err)
	}
	fmt.Printf("Found %d checkins\n", len(checkins))
	if err := uc.repo_backup.SaveCheckins(checkins); err != nil {
		return fmt.Errorf("failed to save checkins: %v", err)
	}

	// Step 8: Backup PT Session Plans
	fmt.Println("Backing up PT session plans...")
	ptPlans, err := uc.repo_prod.GetPTSessionPlans()
	if err != nil {
		return fmt.Errorf("failed to get PT session plans: %v", err)
	}
	fmt.Printf("Found %d PT session plans\n", len(ptPlans))
	if err := uc.repo_backup.SavePTSessionPlans(ptPlans); err != nil {
		return fmt.Errorf("failed to save PT session plans: %v", err)
	}

	// Step 9: Backup Personal Trainer Sessions
	fmt.Println("Backing up personal trainer sessions...")
	ptSessions, err := uc.repo_prod.GetPersonalTrainerSessions()
	if err != nil {
		return fmt.Errorf("failed to get personal trainer sessions: %v", err)
	}
	fmt.Printf("Found %d personal trainer sessions\n", len(ptSessions))
	if err := uc.repo_backup.SavePersonalTrainerSessions(ptSessions); err != nil {
		return fmt.Errorf("failed to save personal trainer sessions: %v", err)
	}

	// Step 10: Backup PT Session Bookings
	fmt.Println("Backing up PT session bookings...")
	ptBookings, err := uc.repo_prod.GetPTSessionBookings()
	if err != nil {
		return fmt.Errorf("failed to get PT session bookings: %v", err)
	}
	fmt.Printf("Found %d PT session bookings\n", len(ptBookings))
	if err := uc.repo_backup.SavePTSessionBookings(ptBookings); err != nil {
		return fmt.Errorf("failed to save PT session bookings: %v", err)
	}

	// Step 11: Backup Checkin PT Sessions
	fmt.Println("Backing up checkin PT sessions...")
	checkinPT, err := uc.repo_prod.GetCheckinPTSessions()
	if err != nil {
		return fmt.Errorf("failed to get checkin PT sessions: %v", err)
	}
	fmt.Printf("Found %d checkin PT sessions\n", len(checkinPT))
	if err := uc.repo_backup.SaveCheckinPTSessions(checkinPT); err != nil {
		return fmt.Errorf("failed to save checkin PT sessions: %v", err)
	}

	// Step 12: Backup Event Plans
	fmt.Println("Backing up event plans...")
	eventPlans, err := uc.repo_prod.GetEventPlans()
	if err != nil {
		return fmt.Errorf("failed to get event plans: %v", err)
	}
	fmt.Printf("Found %d event plans\n", len(eventPlans))
	if err := uc.repo_backup.SaveEventPlans(eventPlans); err != nil {
		return fmt.Errorf("failed to save event plans: %v", err)
	}

	// Step 13: Backup Classes
	fmt.Println("Backing up classes...")
	classes, err := uc.repo_prod.GetClasses()
	if err != nil {
		return fmt.Errorf("failed to get classes: %v", err)
	}
	fmt.Printf("Found %d classes\n", len(classes))
	if err := uc.repo_backup.SaveClasses(classes); err != nil {
		return fmt.Errorf("failed to save classes: %v", err)
	}

	// Step 14: Backup Class Purchases
	fmt.Println("Backing up class purchases...")
	classPurchases, err := uc.repo_prod.GetClassPurchases()
	if err != nil {
		return fmt.Errorf("failed to get class purchases: %v", err)
	}
	fmt.Printf("Found %d class purchases\n", len(classPurchases))
	if err := uc.repo_backup.SaveClassPurchases(classPurchases); err != nil {
		return fmt.Errorf("failed to save class purchases: %v", err)
	}

	// Step 15: Backup Class Attendances
	fmt.Println("Backing up class attendances...")
	classAttendances, err := uc.repo_prod.GetClassAttendances()
	if err != nil {
		return fmt.Errorf("failed to get class attendances: %v", err)
	}
	fmt.Printf("Found %d class attendances\n", len(classAttendances))
	if err := uc.repo_backup.SaveClassAttendances(classAttendances); err != nil {
		return fmt.Errorf("failed to save class attendances: %v", err)
	}

	// Step 16: Backup Staff Schedules
	fmt.Println("Backing up staff schedules...")
	staffSchedules, err := uc.repo_prod.GetStaffSchedules()
	if err != nil {
		return fmt.Errorf("failed to get staff schedules: %v", err)
	}
	fmt.Printf("Found %d staff schedules\n", len(staffSchedules))
	if err := uc.repo_backup.SaveStaffSchedules(staffSchedules); err != nil {
		return fmt.Errorf("failed to save staff schedules: %v", err)
	}

	// Step 17: Backup Membership Freezes
	fmt.Println("Backing up membership freezes...")
	freezes, err := uc.repo_prod.GetMembershipFreezes()
	if err != nil {
		return fmt.Errorf("failed to get membership freezes: %v", err)
	}
	fmt.Printf("Found %d membership freezes\n", len(freezes))
	if err := uc.repo_backup.SaveMembershipFreezes(freezes); err != nil {
		return fmt.Errorf("failed to save membership freezes: %v", err)
	}

	// Step 18: Backup Membership Transfers
	fmt.Println("Backing up membership transfers...")
	transfers, err := uc.repo_prod.GetMembershipTransfers()
	if err != nil {
		return fmt.Errorf("failed to get membership transfers: %v", err)
	}
	fmt.Printf("Found %d membership transfers\n", len(transfers))
	if err := uc.repo_backup.SaveMembershipTransfers(transfers); err != nil {
		return fmt.Errorf("failed to save membership transfers: %v", err)
	}

	// Step 19: Backup Transactions
	fmt.Println("Backing up transactions...")
	transactions, err := uc.repo_prod.GetTransactions()
	if err != nil {
		return fmt.Errorf("failed to get transactions: %v", err)
	}
	fmt.Printf("Found %d transactions\n", len(transactions))
	if err := uc.repo_backup.SaveTransactions(transactions); err != nil {
		return fmt.Errorf("failed to save transactions: %v", err)
	}

	// Step 20: Backup Audit Logs
	fmt.Println("Backing up audit logs...")
	auditLogs, err := uc.repo_prod.GetAuditLogs()
	if err != nil {
		return fmt.Errorf("failed to get audit logs: %v", err)
	}
	fmt.Printf("Found %d audit logs\n", len(auditLogs))
	if err := uc.repo_backup.SaveAuditLogs(auditLogs); err != nil {
		return fmt.Errorf("failed to save audit logs: %v", err)
	}

	// Save metadata for all tables
	fmt.Println("\nSaving backup metadata...")
	backupTime := time.Now()
	tables := map[string]int{
		"user":                     len(users),
		"membership_plan":          len(plans),
		"membership_plan_schedule": len(schedules),
		"membership":               len(memberships),
		"membership_checkin_log":   len(checkinLogs),
		"checkin":                  len(checkins),
		"pt_session_plan":          len(ptPlans),
		"personal_trainer_session": len(ptSessions),
		"pt_session_booking":       len(ptBookings),
		"checkin_pt_session":       len(checkinPT),
		"event_plan":               len(eventPlans),
		"class":                    len(classes),
		"class_purchase":           len(classPurchases),
		"class_attendance":         len(classAttendances),
		"staff_schedule":           len(staffSchedules),
		"membership_freeze":        len(freezes),
		"membership_transfer":      len(transfers),
		"transaction":              len(transactions),
		"audit_log":                len(auditLogs),
	}

	for tableName, count := range tables {
		metadata := model.BackupMetadata{
			Table:          tableName,
			LastBackupAt:   backupTime,
			TotalRecords:   count,
			NewRecords:     count,
			UpdatedRecords: 0,
			BackupType:     "initial",
			Status:         "success",
		}
		if err := uc.repo_backup.SaveBackupMetadata(metadata); err != nil {
			warnMsg := fmt.Sprintf("Warning: failed to save metadata for %s: %v\n", tableName, err)
			fmt.Print(warnMsg)
			logBuilder += warnMsg
		}
	}

	summary := "\n===========================================\n"
	summary += "INITIAL BACKUP completed successfully!\n"
	summary += "===========================================\n"

	// Calculate total records and save backup log
	totalRecs := len(users) + len(plans) + len(schedules) + len(memberships) +
		len(checkinLogs) + len(checkins) + len(ptPlans) + len(ptSessions) +
		len(ptBookings) + len(checkinPT) + len(eventPlans) + len(classes) +
		len(classPurchases) + len(classAttendances) + len(staffSchedules) +
		len(freezes) + len(transfers) + len(transactions) + len(auditLogs)

	endTime := time.Now()
	duration := endTime.Sub(startTime).Seconds()

	summary += fmt.Sprintf("Total Records Backed Up: %d\n", totalRecs)
	summary += fmt.Sprintf("Total Time: %.2f seconds\n", duration)
	summary += fmt.Sprintf("Completed at: %s\n", endTime.Format("2006-01-02 15:04:05"))
	summary += "===========================================\n"

	fmt.Print(summary)
	logBuilder += summary

	backupLog := model.BackupLog{
		BackupType:     "initial",
		StartTime:      startTime,
		EndTime:        endTime,
		Duration:       duration,
		Status:         "success",
		TotalNew:       totalRecs,
		TotalUpdated:   0,
		TotalRecords:   totalRecs,
		TablesAffected: 19,
		LogDetails:     &logBuilder,
	}

	if err := uc.repo_backup.SaveBackupLog(backupLog); err != nil {
		fmt.Printf("Warning: failed to save backup log: %v\n", err)
	}

	return nil
}
