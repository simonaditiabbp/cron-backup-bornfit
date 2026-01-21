package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/simonaditiabbp/cron-backup-bornfit/model"
)

func (uc *UsecaseConnection) IncrementalBackup() error {
	startTime := time.Now()
	logBuilder := ""

	// Maps to store changes per table
	tableChanges := make(map[string]map[string]interface{})

	// Log header
	header := "===========================================\n"
	header += "Starting INCREMENTAL BACKUP process...\n"
	header += "===========================================\n"
	header += fmt.Sprintf("Started at: %s\n", startTime.Format("2006-01-02 15:04:05"))

	fmt.Print(header)
	logBuilder += header

	// Step 1: Initialize all tables (create if not exist, add new columns if any)
	msg := "\nStep 1: Initializing database schema...\n"
	fmt.Print(msg)
	logBuilder += msg
	if err := uc.repo_backup.InitializeAllTables(); err != nil {
		errMsg := fmt.Sprintf("ERROR: failed to initialize tables: %v\n", err)
		fmt.Print(errMsg)
		logBuilder += errMsg
		return fmt.Errorf("failed to initialize tables: %v", err)
	}

	// Step 2: Initialize backup metadata and logs table (add missing columns if any)
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
	msg = "✓ All tables initialized successfully\n"
	fmt.Print(msg)
	logBuilder += msg

	totalNew := 0
	totalUpdated := 0

	// Helper function to backup table incrementally
	backupTableIncremental := func(
		tableName string,
		getModified func(time.Time) (interface{}, error),
		upsert func(interface{}) (int, int, error),
	) error {
		msg := fmt.Sprintf("\n[%s] Checking for changes...\n", tableName)
		fmt.Print(msg)
		logBuilder += msg

		// Get last backup time
		lastBackup, err := uc.repo_backup.GetLastBackupTime(tableName)
		if err != nil {
			errMsg := fmt.Sprintf("[%s] ERROR: failed to get last backup time: %v\n", tableName, err)
			fmt.Print(errMsg)
			logBuilder += errMsg
			return fmt.Errorf("failed to get last backup time for %s: %v", tableName, err)
		}

		// If no previous backup, use a very old date
		if lastBackup.IsZero() {
			lastBackup = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			msg := fmt.Sprintf("[%s] No previous backup found, fetching all data\n", tableName)
			fmt.Print(msg)
			logBuilder += msg
		} else {
			msg := fmt.Sprintf("[%s] Last backup: %s\n", tableName, lastBackup.Format("2006-01-02 15:04:05"))
			fmt.Print(msg)
			logBuilder += msg
		}

		// Get modified data
		data, err := getModified(lastBackup)
		if err != nil {
			errMsg := fmt.Sprintf("[%s] ERROR: failed to get modified data: %v\n", tableName, err)
			fmt.Print(errMsg)
			logBuilder += errMsg
			return fmt.Errorf("failed to get modified data for %s: %v", tableName, err)
		}

		// Count records
		var count int
		switch v := data.(type) {
		case []model.User:
			count = len(v)
		case []model.MembershipPlan:
			count = len(v)
		case []model.MembershipPlanSchedule:
			count = len(v)
		case []model.Membership:
			count = len(v)
		case []model.MembershipCheckinLog:
			count = len(v)
		case []model.Checkin:
			count = len(v)
		case []model.CheckinPTSession:
			count = len(v)
		case []model.PTSessionPlan:
			count = len(v)
		case []model.PersonalTrainerSession:
			count = len(v)
		case []model.PTSessionBooking:
			count = len(v)
		case []model.EventPlan:
			count = len(v)
		case []model.Class:
			count = len(v)
		case []model.ClassPurchase:
			count = len(v)
		case []model.ClassAttendance:
			count = len(v)
		case []model.StaffSchedule:
			count = len(v)
		case []model.MembershipFreeze:
			count = len(v)
		case []model.MembershipTransfer:
			count = len(v)
		case []model.Transaction:
			count = len(v)
		case []model.AuditLog:
			count = len(v)
		case []model.PrismaMigrations:
			count = len(v)
		}

		if count == 0 {
			msg := fmt.Sprintf("[%s] No changes detected\n", tableName)
			fmt.Print(msg)
			logBuilder += msg
			return nil
		}

		msg = fmt.Sprintf("[%s] Found %d modified/new records\n", tableName, count)
		fmt.Print(msg)
		logBuilder += msg

		// Get existing IDs from backup DB to separate new vs updated records
		existingMap, err := uc.repo_backup.GetExistingRecordsByTable(tableName, data)
		if err != nil {
			errMsg := fmt.Sprintf("[%s] WARNING: failed to get existing records: %v\n", tableName, err)
			fmt.Print(errMsg)
			logBuilder += errMsg
			existingMap = make(map[int]interface{}) // Continue with empty map
		}

		// Separate new and updated records
		var newRecords []interface{}
		var updatedRecords []map[string]interface{}

		// Process each record to separate new vs updated
		switch v := data.(type) {
		case []model.User:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					// Record exists - this is an update
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					// Record doesn't exist - this is new
					newRecords = append(newRecords, record)
				}
			}
		case []model.MembershipPlan:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.MembershipPlanSchedule:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.Membership:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.MembershipCheckinLog:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.Checkin:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.CheckinPTSession:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.PTSessionPlan:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.PersonalTrainerSession:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.PTSessionBooking:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.EventPlan:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.Class:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.ClassPurchase:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.ClassAttendance:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.StaffSchedule:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.MembershipFreeze:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.MembershipTransfer:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.Transaction:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.AuditLog:
			for _, record := range v {
				if beforeData, exists := existingMap[record.ID]; exists {
					updatedRecords = append(updatedRecords, map[string]interface{}{
						"before": beforeData,
						"after":  record,
					})
				} else {
					newRecords = append(newRecords, record)
				}
			}
		case []model.PrismaMigrations:
			for _, record := range v {
				newRecords = append(newRecords, record)
			}
		}

		// Upsert data and get new/update counts
		newCount, updateCount, err := upsert(data)
		if err != nil {
			errMsg := fmt.Sprintf("[%s] ERROR: failed to upsert: %v\n", tableName, err)
			fmt.Print(errMsg)
			logBuilder += errMsg

			// Save failed metadata
			metadata := model.BackupMetadata{
				Table:          tableName,
				LastBackupAt:   time.Now(),
				TotalRecords:   count,
				NewRecords:     newCount,
				UpdatedRecords: updateCount,
				BackupType:     "incremental",
				Status:         "failed",
				ErrorMessage:   stringPtr(err.Error()),
			}
			uc.repo_backup.SaveBackupMetadata(metadata)
			return fmt.Errorf("failed to upsert %s: %v", tableName, err)
		}

		// Store changes for this table
		if newCount > 0 || updateCount > 0 {
			tableChanges[tableName] = map[string]interface{}{
				"new_records":     newRecords,     // Only truly new records
				"updated_records": updatedRecords, // Updated with before/after
				"new_count":       newCount,
				"updated_count":   updateCount,
				"total_changed":   newCount + updateCount,
			}
		}

		totalNew += newCount
		totalUpdated += updateCount

		msg = fmt.Sprintf("[%s] ✓ New: %d, Updated: %d\n", tableName, newCount, updateCount)
		fmt.Print(msg)
		logBuilder += msg

		// Save successful metadata
		metadata := model.BackupMetadata{
			Table:          tableName,
			LastBackupAt:   time.Now(),
			TotalRecords:   count,
			NewRecords:     newCount,
			UpdatedRecords: updateCount,
			BackupType:     "incremental",
			Status:         "success",
		}
		if err := uc.repo_backup.SaveBackupMetadata(metadata); err != nil {
			fmt.Printf("[%s] Warning: failed to save metadata: %v\n", tableName, err)
		}

		return nil
	}

	// Backup each table
	tables := []struct {
		name        string
		getModified func(time.Time) (interface{}, error)
		upsert      func(interface{}) (int, int, error)
	}{
		{
			"user",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetUsersModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertUsers(d.([]model.User)) },
		},
		{
			"membership_plan",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipPlansModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertMembershipPlans(d.([]model.MembershipPlan))
			},
		},
		{
			"membership_plan_schedule",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipPlanSchedulesModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertMembershipPlanSchedules(d.([]model.MembershipPlanSchedule))
			},
		},
		{
			"membership",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipsModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertMemberships(d.([]model.Membership)) },
		},
		{
			"membership_checkin_log",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipCheckinLogsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertMembershipCheckinLogs(d.([]model.MembershipCheckinLog))
			},
		},
		{
			"checkin",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetCheckinsModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertCheckins(d.([]model.Checkin)) },
		},
		{
			"pt_session_plan",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetPTSessionPlansModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertPTSessionPlans(d.([]model.PTSessionPlan))
			},
		},
		{
			"personal_trainer_session",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetPersonalTrainerSessionsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertPersonalTrainerSessions(d.([]model.PersonalTrainerSession))
			},
		},
		{
			"pt_session_booking",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetPTSessionBookingsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertPTSessionBookings(d.([]model.PTSessionBooking))
			},
		},
		{
			"checkin_pt_session",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetCheckinPTSessionsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertCheckinPTSessions(d.([]model.CheckinPTSession))
			},
		},
		{
			"event_plan",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetEventPlansModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertEventPlans(d.([]model.EventPlan)) },
		},
		{
			"class",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetClassesModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertClasses(d.([]model.Class)) },
		},
		{
			"class_purchase",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetClassPurchasesModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertClassPurchases(d.([]model.ClassPurchase))
			},
		},
		{
			"class_attendance",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetClassAttendancesModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertClassAttendances(d.([]model.ClassAttendance))
			},
		},
		{
			"staff_schedule",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetStaffSchedulesModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertStaffSchedules(d.([]model.StaffSchedule))
			},
		},
		{
			"membership_freeze",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipFreezesModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertMembershipFreezes(d.([]model.MembershipFreeze))
			},
		},
		{
			"membership_transfer",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetMembershipTransfersModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertMembershipTransfers(d.([]model.MembershipTransfer))
			},
		},
		{
			"transaction",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetTransactionsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertTransactions(d.([]model.Transaction))
			},
		},
		{
			"audit_log",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetAuditLogsModifiedAfter(t) },
			func(d interface{}) (int, int, error) { return uc.repo_backup.UpsertAuditLogs(d.([]model.AuditLog)) },
		},
		{
			"_prisma_migrations",
			func(t time.Time) (interface{}, error) { return uc.repo_prod.GetPrismaMigrationsModifiedAfter(t) },
			func(d interface{}) (int, int, error) {
				return uc.repo_backup.UpsertPrismaMigrations(d.([]model.PrismaMigrations))
			},
		},
	}

	// Process each table
	tablesAffected := 0
	for _, table := range tables {
		if err := backupTableIncremental(table.name, table.getModified, table.upsert); err != nil {
			errMsg := fmt.Sprintf("Error backing up %s: %v\n", table.name, err)
			fmt.Print(errMsg)
			logBuilder += errMsg
			// Continue with other tables even if one fails
		} else {
			tablesAffected++
		}
	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Seconds()

	summary := "\n===========================================\n"
	summary += "INCREMENTAL BACKUP SUMMARY\n"
	summary += "===========================================\n"
	summary += fmt.Sprintf("Total New Records: %d\n", totalNew)
	summary += fmt.Sprintf("Total Updated Records: %d\n", totalUpdated)
	summary += fmt.Sprintf("Total Time: %.2f seconds\n", elapsed)
	summary += fmt.Sprintf("Completed at: %s\n", endTime.Format("2006-01-02 15:04:05"))
	summary += "===========================================\n"

	fmt.Print(summary)
	logBuilder += summary

	// Create backup log structure
	backupLog := model.BackupLog{
		BackupType:     "incremental",
		StartTime:      startTime,
		EndTime:        endTime,
		Duration:       elapsed,
		Status:         "success",
		TotalNew:       totalNew,
		TotalUpdated:   totalUpdated,
		TotalRecords:   totalNew + totalUpdated,
		TablesAffected: tablesAffected,
		LogDetails:     &logBuilder,
	}

	// Map table changes to appropriate JSONB columns
	for tableName, changes := range tableChanges {
		jsonBytes, err := json.MarshalIndent(changes, "", "  ")
		if err != nil {
			errMsg := fmt.Sprintf("ERROR: failed to marshal changes for %s: %v\n", tableName, err)
			fmt.Print(errMsg)
			continue
		}
		changesJSON := string(jsonBytes)

		// Assign to appropriate column based on table name
		switch tableName {
		case "user":
			backupLog.UserChanges = &changesJSON
		case "membership_plan":
			backupLog.MembershipPlanChanges = &changesJSON
		case "membership_plan_schedule":
			backupLog.MembershipPlanScheduleChanges = &changesJSON
		case "membership":
			backupLog.MembershipChanges = &changesJSON
		case "membership_checkin_log":
			backupLog.MembershipCheckinLogChanges = &changesJSON
		case "checkin":
			backupLog.CheckinChanges = &changesJSON
		case "checkin_pt_session":
			backupLog.CheckinPTSessionChanges = &changesJSON
		case "pt_session_plan":
			backupLog.PTSessionPlanChanges = &changesJSON
		case "personal_trainer_session":
			backupLog.PersonalTrainerSessionChanges = &changesJSON
		case "pt_session_booking":
			backupLog.PTSessionBookingChanges = &changesJSON
		case "event_plan":
			backupLog.EventPlanChanges = &changesJSON
		case "class":
			backupLog.ClassChanges = &changesJSON
		case "class_purchase":
			backupLog.ClassPurchaseChanges = &changesJSON
		case "class_attendance":
			backupLog.ClassAttendanceChanges = &changesJSON
		case "staff_schedule":
			backupLog.StaffScheduleChanges = &changesJSON
		case "membership_freeze":
			backupLog.MembershipFreezeChanges = &changesJSON
		case "membership_transfer":
			backupLog.MembershipTransferChanges = &changesJSON
		case "transaction":
			backupLog.TransactionChanges = &changesJSON
		case "audit_log":
			backupLog.AuditLogChanges = &changesJSON
		case "_prisma_migrations":
			backupLog.PrismaMigrationsChanges = &changesJSON
		}
	}

	// Save backup log
	if err := uc.repo_backup.SaveBackupLog(backupLog); err != nil {
		fmt.Printf("Warning: failed to save backup log: %v\n", err)
	}

	return nil
}

func stringPtr(s string) *string {
	return &s
}
