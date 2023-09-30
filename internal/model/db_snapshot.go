package model

type DatabaseSnapshot struct {
	Id                 *string                            `json:"id,omitempty"`           // DB Service snapshot Id
	Name               *string                            `json:"name,omitempty"`         // DB Service snapshot name
	Description        *string                            `json:"description,omitempty"`  // Description for the snapshot
	SnapshotTime       *string                            `json:"snapshotTime,omitempty"` // DB Service snapshot capture time
	Status             *string                            `json:"status,omitempty"`       // Database Backup Status
	Size               *int                               `json:"size,omitempty"`         // Database Backup size in bytes
	Manual             *bool                              `json:"manual,omitempty"`       // Specifies whether the backup is captured manually
	CloudAvailability  *[]DatabaseSnapshotCloudRegionInfo `json:"cloudAvailability,omitempty"`
	AvailabilityConfig *[]SnapshotAvailabilityConfig      `json:"availabilityConfig,omitempty"`
	Databases          *[]BackupDatabaseInfo              `json:"databases,omitempty"` // The databases that are captured as part of the snapshot
	SharedWith         *EntityAclSharingSummaryInfo       `json:"sharedWith,omitempty"`
	BackupStatus       *string                            `json:"backupStatus,omitempty"`
}

type CreateDatabaseSnapshotTaskPayload struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

type APIStatus struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

type GetDatabaseSnapshotsResponse struct {
	AvailabilityMachineId *string                    `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                    `json:"tessellServiceId,omitempty"`
	ServiceName           *string                    `json:"serviceName,omitempty"`
	EngineType            *string                    `json:"engineType,omitempty"`  // Database Engine Type
	TimeZone              *string                    `json:"timeZone,omitempty"`    // Output timezone
	Owner                 *string                    `json:"owner,omitempty"`       // Owner of the Availability Machine
	PITRCatalog           *[]TessellDataflixPITRInfo `json:"pitrCatalog,omitempty"` // PITR availability catalog
	Snapshots             *[]DatabaseSnapshot        `json:"snapshots,omitempty"`
}
