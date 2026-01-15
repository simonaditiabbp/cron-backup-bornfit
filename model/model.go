package model

import (
	"time"
)

type User struct {
	ID                    int        `gorm:"column:id;primary_key" json:"id"`
	Name                  string     `gorm:"column:name" json:"name"`
	Email                 *string    `gorm:"column:email" json:"email"`
	Password              *string    `gorm:"column:password" json:"password"`
	Phone                 *string    `gorm:"column:phone" json:"phone"`
	Role                  string     `gorm:"column:role" json:"role"`
	DateOfBirth           *time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	NikPassport           *string    `gorm:"column:nik_passport" json:"nik_passport"`
	EmergencyContactName  *string    `gorm:"column:emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactPhone *string    `gorm:"column:emergency_contact_phone" json:"emergency_contact_phone"`
	Photo                 *string    `gorm:"column:photo" json:"photo"`
	QrCode                string     `gorm:"column:qr_code" json:"qr_code"`
	Latitude              *float64   `gorm:"column:latitude" json:"latitude"`
	Longitude             *float64   `gorm:"column:longitude" json:"longitude"`
	IsDeleted             bool       `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy             *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy             *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt             time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy             *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (User) TableName() string {
	return "user"
}

type MembershipPlan struct {
	ID                    int        `gorm:"column:id;primary_key" json:"id"`
	Name                  string     `gorm:"column:name" json:"name"`
	DurationValue         int        `gorm:"column:duration_value" json:"duration_value"`
	DurationUnit          string     `gorm:"column:duration_unit" json:"duration_unit"`
	Price                 float64    `gorm:"column:price" json:"price"`
	Category              string     `gorm:"column:category" json:"category"`
	LoyaltyPoint          int        `gorm:"column:loyalty_point" json:"loyalty_point"`
	Description           *string    `gorm:"column:description" json:"description"`
	AccessType            string     `gorm:"column:access_type" json:"access_type"`
	ClassAccessType       string     `gorm:"column:class_access_type" json:"class_access_type"`
	MaxSession            *int       `gorm:"column:max_session" json:"max_session"`
	AllowUnlimitedSession bool       `gorm:"column:allow_unlimited_session" json:"allow_unlimited_session"`
	AvailableFrom         *time.Time `gorm:"column:available_from" json:"available_from"`
	AvailableUntil        *time.Time `gorm:"column:available_until" json:"available_until"`
	QuotaMaxSold          *int       `gorm:"column:quota_max_sold" json:"quota_max_sold"`
	AlwaysAvailable       bool       `gorm:"column:always_available" json:"always_available"`
	Level                 int        `gorm:"column:level" json:"level"`
	IsDeleted             bool       `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy             *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy             *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt             time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy             *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (MembershipPlan) TableName() string {
	return "membership_plan"
}

type MembershipPlanSchedule struct {
	ID               int        `gorm:"column:id;primary_key" json:"id"`
	MembershipPlanID int        `gorm:"column:membership_plan_id" json:"membership_plan_id"`
	DayOfWeek        string     `gorm:"column:day_of_week" json:"day_of_week"`
	StartTime        string     `gorm:"column:start_time" json:"start_time"`
	EndTime          string     `gorm:"column:end_time" json:"end_time"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy        *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy        *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (MembershipPlanSchedule) TableName() string {
	return "membership_plan_schedule"
}

type Membership struct {
	ID                    int        `gorm:"column:id;primary_key" json:"id"`
	UserID                int        `gorm:"column:user_id" json:"user_id"`
	MembershipPlanID      int        `gorm:"column:membership_plan_id" json:"membership_plan_id"`
	ReferralUserMemberID  *int       `gorm:"column:referral_user_member_id" json:"referral_user_member_id"`
	ReferralUserStaffID   *int       `gorm:"column:referral_user_staff_id" json:"referral_user_staff_id"`
	StartDate             time.Time  `gorm:"column:start_date" json:"start_date"`
	EndDate               time.Time  `gorm:"column:end_date" json:"end_date"`
	SalesType             string     `gorm:"column:sales_type" json:"sales_type"`
	AdditionalFee         *float64   `gorm:"column:additional_fee" json:"additional_fee"`
	DiscountType          *string    `gorm:"column:discount_type" json:"discount_type"`
	DiscountAmount        *float64   `gorm:"column:discount_amount" json:"discount_amount"`
	DiscountPercent       *float64   `gorm:"column:discount_percent" json:"discount_percent"`
	ExtraDurationDays     *int       `gorm:"column:extra_duration_days" json:"extra_duration_days"`
	ExtraSession          *int       `gorm:"column:extra_session" json:"extra_session"`
	Note                  *string    `gorm:"column:note" json:"note"`
	FinalPrice            float64    `gorm:"column:final_price" json:"final_price"`
	Status                string     `gorm:"column:status" json:"status"`
	IsActive              bool       `gorm:"column:is_active" json:"is_active"`
	IsTransferred         bool       `gorm:"column:is_transferred" json:"is_transferred"`
	PreviousMembershipID  *int       `gorm:"column:previous_membership_id" json:"previous_membership_id"`
	RenewalReminderSentAt *time.Time `gorm:"column:renewal_reminder_sent_at" json:"renewal_reminder_sent_at"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy             *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt             time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy             *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt             time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy             *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (Membership) TableName() string {
	return "membership"
}

type MembershipCheckinLog struct {
	ID           int       `gorm:"column:id;primary_key" json:"id"`
	UserID       int       `gorm:"column:user_id" json:"user_id"`
	MembershipID int       `gorm:"column:membership_id" json:"membership_id"`
	CheckinTime  time.Time `gorm:"column:checkin_time" json:"checkin_time"`
	Latitude     *float64  `gorm:"column:latitude" json:"latitude"`
	Longitude    *float64  `gorm:"column:longitude" json:"longitude"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy    *string   `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    *string   `gorm:"column:updated_by" json:"updated_by"`
}

func (MembershipCheckinLog) TableName() string {
	return "membership_checkin_log"
}

type Checkin struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	CheckinTime time.Time `gorm:"column:checkin_time" json:"checkin_time"`
	Latitude    *float64  `gorm:"column:latitude" json:"latitude"`
	Longitude   *float64  `gorm:"column:longitude" json:"longitude"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Checkin) TableName() string {
	return "checkin"
}

type CheckinPTSession struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	PTSessionID int       `gorm:"column:ptsession_id" json:"ptsession_id"`
	CheckinTime time.Time `gorm:"column:checkin_time" json:"checkin_time"`
	Latitude    *float64  `gorm:"column:latitude" json:"latitude"`
	Longitude   *float64  `gorm:"column:longitude" json:"longitude"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (CheckinPTSession) TableName() string {
	return "checkin_pt_session"
}

type PTSessionPlan struct {
	ID                int        `gorm:"column:id;primary_key" json:"id"`
	Name              string     `gorm:"column:name" json:"name"`
	Duration          *int       `gorm:"column:duration" json:"duration"`
	DurationValue     int        `gorm:"column:duration_value" json:"duration_value"`
	DurationUnit      string     `gorm:"column:duration_unit" json:"duration_unit"`
	MaxSession        int        `gorm:"column:max_session" json:"max_session"`
	Price             float64    `gorm:"column:price" json:"price"`
	MinutesPerSession int        `gorm:"column:minutes_per_session" json:"minutes_per_session"`
	Description       *string    `gorm:"column:description" json:"description"`
	CreatedAt         time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at" json:"updated_at"`
	CreatedBy         *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy         *string    `gorm:"column:updated_by" json:"updated_by"`
	IsDeleted         bool       `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy         *string    `gorm:"column:deleted_by" json:"deleted_by"`
}

func (PTSessionPlan) TableName() string {
	return "pt_session_plan"
}

type PersonalTrainerSession struct {
	ID               int        `gorm:"column:id;primary_key" json:"id"`
	PTSessionPlanID  int        `gorm:"column:pt_session_plan_id" json:"pt_session_plan_id"`
	UserMemberID     int        `gorm:"column:user_member_id" json:"user_member_id"`
	UserPTID         int        `gorm:"column:user_pt_id" json:"user_pt_id"`
	JoinDate         time.Time  `gorm:"column:join_date" json:"join_date"`
	StartDate        time.Time  `gorm:"column:start_date" json:"start_date"`
	EndDate          time.Time  `gorm:"column:end_date" json:"end_date"`
	Status           string     `gorm:"column:status" json:"status"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	CreatedBy        *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy        *string    `gorm:"column:updated_by" json:"updated_by"`
	IsDeleted        bool       `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy        *string    `gorm:"column:deleted_by" json:"deleted_by"`
	Name             *string    `gorm:"column:name" json:"name"`
	RemainingSession *int       `gorm:"column:remaining_session" json:"remaining_session"`
	QrCode           *string    `gorm:"column:qr_code" json:"qr_code"`
	BookingStart     *time.Time `gorm:"column:booking_start" json:"booking_start"`
	BookingEnd       *time.Time `gorm:"column:booking_end" json:"booking_end"`
}

func (PersonalTrainerSession) TableName() string {
	return "personal_trainer_session"
}

type PTSessionBooking struct {
	ID                       int        `gorm:"column:id;primary_key" json:"id"`
	UserMemberID             int        `gorm:"column:user_member_id" json:"user_member_id"`
	PTSessionPlanID          int        `gorm:"column:pt_session_plan_id" json:"pt_session_plan_id"`
	PersonalTrainerSessionID int        `gorm:"column:personal_trainer_session_id" json:"personal_trainer_session_id"`
	BookingTime              time.Time  `gorm:"column:booking_time" json:"booking_time"`
	Status                   string     `gorm:"column:status" json:"status"`
	CreatedAt                time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                time.Time  `gorm:"column:updated_at" json:"updated_at"`
	CreatedBy                *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedBy                *string    `gorm:"column:updated_by" json:"updated_by"`
	IsDeleted                bool       `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt                *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy                *string    `gorm:"column:deleted_by" json:"deleted_by"`
}

func (PTSessionBooking) TableName() string {
	return "pt_session_booking"
}

type EventPlan struct {
	ID                      int        `gorm:"column:id;primary_key" json:"id"`
	Name                    string     `gorm:"column:name" json:"name"`
	AccessType              string     `gorm:"column:access_type" json:"access_type"`
	MaxVisitor              int        `gorm:"column:max_visitor" json:"max_visitor"`
	MinutesPerSession       int        `gorm:"column:minutes_per_session" json:"minutes_per_session"`
	Description             *string    `gorm:"column:description" json:"description"`
	UnlimitedMonthlySession bool       `gorm:"column:unlimited_monthly_session" json:"unlimited_monthly_session"`
	MonthlyLimit            int        `gorm:"column:monthly_limit" json:"monthly_limit"`
	UnlimitedDailySession   bool       `gorm:"column:unlimited_daily_session" json:"unlimited_daily_session"`
	DailyLimit              int        `gorm:"column:daily_limit" json:"daily_limit"`
	IsActive                bool       `gorm:"column:is_active" json:"is_active"`
	DeletedAt               *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy               *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt               time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy               *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt               time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy               *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (EventPlan) TableName() string {
	return "event_plan"
}

type Class struct {
	ID                  int        `gorm:"column:id;primary_key" json:"id"`
	EventPlanID         int        `gorm:"column:event_plan_id" json:"event_plan_id"`
	InstructorID        int        `gorm:"column:instructor_id" json:"instructor_id"`
	Name                *string    `gorm:"column:name" json:"name"`
	ClassDate           time.Time  `gorm:"column:class_date" json:"class_date"`
	StartTime           time.Time  `gorm:"column:start_time" json:"start_time"`
	EndTime             *time.Time `gorm:"column:end_time" json:"end_time"`
	ClassType           string     `gorm:"column:class_type" json:"class_type"`
	TotalManualCheckin  int        `gorm:"column:total_manual_checkin" json:"total_manual_checkin"`
	Notes               *string    `gorm:"column:notes" json:"notes"`
	IsRecurring         bool       `gorm:"column:is_recurring" json:"is_recurring"`
	RecurrenceDays      *string    `gorm:"column:recurrence_days" json:"recurrence_days"`
	RecurrenceStartTime *string    `gorm:"column:recurrence_start_time" json:"recurrence_start_time"`
	RecurrenceEndTime   *string    `gorm:"column:recurrence_end_time" json:"recurrence_end_time"`
	ValidFrom           *time.Time `gorm:"column:valid_from" json:"valid_from"`
	ValidUntil          *time.Time `gorm:"column:valid_until" json:"valid_until"`
	ParentClassID       *int       `gorm:"column:parent_class_id" json:"parent_class_id"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy           *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt           time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy           *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt           time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy           *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (Class) TableName() string {
	return "class"
}

type ClassPurchase struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	UserID       int        `gorm:"column:user_id" json:"user_id"`
	ClassID      int        `gorm:"column:class_id" json:"class_id"`
	PurchaseDate time.Time  `gorm:"column:purchase_date" json:"purchase_date"`
	Price        float64    `gorm:"column:price" json:"price"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy    *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (ClassPurchase) TableName() string {
	return "class_purchase"
}

type ClassAttendance struct {
	ID          int        `gorm:"column:id;primary_key" json:"id"`
	ClassID     int        `gorm:"column:class_id" json:"class_id"`
	MemberID    int        `gorm:"column:member_id" json:"member_id"`
	CheckedInAt *time.Time `gorm:"column:checked_in_at" json:"checked_in_at"`
	Status      string     `gorm:"column:status" json:"status"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy   *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy   *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (ClassAttendance) TableName() string {
	return "class_attendance"
}

type StaffSchedule struct {
	ID                 int        `gorm:"column:id;primary_key" json:"id"`
	StaffID            int        `gorm:"column:staff_id" json:"staff_id"`
	ScheduleDate       time.Time  `gorm:"column:schedule_date" json:"schedule_date"`
	StartTime          time.Time  `gorm:"column:start_time" json:"start_time"`
	EndTime            time.Time  `gorm:"column:end_time" json:"end_time"`
	ScheduleType       string     `gorm:"column:schedule_type" json:"schedule_type"`
	RelatedClassID     *int       `gorm:"column:related_class_id" json:"related_class_id"`
	RelatedPTSessionID *int       `gorm:"column:related_pt_session_id" json:"related_pt_session_id"`
	Title              *string    `gorm:"column:title" json:"title"`
	Notes              *string    `gorm:"column:notes" json:"notes"`
	Status             string     `gorm:"column:status" json:"status"`
	IsRecurring        bool       `gorm:"column:is_recurring" json:"is_recurring"`
	RecurrencePattern  *string    `gorm:"column:recurrence_pattern" json:"recurrence_pattern"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy          *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy          *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy          *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (StaffSchedule) TableName() string {
	return "staff_schedule"
}

type MembershipFreeze struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	MembershipID int        `gorm:"column:membership_id" json:"membership_id"`
	FreezeAt     time.Time  `gorm:"column:freeze_at" json:"freeze_at"`
	UnfreezeAt   *time.Time `gorm:"column:unfreeze_at" json:"unfreeze_at"`
	Fee          float64    `gorm:"column:fee" json:"fee"`
	Reason       *string    `gorm:"column:reason" json:"reason"`
	Status       string     `gorm:"column:status" json:"status"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy    *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy    *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (MembershipFreeze) TableName() string {
	return "membership_freeze"
}

type MembershipTransfer struct {
	ID               int        `gorm:"column:id;primary_key" json:"id"`
	FromMembershipID int        `gorm:"column:from_membership_id" json:"from_membership_id"`
	ToMembershipID   *int       `gorm:"column:to_membership_id" json:"to_membership_id"`
	FromUserID       int        `gorm:"column:from_user_id" json:"from_user_id"`
	ToUserID         int        `gorm:"column:to_user_id" json:"to_user_id"`
	TransferDate     time.Time  `gorm:"column:transfer_date" json:"transfer_date"`
	Fee              float64    `gorm:"column:fee" json:"fee"`
	Reason           *string    `gorm:"column:reason" json:"reason"`
	Status           string     `gorm:"column:status" json:"status"`
	ApprovedBy       *string    `gorm:"column:approved_by" json:"approved_by"`
	ApprovedAt       *time.Time `gorm:"column:approved_at" json:"approved_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy        *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy        *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (MembershipTransfer) TableName() string {
	return "membership_transfer"
}

type Transaction struct {
	ID               int        `gorm:"column:id;primary_key" json:"id"`
	InvoiceNumber    string     `gorm:"column:invoice_number" json:"invoice_number"`
	TransactionDate  time.Time  `gorm:"column:transaction_date" json:"transaction_date"`
	UserID           int        `gorm:"column:user_id" json:"user_id"`
	TransactionType  string     `gorm:"column:transaction_type" json:"transaction_type"`
	MembershipID     *int       `gorm:"column:membership_id" json:"membership_id"`
	PTSessionID      *int       `gorm:"column:pt_session_id" json:"pt_session_id"`
	ClassPurchaseID  *int       `gorm:"column:class_purchase_id" json:"class_purchase_id"`
	ItemName         string     `gorm:"column:item_name" json:"item_name"`
	ItemDescription  *string    `gorm:"column:item_description" json:"item_description"`
	Quantity         int        `gorm:"column:quantity" json:"quantity"`
	UnitPrice        float64    `gorm:"column:unit_price" json:"unit_price"`
	Subtotal         float64    `gorm:"column:subtotal" json:"subtotal"`
	DiscountAmount   float64    `gorm:"column:discount_amount" json:"discount_amount"`
	DiscountPercent  *float64   `gorm:"column:discount_percent" json:"discount_percent"`
	TaxAmount        float64    `gorm:"column:tax_amount" json:"tax_amount"`
	TaxPercent       float64    `gorm:"column:tax_percent" json:"tax_percent"`
	AdditionalFee    float64    `gorm:"column:additional_fee" json:"additional_fee"`
	TotalAmount      float64    `gorm:"column:total_amount" json:"total_amount"`
	PTName           *string    `gorm:"column:pt_name" json:"pt_name"`
	PaymentMethod    *string    `gorm:"column:payment_method" json:"payment_method"`
	PaymentStatus    string     `gorm:"column:payment_status" json:"payment_status"`
	PaymentDate      *time.Time `gorm:"column:payment_date" json:"payment_date"`
	InvoiceSent      bool       `gorm:"column:invoice_sent" json:"invoice_sent"`
	InvoiceSentAt    *time.Time `gorm:"column:invoice_sent_at" json:"invoice_sent_at"`
	AgreementSent    bool       `gorm:"column:agreement_sent" json:"agreement_sent"`
	AgreementSentAt  *time.Time `gorm:"column:agreement_sent_at" json:"agreement_sent_at"`
	ProcessedByName  *string    `gorm:"column:processed_by_name" json:"processed_by_name"`
	ProcessedByEmail *string    `gorm:"column:processed_by_email" json:"processed_by_email"`
	Notes            *string    `gorm:"column:notes" json:"notes"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy        *string    `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt        time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy        *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy        *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (Transaction) TableName() string {
	return "transaction"
}

type AuditLog struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	Timestamp    time.Time  `gorm:"column:timestamp" json:"timestamp"`
	Action       string     `gorm:"column:action" json:"action"`
	Entity       string     `gorm:"column:entity" json:"entity"`
	EntityID     *int       `gorm:"column:entity_id" json:"entity_id"`
	FlagName     string     `gorm:"column:flag_name" json:"flag_name"`
	UserName     string     `gorm:"column:user_name" json:"user_name"`
	UserEmail    *string    `gorm:"column:user_email" json:"user_email"`
	UserID       *int       `gorm:"column:user_id" json:"user_id"`
	RequestBody  *string    `gorm:"column:request_body" json:"request_body"`
	Changes      *string    `gorm:"column:changes" json:"changes"`
	IPAddress    *string    `gorm:"column:ip_address" json:"ip_address"`
	UserAgent    *string    `gorm:"column:user_agent" json:"user_agent"`
	Status       string     `gorm:"column:status" json:"status"`
	Message      *string    `gorm:"column:message" json:"message"`
	ErrorCode    *string    `gorm:"column:error_code" json:"error_code"`
	ErrorStack   *string    `gorm:"column:error_stack" json:"error_stack"`
	ResponseData *string    `gorm:"column:response_data" json:"response_data"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    *string    `gorm:"column:updated_by" json:"updated_by"`
}

func (AuditLog) TableName() string {
	return "audit_log"
}

// BackupMetadata untuk tracking last backup timestamp
type BackupMetadata struct {
	ID             int       `gorm:"column:id;primary_key" json:"id"`
	Table          string    `gorm:"column:table_name;unique" json:"table_name"`
	LastBackupAt   time.Time `gorm:"column:last_backup_at" json:"last_backup_at"`
	TotalRecords   int       `gorm:"column:total_records" json:"total_records"`
	NewRecords     int       `gorm:"column:new_records" json:"new_records"`
	UpdatedRecords int       `gorm:"column:updated_records" json:"updated_records"`
	BackupType     string    `gorm:"column:backup_type" json:"backup_type"` // initial, incremental
	Status         string    `gorm:"column:status" json:"status"`           // success, failed
	ErrorMessage   *string   `gorm:"column:error_message" json:"error_message"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BackupMetadata) TableName() string {
	return "backup_metadata"
}

// BackupLog untuk menyimpan history log backup
type BackupLog struct {
	ID             int       `gorm:"column:id;primary_key;auto_increment" json:"id"`
	BackupType     string    `gorm:"column:backup_type" json:"backup_type"` // initial, incremental
	StartTime      time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime        time.Time `gorm:"column:end_time" json:"end_time"`
	Duration       float64   `gorm:"column:duration" json:"duration"` // in seconds
	Status         string    `gorm:"column:status" json:"status"`     // success, failed, partial
	TotalNew       int       `gorm:"column:total_new" json:"total_new"`
	TotalUpdated   int       `gorm:"column:total_updated" json:"total_updated"`
	TotalRecords   int       `gorm:"column:total_records" json:"total_records"`
	TablesAffected int       `gorm:"column:tables_affected" json:"tables_affected"`
	ErrorMessage   *string   `gorm:"column:error_message" json:"error_message"`
	LogDetails     *string   `gorm:"column:log_details;type:text" json:"log_details"` // Full log text

	// Detail changes per table in JSONB format
	UserChanges                   *string `gorm:"column:user_changes;type:jsonb" json:"user_changes"`
	MembershipPlanChanges         *string `gorm:"column:membership_plan_changes;type:jsonb" json:"membership_plan_changes"`
	MembershipPlanScheduleChanges *string `gorm:"column:membership_plan_schedule_changes;type:jsonb" json:"membership_plan_schedule_changes"`
	MembershipChanges             *string `gorm:"column:membership_changes;type:jsonb" json:"membership_changes"`
	MembershipCheckinLogChanges   *string `gorm:"column:membership_checkin_log_changes;type:jsonb" json:"membership_checkin_log_changes"`
	CheckinChanges                *string `gorm:"column:checkin_changes;type:jsonb" json:"checkin_changes"`
	CheckinPTSessionChanges       *string `gorm:"column:checkin_pt_session_changes;type:jsonb" json:"checkin_pt_session_changes"`
	PTSessionPlanChanges          *string `gorm:"column:pt_session_plan_changes;type:jsonb" json:"pt_session_plan_changes"`
	PersonalTrainerSessionChanges *string `gorm:"column:personal_trainer_session_changes;type:jsonb" json:"personal_trainer_session_changes"`
	PTSessionBookingChanges       *string `gorm:"column:pt_session_booking_changes;type:jsonb" json:"pt_session_booking_changes"`
	EventPlanChanges              *string `gorm:"column:event_plan_changes;type:jsonb" json:"event_plan_changes"`
	ClassChanges                  *string `gorm:"column:class_changes;type:jsonb" json:"class_changes"`
	ClassPurchaseChanges          *string `gorm:"column:class_purchase_changes;type:jsonb" json:"class_purchase_changes"`
	ClassAttendanceChanges        *string `gorm:"column:class_attendance_changes;type:jsonb" json:"class_attendance_changes"`
	StaffScheduleChanges          *string `gorm:"column:staff_schedule_changes;type:jsonb" json:"staff_schedule_changes"`
	MembershipFreezeChanges       *string `gorm:"column:membership_freeze_changes;type:jsonb" json:"membership_freeze_changes"`
	MembershipTransferChanges     *string `gorm:"column:membership_transfer_changes;type:jsonb" json:"membership_transfer_changes"`
	TransactionChanges            *string `gorm:"column:transaction_changes;type:jsonb" json:"transaction_changes"`
	AuditLogChanges               *string `gorm:"column:audit_log_changes;type:jsonb" json:"audit_log_changes"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (BackupLog) TableName() string {
	return "backup_logs"
}
