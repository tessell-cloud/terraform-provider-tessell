package model

type BackupSourceInfo struct {
	SourceSnapshotId *string `json:"sourceSnapshotId,omitempty"` // snapshot from which backup was created
	SnapshotName     *string `json:"snapshotName,omitempty"`
	SnapshotTime     *string `json:"snapshotTime,omitempty"` // snapshot creation time
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
	ExpireAt *string `json:"expireAt,omitempty"` // time till backup will be live
}

type DatabaseBackup struct {
	Id                 *string                       `json:"id,omitempty"`         // DB Service backup Id
	Name               *string                       `json:"name,omitempty"`       // DB Service backup name
	BackupTime         *string                       `json:"backupTime,omitempty"` // DB Service backup capture time
	Status             *string                       `json:"status,omitempty"`     // Database Backup Status
	Size               *int                          `json:"size,omitempty"`       // Backup size in bytes
	Manual             *bool                         `json:"manual,omitempty"`     // Specifies whether the backup is captured manually
	CloudAvailability  *[]CloudRegionInfo            `json:"cloudAvailability,omitempty"`
	AvailabilityConfig *[]SnapshotAvailabilityConfig `json:"availabilityConfig,omitempty"`
	Databases          *[]BackupDatabaseInfo         `json:"databases,omitempty"` // The databases that are captured as part of the backup
	BackupInfo         *BackupSourceInfo             `json:"backupInfo,omitempty"`
	SharedWith         *DatabaseBackupSharedWith     `json:"sharedWith,omitempty"`
	DownloadUrlStatus  *string                       `json:"downloadUrlStatus,omitempty"`
}

type GetDatabaseBackupsResponse struct {
	AvailabilityMachineId *string           `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string           `json:"tessellServiceId,omitempty"`
	ServiceName           *string           `json:"serviceName,omitempty"`
	EngineType            *string           `json:"engineType,omitempty"` // Database Engine Type
	TimeZone              *string           `json:"timeZone,omitempty"`   // Output timezone
	Owner                 *string           `json:"owner,omitempty"`      // Owner of the Availability Machine
	Backups               *[]DatabaseBackup `json:"backups,omitempty"`
}
