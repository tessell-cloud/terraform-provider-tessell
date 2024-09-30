package model

type TessellDMMAvailabilityServiceView struct {
	AvailabilityMachineId *string              `json:"availabilityMachineId,omitempty"`
	AvailabilityMachine   *string              `json:"availabilityMachine,omitempty"` // Associated Availability Machine Name
	Topology              *[]DBServiceTopology `json:"topology,omitempty"`            // The availability location details: cloudAccount to region
	RPOSLAStatus          *string              `json:"rpoSlaStatus,omitempty"`        // The availability status
	SLA                   *string              `json:"sla,omitempty"`                 // Associated SLA
	SLARetentionInfo      *TamRetentionInfo    `json:"slaRetentionInfo,omitempty"`
	Schedule              *ScheduleInfo        `json:"schedule,omitempty"`
}

type DBServiceTopology struct {
	Type              *string   `json:"type,omitempty"`
	CloudType         *string   `json:"cloudType,omitempty"`
	Region            *string   `json:"region,omitempty"`
	AvailabilityZones *[]string `json:"availabilityZones,omitempty"`
}

type TamRetentionInfo struct {
	PITR    *int `json:"pitr,omitempty"`    // Retention time (in days) for Point-In-Time recoverability
	Daily   *int `json:"daily,omitempty"`   // Retention time (in days) to retain daily snapshots
	Weekly  *int `json:"weekly,omitempty"`  // Retention time (number of weeks) to retain weekly snapshots
	Monthly *int `json:"monthly,omitempty"` // Retention time (number of months) to retain monthly snapshots
	Yearly  *int `json:"yearly,omitempty"`  // Retention time (number of years) to retain yearly snapshots
}

type TessellDAPServiceDTO struct {
	Id                    *string               `json:"id,omitempty"`                    // ID of the Access Policy
	Name                  *string               `json:"name,omitempty"`                  // Name of the Access Policy
	AvailabilityMachineId *string               `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId      *string               `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName           *string               `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType            *string               `json:"engineType,omitempty"`            // Database engine type of the associated DB Service
	ContentType           *string               `json:"contentType,omitempty"`           // Content Type for the Data Access Policy
	Status                *string               `json:"status,omitempty"`                // Database Access Policy Status
	ContentInfo           *DAPContentInfo       `json:"contentInfo,omitempty"`
	CloudAvailability     *[]CloudRegionInfo    `json:"cloudAvailability,omitempty"` // The cloud and region information where the data is being managed by this Access Policy
	DataAccessConfig      *DAPRetentionInfo     `json:"dataAccessConfig,omitempty"`
	Owner                 *string               `json:"owner,omitempty"`            // Owner of the Access Policy
	LoggedInUserRole      *string               `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the Availability Machine
	SharedWith            *EntityAclSharingInfo `json:"sharedWith,omitempty"`
	DateCreated           *string               `json:"dateCreated,omitempty"`  // Timestamp when this Access Policy was created at
	DateModified          *string               `json:"dateModified,omitempty"` // Timestamp when this Access Policy was last updated at
}

type DAPContentInfo struct {
	AsIsContent      *AsIsDAPContent         `json:"asIsContent,omitempty"`
	SanitizedContent *SanitizationDAPContent `json:"sanitizedContent,omitempty"`
	BackupContent    *BackupDAPContent       `json:"backupContent,omitempty"`
}

type AsIsDAPContent struct {
	Automated *bool            `json:"automated,omitempty"` // Share the automated as-is snapshots. This is exclusive with manual specification.
	Manual    *[]DAPManualInfo `json:"manual,omitempty"`    // The list of snapshots that are to be shared as part of this access policy
}

type DAPManualInfo struct {
	Id           *string `json:"id,omitempty"`           // The DB Service snapshot id
	Name         *string `json:"name,omitempty"`         // The DB Service snapshot name
	CreationTime *string `json:"creationTime,omitempty"` // DB Service snapshot capture time
	SharedAt     *string `json:"sharedAt,omitempty"`     // The timestamp when the snapshot was added to DAP for sharing
}

type SanitizationDAPContent struct {
	Automated *SanitizationDAPContentAutomated `json:"automated,omitempty"`
	Manual    *[]DAPManualInfo                 `json:"manual,omitempty"` // The list of sanitized snapshots that are to be shared as part of this access policy
}

type SanitizationDAPContentAutomated struct {
	SanitizationScheduleId *string `json:"sanitizationScheduleId"` // Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.
}

type BackupDAPContent struct {
	Automated *bool            `json:"automated,omitempty"` // Share the automated backups. This is exclusive with manual specification.
	Manual    *[]DAPManualInfo `json:"manual,omitempty"`    // The list of backups that are to be shared as part of this access policy
}

type DAPRetentionInfo struct {
	PITR         *int `json:"pitr,omitempty"`         // Retention time (in days) for Point-In-Time recoverability
	DailyBackups *int `json:"dailyBackups,omitempty"` // Retention time (in days) to retain daily snapshots
}

type TessellCloneSummaryInfo struct {
	Id                *string            `json:"id,omitempty"`
	Name              *string            `json:"name"`                   // Name of the clone database
	Subscription      *string            `json:"subscription,omitempty"` // Clone&#39;s subscription name
	ComputeType       *string            `json:"computeType,omitempty"`  // Clone&#39;s compute type
	Status            *string            `json:"status,omitempty"`       // Status of the clone database
	CloudAvailability *[]CloudRegionInfo `json:"cloudAvailability,omitempty"`
	CloneInfo         *map[string]string `json:"cloneInfo,omitempty"`   // Miscellaneous information
	Owner             *string            `json:"owner,omitempty"`       // The user who created database clone
	DateCreated       *string            `json:"dateCreated,omitempty"` // Timestamp when the entity was created
}

type BackupDownloadConfig struct {
	AllowBackupDownloadsForAllUsers *bool `json:"allowBackupDownloadsForAllUsers,omitempty"` // Allow all users to download the backup, if false only owner/co-owner(s) will be allowed
	AllowBackupDownloads            *bool `json:"allowBackupDownloads,omitempty"`            // Allow download of the backup for owner/co-owner of the AM
}

type TessellDMMServiceConsumerDTO struct {
	Id                   *string                            `json:"id,omitempty"`                  // ID of the Availability Machine
	TessellServiceId     *string                            `json:"tessellServiceId,omitempty"`    // ID of the DB Service that is associated with the Availability Machine
	ServiceName          *string                            `json:"serviceName,omitempty"`         // Name of the DB Service that is associated with the Availability Machine
	Tenant               *string                            `json:"tenant,omitempty"`              // ID of the tenant under which this Availability Machine is effective
	Subscription         *string                            `json:"subscription,omitempty"`        // Name of the subscription under which the associated DB Service is hosted
	EngineType           *string                            `json:"engineType,omitempty"`          // Database Engine Type
	DataIngestionStatus  *string                            `json:"dataIngestionStatus,omitempty"` // Availability Machine&#39;s data ingestion status
	UserId               *string                            `json:"userId,omitempty"`              // User details representing the owner for the Availability Machine
	Owner                *string                            `json:"owner,omitempty"`               // User details representing the owner for the Availability Machine
	LoggedInUserRole     *string                            `json:"loggedInUserRole,omitempty"`    // The role of the logged in user for accessing this Availability Machine
	SharedWith           *EntityAclSharingInfo              `json:"sharedWith,omitempty"`
	CloudAvailability    *[]CloudRegionInfo                 `json:"cloudAvailability,omitempty"` // Availability Machine manages data across multiple regions within a cloud. This sections provides information about the cloud and regions where this Availability Machine is managing the data.
	RPOSLA               *TessellDMMAvailabilityServiceView `json:"rpoSla,omitempty"`
	DAPs                 *[]TessellDAPServiceDTO            `json:"daps,omitempty"`         // The Access Policies (DAP) that have configured for this Availability Machine
	Clones               *[]TessellCloneSummaryInfo         `json:"clones,omitempty"`       // The clone DB Services that have been created using contents (snapshots, Sanitized Snapshots, PITR, backups) from this Availability Machine
	DateCreated          *string                            `json:"dateCreated,omitempty"`  // The timestamp when the Availability Machine was incarnated
	DateModified         *string                            `json:"dateModified,omitempty"` // The timestamp when the Availability Machine was last updated
	Tsm                  *bool                              `json:"tsm,omitempty"`          // Specify whether the associated DB Service is created using TSM compute type
	BackupDownloadConfig *BackupDownloadConfig              `json:"backupDownloadConfig,omitempty"`
	StorageConfig        *StorageConfigPayload              `json:"storageConfig,omitempty"`
}

type GetDMMsServiceView struct {
	Metadata *APIMetadata                    `json:"metadata,omitempty"`
	Response *[]TessellDMMServiceConsumerDTO `json:"response,omitempty"`
}
