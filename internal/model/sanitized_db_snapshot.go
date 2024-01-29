package model

type SanitizedDatabaseSnapshotSanitizationInfo struct {
	SourceSnapshotId       *string `json:"sourceSnapshotId,omitempty"`       // ID of the as-is snapshot from which it was sanitized
	SanitizationScheduleId *string `json:"sanitizationScheduleId,omitempty"` // ID of the Sanitization Schedule which has created this snapshot
	SanitizationSchedule   *string `json:"sanitizationSchedule,omitempty"`   // Name of the Sanitization Schedule which has created this snapshot
	SanitizationScriptId   *string `json:"sanitizationScriptId,omitempty"`   // ID of the script which was used to create this snapshot
	SanitizationScript     *string `json:"sanitizationScript,omitempty"`     // Name of the script which was used to create this snapshot
	ScriptVersion          *string `json:"scriptVersion,omitempty"`          // Version of the script which was used to create this snapshot
}

type SanitizedDatabaseSnapshot struct {
	Id                 *string                                    `json:"id,omitempty"`                 // ID of the sanitized snapshot
	Name               *string                                    `json:"name,omitempty"`               // Name of the sanitized snapshot
	SnapshotTime       *string                                    `json:"snapshotTime,omitempty"`       // Capture time of the source snapshot from which this sanitized snapshot is created
	Status             *string                                    `json:"status,omitempty"`             // Database Backup Status
	Size               *int                                       `json:"size,omitempty"`               // Size of this snapshot (in bytes)
	Manual             *bool                                      `json:"manual,omitempty"`             // Specifies whether this snapshot is created based on a manual user request or through an automated schedule
	CloudAvailability  *[]DatabaseSnapshotCloudRegionInfo         `json:"cloudAvailability,omitempty"`  // The cloud and region information where this snapshot has been made available at
	AvailabilityConfig *[]SnapshotAvailabilityConfig              `json:"availabilityConfig,omitempty"` // The config information for cloud and region availability for this snapshot
	SanitizationInfo   *SanitizedDatabaseSnapshotSanitizationInfo `json:"sanitizationInfo,omitempty"`
	Databases          *[]BackupDatabaseInfo                      `json:"databases,omitempty"` // The databases that are captured as part of the snapshot
	SharedWith         *EntityAclSharingSummaryInfo               `json:"sharedWith,omitempty"`
	BackupStatus       *string                                    `json:"backupStatus,omitempty"`
}

type GetSanitizedDatabaseSnapshotsResponse struct {
	AvailabilityMachineId *string                      `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId      *string                      `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName           *string                      `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType            *string                      `json:"engineType,omitempty"`            // Database Engine Type
	TimeZone              *string                      `json:"timeZone,omitempty"`              // Timezone applicable for timestamps that are returned in this response
	Owner                 *string                      `json:"owner,omitempty"`                 // Owner of the Availability Machine
	Snapshots             *[]SanitizedDatabaseSnapshot `json:"snapshots,omitempty"`             // Catalog information for the available Sanitized Snapshots
}
