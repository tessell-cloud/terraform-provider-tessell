package model

type ApiMetadata struct {
	TimeZone   *string            `json:"timeZone,omitempty"`
	Records    *int               `json:"records,omitempty"`
	Pagination *ApiPaginationInfo `json:"pagination,omitempty"`
}

type EntityUserAclSharingInfo struct {
	EmailId *string `json:"emailId,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type TaskSummary struct {
	TaskId     *string            `json:"taskId,omitempty"`
	TaskType   *string            `json:"taskType,omitempty"`
	ResourceId *string            `json:"resourceId,omitempty"`
	Details    *map[string]string `json:"details,omitempty"`
}

type ApiPaginationInfo struct {
	PageSize   *int `json:"pageSize,omitempty"`
	PageOffset *int `json:"pageOffset,omitempty"`
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

type EntityAclSharingInfo struct {
	Users *[]EntityUserAclSharingInfo `json:"users,omitempty"`
}

type BackupDatabaseInfo struct {
	Id     *string `json:"id,omitempty"`     // Databases Id
	Name   *string `json:"name,omitempty"`   // Databases name
	Status *string `json:"status,omitempty"` // Databases status
}
