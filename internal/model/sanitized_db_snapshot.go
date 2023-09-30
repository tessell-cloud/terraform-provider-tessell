package model

type SanitizedDatabaseSnapshotSanitizationInfo struct {
	SourceSnapshotId       *string `json:"sourceSnapshotId,omitempty"`       // The as-is snapshot from which it was sanitized
	SanitizationScheduleId *string `json:"sanitizationScheduleId,omitempty"` // Id of the sanitization schedule which has created this snapshot
	SanitizationSchedule   *string `json:"sanitizationSchedule,omitempty"`   // Name of the sanitization schedule which has created this snapshot
	SanitizationScriptId   *string `json:"sanitizationScriptId,omitempty"`   // Id of the script which was used to create this backup
	SanitizationScript     *string `json:"sanitizationScript,omitempty"`     // Name of the script which was used to create this backup
	ScriptVersion          *string `json:"scriptVersion,omitempty"`          // Version of the script which was used to create this backup
}

type SanitizedDatabaseSnapshot struct {
	Id                 *string                                    `json:"id,omitempty"`           // DB Service snapshot Id
	Name               *string                                    `json:"name,omitempty"`         // DB Service snapshot name
	SnapshotTime       *string                                    `json:"snapshotTime,omitempty"` // DB Service snapshot capture time
	Status             *string                                    `json:"status,omitempty"`       // Database Backup Status
	Size               *int                                       `json:"size,omitempty"`         // Snapshot size in bytes
	Manual             *bool                                      `json:"manual,omitempty"`       // Specifies whether the snapshot is captured manually
	CloudAvailability  *[]DatabaseSnapshotCloudRegionInfo         `json:"cloudAvailability,omitempty"`
	AvailabilityConfig *[]SnapshotAvailabilityConfig              `json:"availabilityConfig,omitempty"`
	SanitizationInfo   *SanitizedDatabaseSnapshotSanitizationInfo `json:"sanitizationInfo,omitempty"`
	Databases          *[]BackupDatabaseInfo                      `json:"databases,omitempty"` // The databases that are captured as part of the snapshot
	SharedWith         *EntityAclSharingSummaryInfo               `json:"sharedWith,omitempty"`
	BackupStatus       *string                                    `json:"backupStatus,omitempty"`
}

type GetSanitizedDatabaseSnapshotsResponse struct {
	AvailabilityMachineId *string                      `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                      `json:"tessellServiceId,omitempty"`
	ServiceName           *string                      `json:"serviceName,omitempty"`
	EngineType            *string                      `json:"engineType,omitempty"` // Database Engine Type
	TimeZone              *string                      `json:"timeZone,omitempty"`   // Output timezone
	Owner                 *string                      `json:"owner,omitempty"`      // Owner of the Availability Machine
	Snapshots             *[]SanitizedDatabaseSnapshot `json:"snapshots,omitempty"`
}
