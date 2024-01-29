package model

type APIPaginationInfo struct {
	PageSize   *int `json:"pageSize,omitempty"`
	PageOffset *int `json:"pageOffset,omitempty"`
}

type EntityUserAclSharingInfo struct {
	EmailId *string `json:"emailId,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type DatabaseSnapshotRegionInfo struct {
	Region *string `json:"region"`           // The region name
	Status *string `json:"status,omitempty"` // The current status of the snapshot in the respective region
}

type SnapshotCloudAvailabilityInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]SnapshotRegionAvailability `json:"regions,omitempty"` // The list of regions and respective avaoilability status
}

type TaskSummary struct {
	TaskId     *string            `json:"taskId,omitempty"`
	TaskType   *string            `json:"taskType,omitempty"`
	ResourceId *string            `json:"resourceId,omitempty"`
	Details    *map[string]string `json:"details,omitempty"`
}

type TessellDataflixPITRInfo struct {
	Cloud   *string                             `json:"cloud,omitempty"`
	Regions *[]TessellDataflixPITRInfoForRegion `json:"regions,omitempty"`
}

type TessellDataflixPITRInfoForRegion struct {
	Region     *string                        `json:"region,omitempty"` // Region name
	TimeRanges *[]TessellDataflixFromTimeInfo `json:"timeRanges,omitempty"`
}

type APIError struct {
	Code           *string          `json:"code,omitempty"`    // Status code for the error response
	Message        *string          `json:"message,omitempty"` // Error message for API response
	Details        *APIErrorDetails `json:"details,omitempty"`
	DefaultCodeSet *bool            `json:"defaultCodeSet,omitempty"`
	ContextId      *string          `json:"contextId,omitempty"` // ContextId of API request
	SessionId      *string          `json:"sessionId,omitempty"` // SessionId of API request
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

type RegionInfo struct {
	Region            *string   `json:"region"` // The cloud region name
	AvailabilityZones *[]string `json:"availabilityZones,omitempty"`
}

type CloudRegionInfo struct {
	Cloud   *string       `json:"cloud"`
	Regions *[]RegionInfo `json:"regions,omitempty"` // The regions details
}

type EntityAclSharingSummaryInfo struct {
	Users *[]string `json:"users,omitempty"`
}

type APIMetadata struct {
	TimeZone   *string            `json:"timeZone,omitempty"`
	Records    *int               `json:"records,omitempty"`
	Pagination *APIPaginationInfo `json:"pagination,omitempty"`
}

type EntityAclSharingInfo struct {
	Users *[]EntityUserAclSharingInfo `json:"users,omitempty"`
}

type APIErrorDetails struct {
	Resolution *string `json:"resolution,omitempty"` // Resolution detail for API exception
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
