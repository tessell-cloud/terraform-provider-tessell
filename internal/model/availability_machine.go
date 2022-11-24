package model

type TessellDmmServiceConsumerDTO struct {
	Id                  *string                            `json:"id,omitempty"`
	TessellServiceId    *string                            `json:"tessellServiceId,omitempty"`
	ServiceName         *string                            `json:"serviceName,omitempty"`
	Tenant              *string                            `json:"tenant,omitempty"`              // Dmm&#39;s tenancy details
	Subscription        *string                            `json:"subscription,omitempty"`        // Dmm&#39;s subscription name
	EngineType          *string                            `json:"engineType,omitempty"`          // Database Engine Type
	DataIngestionStatus *string                            `json:"dataIngestionStatus,omitempty"` // Availability Machine&#39;s data ingestion status
	UserId              *string                            `json:"userId,omitempty"`              // Data Management Machine&#39;s user
	Owner               *string                            `json:"owner,omitempty"`               // Availability Machine&#39;s owner
	LoggedInUserRole    *string                            `json:"loggedInUserRole,omitempty"`    // The role of the logged in user for accessing the Availability Machine
	SharedWith          *EntityAclSharingInfo              `json:"sharedWith,omitempty"`
	CloudAvailability   *[]CloudRegionInfo                 `json:"cloudAvailability,omitempty"`
	RpoSla              *TessellDmmAvailabilityServiceView `json:"rpoSla,omitempty"`
	Daps                *[]TessellDapServiceDTO            `json:"daps,omitempty"`
	Clones              *[]TessellCloneSummaryInfo         `json:"clones,omitempty"` // Clone databases that are created from this Availability Machine
	DateCreated         *string                            `json:"dateCreated,omitempty"`
	DateModified        *string                            `json:"dateModified,omitempty"`
}

type TessellDmmAvailabilityServiceView struct {
	AvailabilityMachineId *string            `json:"availabilityMachineId,omitempty"`
	AvailabilityMachine   *string            `json:"availabilityMachine,omitempty"` // Associated Availability Machine Name
	CloudAvailability     *[]CloudRegionInfo `json:"cloudAvailability,omitempty"`   // The availability location details: cloudAccount to region
	RpoSlaStatus          *string            `json:"rpoSlaStatus,omitempty"`        // The availability status
	Sla                   *string            `json:"sla,omitempty"`                 // Associated SLA
	Schedule              *ScheduleInfo      `json:"schedule,omitempty"`
}

type ScheduleInfo struct {
	BackupStartTime *ScheduleTimeFormat `json:"backupStartTime,omitempty"`
	DailySchedule   *DailySchedule      `json:"dailySchedule,omitempty"`
}

type ScheduleTimeFormat struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

type DailySchedule struct {
	BackupsPerDay    *int                  `json:"backupsPerDay,omitempty"`    // The number of backups to be captured per day, this is exclusive with &#39;backupStartTimes&#39;
	BackupStartTimes *[]ScheduleTimeFormat `json:"backupStartTimes,omitempty"` // List of times when backup(s) has to be captured at. If this is specified, the &#39;backupStartTime&#39; (at top level) value would be overridern by the first entry in this list
}

type TessellDapServiceDTO struct {
	Id                    *string                   `json:"id,omitempty"`
	Name                  *string                   `json:"name,omitempty"`
	AvailabilityMachineId *string                   `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                   `json:"tessellServiceId,omitempty"`
	ServiceName           *string                   `json:"serviceName,omitempty"`
	EngineType            *string                   `json:"engineType,omitempty"`
	ContentType           *string                   `json:"contentType,omitempty"` // Content Type for the Data Access Policy
	Status                *string                   `json:"status,omitempty"`      // Database Access Policy Status
	ContentInfo           *DapContentInfo           `json:"contentInfo,omitempty"`
	CloudAvailability     *[]CloudRegionInfo        `json:"cloudAvailability,omitempty"`
	DataAccessConfig      *RetentionAndScheduleInfo `json:"dataAccessConfig,omitempty"`
	Owner                 *string                   `json:"owner,omitempty"`
	LoggedInUserRole      *string                   `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the Availability Machine
	SharedWith            *EntityAclSharingInfo     `json:"sharedWith,omitempty"`
	DateCreated           *string                   `json:"dateCreated,omitempty"`
	DateModified          *string                   `json:"dateModified,omitempty"`
}

type DapContentInfo struct {
	AsIsContent      *AsIsDapContent         `json:"asIsContent,omitempty"`
	SanitizedContent *SanitizationDapContent `json:"sanitizedContent,omitempty"`
}

type AsIsDapContent struct {
	Automated *bool              `json:"automated,omitempty"` // Share the automated as-is snapshots. This is exclusive with manual specification.
	Manual    *[]DapSnapshotInfo `json:"manual,omitempty"`    // The list of snapshots that are to be shared as part of this access policy
}

type DapSnapshotInfo struct {
	SnapshotId   *string `json:"snapshotId,omitempty"`   // The DB Service snapshot id
	SnapshotName *string `json:"snapshotName,omitempty"` // The DB Service snapshot name
	SnapshotTime *string `json:"snapshotTime,omitempty"` // DB Service snapshot capture time
	SharedAt     *string `json:"sharedAt,omitempty"`     // The timestamp when the snapshot was added to DAP for sharing
}

type SanitizationDapContent struct {
	Automated *SanitizationDapContentAutomated `json:"automated,omitempty"`
	Manual    *[]DapSnapshotInfo               `json:"manual,omitempty"` // The list of sanitized snapshots that are to be shared as part of this access policy
}

type SanitizationDapContentAutomated struct {
	SanitizationScheduleId *string `json:"sanitizationScheduleId"` // Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.
}

type RetentionAndScheduleInfo struct {
	DailyBackups *int `json:"dailyBackups,omitempty"` // Number of daily backups to replicate
}

type TessellCloneSummaryInfo struct {
	Id                *string            `json:"id,omitempty"`
	Name              *string            `json:"name"`                   // Name of the clone database
	Subscription      *string            `json:"subscription,omitempty"` // Clone&#39;s subsription name
	Status            *string            `json:"status,omitempty"`       // Status of the clone database
	CloudAvailability *[]CloudRegionInfo `json:"cloudAvailability,omitempty"`
	CloneInfo         *map[string]string `json:"cloneInfo,omitempty"`   // Miscellaneous information
	Owner             *string            `json:"owner,omitempty"`       // The user who created database clone
	DateCreated       *string            `json:"dateCreated,omitempty"` // Timestamp when the entity was created
}

type GetDmmsServiceView struct {
	Metadata *ApiMetadata                    `json:"metadata,omitempty"`
	Response *[]TessellDmmServiceConsumerDTO `json:"response,omitempty"`
}
