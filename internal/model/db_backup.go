package model

type BackupSourceInfo struct {
	SourceSnapshotId *string `json:"sourceSnapshotId,omitempty"` // ID of snapshot from which this backup was created
	SnapshotName     *string `json:"snapshotName,omitempty"`     // Name of snapshot from which this backup was created
	SnapshotTime     *string `json:"snapshotTime,omitempty"`     // Capture time of snapshot from which this backup was created
}

type DatabaseBackupSharedWith struct {
	Users *[]BackupUserInfo `json:"users,omitempty"`
}

type BackupUserInfo struct {
	UserEmail         *string       `json:"userEmail,omitempty"` // email of the user
	DownloadUrlStatus *string       `json:"downloadUrlStatus,omitempty"`
	ExpiryConfig      *ExpiryConfig `json:"expiryConfig,omitempty"`
}

type ExpiryConfig struct {
	ExpireAt *string `json:"expireAt,omitempty"` // time-to-live for the Pre auth url
}

type DatabaseBackup struct {
	Id                 *string                       `json:"id,omitempty"`                 // ID of the backup
	Name               *string                       `json:"name,omitempty"`               // Name of the backup
	BackupTime         *string                       `json:"backupTime,omitempty"`         // Backup capture time
	Status             *string                       `json:"status,omitempty"`             // Database Backup Status
	Size               *int                          `json:"size,omitempty"`               // Size of this backup (in bytes)
	Manual             *bool                         `json:"manual,omitempty"`             // Specifies whether the backup is captured as per manual user request or per automated schedule
	IsIncremental      *bool                         `json:"isIncremental,omitempty"`      // Specifies whether this backup is incremental
	CloudAvailability  *[]CloudRegionInfo            `json:"cloudAvailability,omitempty"`  // The cloud and region information where this backup has been made available at
	AvailabilityConfig *[]SnapshotAvailabilityConfig `json:"availabilityConfig,omitempty"` // The config information for cloud and region availability for this backup
	Databases          *[]BackupDatabaseInfo         `json:"databases,omitempty"`          // The databases that are captured as part of this backup
	BackupSource       *string                       `json:"backupSource,omitempty"`
	BackupInfo         *BackupSourceInfo             `json:"backupInfo,omitempty"`
	SharedWith         *DatabaseBackupSharedWith     `json:"sharedWith,omitempty"`
	DownloadUrlStatus  *string                       `json:"downloadUrlStatus,omitempty"`
}

type GetDatabaseBackupsResponse struct {
	AvailabilityMachineId *string           `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId      *string           `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName           *string           `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType            *string           `json:"engineType,omitempty"`            // Database Engine Type
	TimeZone              *string           `json:"timeZone,omitempty"`              // Timezone applicable for timestamps that are returned in this response
	Owner                 *string           `json:"owner,omitempty"`                 // Owner of the Availability Machine
	Backups               *[]DatabaseBackup `json:"backups,omitempty"`               // Catalog information for the available backups
}
