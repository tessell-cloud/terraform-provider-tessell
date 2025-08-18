package model

type DatabaseSnapshot struct {
	Id                 *string                            `json:"id,omitempty"`                 // ID of the snapshot
	Name               *string                            `json:"name,omitempty"`               // Name of the snapshot
	Description        *string                            `json:"description,omitempty"`        // Description of the snapshot
	SnapshotTime       *string                            `json:"snapshotTime,omitempty"`       // Capture time of the snapshot
	Status             *string                            `json:"status,omitempty"`             // Database Backup Status
	Size               *int                               `json:"size,omitempty"`               // Size of this snapshot (in bytes)
	Manual             *bool                              `json:"manual,omitempty"`             // Specifies whether this snapshot is captured as per manual user request or per automated schedule
	Incremental        *bool                              `json:"incremental,omitempty"`        // Specifies if Database Backup&#39;s is incremental
	CloudAvailability  *[]DatabaseSnapshotCloudRegionInfo `json:"cloudAvailability,omitempty"`  // The cloud and region information where this snapshot has been made available at
	AvailabilityConfig *[]SnapshotAvailabilityConfig      `json:"availabilityConfig,omitempty"` // The config information for cloud and region availability for this snapshot
	Databases          *[]BackupDatabaseInfo              `json:"databases,omitempty"`          // The databases that are captured as part of this snapshot
	SharedWith         *EntityAclSharingSummaryInfo       `json:"sharedWith,omitempty"`
	BackupStatus       *string                            `json:"backupStatus,omitempty"`
}

type CreateDatabaseSnapshotTaskPayload struct {
	Name        *string `json:"name"`                  // Name for the snapshot
	Description *string `json:"description,omitempty"` // Description for the snapshot
}

type GetDatabaseSnapshotsResponse struct {
	AvailabilityMachineId *string                    `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId      *string                    `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName           *string                    `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType            *string                    `json:"engineType,omitempty"`            // Database Engine Type
	TimeZone              *string                    `json:"timeZone,omitempty"`              // Timezone applicable for timestamps that are returned in this response
	Owner                 *string                    `json:"owner,omitempty"`                 // Owner of the Availability Machine
	PITRCatalog           *[]TessellDataflixPITRInfo `json:"pitrCatalog,omitempty"`           // Catalog information for the point-in-time recoverability
	Snapshots             *[]DatabaseSnapshot        `json:"snapshots,omitempty"`             // Catalog information for the available snapshots
}
