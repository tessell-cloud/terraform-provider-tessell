package model

type SnapshotAvailabilityConfig struct {
	AvailabilityConfiguredManually *bool                            `json:"availabilityConfiguredManually,omitempty"`
	DAPId                          *string                          `json:"dapId,omitempty"`
	CloudAvailabilityConfig        *[]SnapshotCloudAvailabilityInfo `json:"cloudAvailabilityConfig,omitempty"`
}

type SnapshotCloudAvailabilityInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]SnapshotRegionAvailability `json:"regions,omitempty"` // The list of regions and respective avaoilability status
}

type SnapshotRegionAvailability struct {
	Region *string `json:"region"`
	Status *string `json:"status"` // Database Backup Status
}

type DatabaseSnapshot struct {
	Id                 *string                       `json:"id,omitempty"`           // DB Service snapshot Id
	Name               *string                       `json:"name,omitempty"`         // DB Service snapshot name
	Description        *string                       `json:"description,omitempty"`  // Description for the snapshot
	SnapshotTime       *string                       `json:"snapshotTime,omitempty"` // DB Service snapshot capture time
	Status             *string                       `json:"status,omitempty"`       // Database Backup Status
	Size               *int                          `json:"size,omitempty"`         // Database Backup size in bytes
	Manual             *bool                         `json:"manual,omitempty"`       // Specifies whether the backup is captured manually
	CloudAvailability  *[]CloudRegionInfo            `json:"cloudAvailability,omitempty"`
	AvailabilityConfig *[]SnapshotAvailabilityConfig `json:"availabilityConfig,omitempty"`
	Databases          *[]BackupDatabaseInfo         `json:"databases,omitempty"` // The databases that are captured as part of the snapshot
	SharedWith         *EntityAclSharingSummaryInfo  `json:"sharedWith,omitempty"`
	BackupStatus       *string                       `json:"backupStatus,omitempty"`
}

type CreateDatabaseSnapshotTaskPayload struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

type APIStatus struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
