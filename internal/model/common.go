package model

type AzureNetAppConfigPayloadConfigurations struct {
	NetworkFeatures *string `json:"networkFeatures,omitempty"`
}

type WeeklySchedule struct {
	Days *[]string `json:"days,omitempty"` // Days in a week to retain weekly backups for
}

type StorageConfigPayload struct {
	Provider          *string                   `json:"provider"`
	FsxNetAppConfig   *FsxNetAppConfigPayload   `json:"fsxNetAppConfig,omitempty"`
	AzureNetAppConfig *AzureNetAppConfigPayload `json:"azureNetAppConfig,omitempty"`
}

type CommonYearlySchedule struct {
	Dates          *[]int    `json:"dates,omitempty"` // Dates in a month to retain monthly backups
	LastDayOfMonth *bool     `json:"lastDayOfMonth,omitempty"`
	Months         *[]string `json:"months,omitempty"`
}

type APIError struct {
	Code             *string `json:"code,omitempty"`    // Status code for the error response
	Message          *string `json:"message,omitempty"` // Error message for API response
	Resolution       *string `json:"resolution,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
	ContextId        *string `json:"contextId,omitempty"`        // ContextId of API request
	SessionId        *string `json:"sessionId,omitempty"`        // SessionId of API request
	TessellErrorCode *string `json:"tessellErrorCode,omitempty"` // Unique error code specific to Tessell
}

type TessellDataflixFromTimeInfo struct {
	FromTime   *string                      `json:"fromTime,omitempty"` // Recoverability start timestamp
	ToTime     *string                      `json:"toTime,omitempty"`   // Recoverability end timestamp
	SharedWith *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}

type SnapshotRegionAvailability struct {
	Region *string `json:"region"`
	Status *string `json:"status"` // Database Backup Status
}

type SnapshotAvailabilityConfig struct {
	AvailabilityConfiguredManually *bool                            `json:"availabilityConfiguredManually,omitempty"`
	DAPId                          *string                          `json:"dapId,omitempty"`
	CloudAvailabilityConfig        *[]SnapshotCloudAvailabilityInfo `json:"cloudAvailabilityConfig,omitempty"`
}

type EntityAclSharingSummaryInfo struct {
	Users *[]string `json:"users,omitempty"`
}

type MonthlySchedule struct {
	CommonSchedule *DatesForEachMonth `json:"commonSchedule,omitempty"`
}

type APIMetadata struct {
	TimeZone   *string            `json:"timeZone,omitempty"`
	Records    *int               `json:"records,omitempty"`
	Pagination *APIPaginationInfo `json:"pagination,omitempty"`
}

type EntityAclSharingInfo struct {
	Users *[]EntityUserAclSharingInfo `json:"users,omitempty"`
}

type TimeFormat struct {
	Hour   *int `json:"hour"`
	Minute *int `json:"minute"`
}

type YearlySchedule struct {
	CommonSchedule        *CommonYearlySchedule `json:"commonSchedule,omitempty"`
	MonthSpecificSchedule *[]MonthWiseDates     `json:"monthSpecificSchedule,omitempty"`
}

type StandardRPOPolicy struct {
	RetentionDays          *int        `json:"retentionDays"`                    // Number of days for which the snapshot of DB Service would be retained
	IncludeTransactionLogs *bool       `json:"includeTransactionLogs,omitempty"` // Determines whether transaction logs should be retained to enable Point-In-Time Recovery (PITR) functionality
	SnapshotStartTime      *TimeFormat `json:"snapshotStartTime"`
}

type ScheduleInfo struct {
	BackupStartTime *TimeFormat      `json:"backupStartTime,omitempty"`
	DailySchedule   *DailySchedule   `json:"dailySchedule,omitempty"`
	WeeklySchedule  *WeeklySchedule  `json:"weeklySchedule,omitempty"`
	MonthlySchedule *MonthlySchedule `json:"monthlySchedule,omitempty"`
	YearlySchedule  *YearlySchedule  `json:"yearlySchedule,omitempty"`
}

type APIPaginationInfo struct {
	PageSize   *int `json:"pageSize,omitempty"`
	PageOffset *int `json:"pageOffset,omitempty"`
}

type APIStatus struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

type EntityUserAclSharingInfo struct {
	EmailId *string `json:"emailId,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type CustomRPOPolicy struct {
	Name     *string       `json:"name"` // Custom RPO policy name
	Schedule *ScheduleInfo `json:"schedule"`
}

type MonthWiseDates struct {
	Month *string `json:"month"` // Name of a month
	Dates *[]int  `json:"dates"`
}

type DatesForEachMonth struct {
	Dates          *[]int `json:"dates,omitempty"` // Dates in a month to retain monthly backups
	LastDayOfMonth *bool  `json:"lastDayOfMonth,omitempty"`
}

type AzureNetAppConfigPayload struct {
	AzureNetAppId  *string                                 `json:"azureNetAppId,omitempty"`  // Azure NetApp Id registered with Tessell
	CapacityPoolId *string                                 `json:"capacityPoolId,omitempty"` // Capacity pool Id of the Azure NetApp registered with Tessell
	Configurations *AzureNetAppConfigPayloadConfigurations `json:"configurations,omitempty"`
}

type DatabaseSnapshotRegionInfo struct {
	Region *string `json:"region"`           // The region name
	Status *string `json:"status,omitempty"` // The current status of the snapshot in the respective region
}

type SnapshotCloudAvailabilityInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]SnapshotRegionAvailability `json:"regions,omitempty"` // The list of regions and respective availability status
}

type TaskSummary struct {
	TaskId                *string            `json:"taskId,omitempty"`
	TaskType              *string            `json:"taskType,omitempty"`
	ResourceId            *string            `json:"resourceId,omitempty"`
	AssociatedResourceIds *[]string          `json:"associatedResourceIds,omitempty"`
	Details               *map[string]string `json:"details,omitempty"`
}

type TessellDataflixPITRInfo struct {
	Cloud   *string                             `json:"cloud,omitempty"`
	Regions *[]TessellDataflixPITRInfoForRegion `json:"regions,omitempty"`
}

type DeletionScheduleDTO struct {
	Id             *string                       `json:"id,omitempty"`
	DeleteAt       *string                       `json:"deleteAt"` // Time at which the DB Service should be deleted at
	DeletionConfig *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
}

type TessellDataflixPITRInfoForRegion struct {
	Region     *string                        `json:"region,omitempty"` // Region name
	TimeRanges *[]TessellDataflixFromTimeInfo `json:"timeRanges,omitempty"`
}

type FsxNetAppConfigPayload struct {
	FileSystemId *string `json:"fileSystemId"` // File System Id of the FSx NetApp registered with Tessell
	SvmId        *string `json:"svmId"`        // Storage Virtual Machine Id of the FSx NetApp registered with Tessell
}

type TessellServiceDeletionConfig struct {
	RetainAvailabilityMachine *bool `json:"retainAvailabilityMachine,omitempty"` // If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained
}

type RegionInfo struct {
	Region            *string   `json:"region"` // The cloud region name
	AvailabilityZones *[]string `json:"availabilityZones,omitempty"`
}

type CloudRegionInfo struct {
	Cloud   *string       `json:"cloud"`
	Regions *[]RegionInfo `json:"regions,omitempty"` // The regions details
}

type DailySchedule struct {
	BackupsPerDay *int `json:"backupsPerDay,omitempty"` // The number of backups to be captured per day.
}

type DatabaseSnapshotCloudRegionInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]DatabaseSnapshotRegionInfo `json:"regions,omitempty"` // Region specific availability details for the snapshot
}

type BackupDatabaseInfo struct {
	Id     *string `json:"id,omitempty"`     // ID of the database
	Name   *string `json:"name,omitempty"`   // Name of the database
	Status *string `json:"status,omitempty"` // Status of the database as of capture of this snapshot
}
