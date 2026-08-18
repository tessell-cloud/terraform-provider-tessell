package model

type DatabaseProfileOptionType struct {
	ApplyNow       *bool                               `json:"applyNow"`
	Name           *string                             `json:"name"` // Option name
	OptionSettings *[]DatabaseProfileOptionSettingType `json:"optionSettings"`
}

type DatabaseProfileOptionSettingType struct {
	Name     *string `json:"name"`
	Value    *string `json:"value,omitempty"`
	IsGlobal *bool   `json:"isGlobal,omitempty"`
}

type DatabaseOptionProfileMetadata struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type DatabaseOptionProfileDriverInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type TerraformDBOptionProfile struct {
	Id             *string                          `json:"id,omitempty"`          // Tessell generated UUID for the entity
	Name           *string                          `json:"name"`                  // Name of the entity
	Description    *string                          `json:"description,omitempty"` // Description of an Options Profile
	EngineType     *string                          `json:"engineType,omitempty"`
	Status         *string                          `json:"status,omitempty"`
	MaturityStatus *string                          `json:"maturityStatus,omitempty"`
	OptionTypeId   *string                          `json:"optionTypeId,omitempty"` // Tessell option type UUID for the entity
	Owner          *string                          `json:"owner,omitempty"`
	TenantId       *string                          `json:"tenantId,omitempty"`
	Version        *string                          `json:"version,omitempty"` // Database Option Profile&#39;s version
	Options        *[]DatabaseProfileOptionType     `json:"options,omitempty"` // Database Option Profile&#39;s associated options
	Metadata       *DatabaseOptionProfileMetadata   `json:"metadata,omitempty"`
	DriverInfo     *DatabaseOptionProfileDriverInfo `json:"driverInfo,omitempty"`
}

type TessellDatabaseOptionProfileConsumptionDTO struct {
	Description    *string                          `json:"description,omitempty"`
	DriverInfo     *DatabaseOptionProfileDriverInfo `json:"driverInfo,omitempty"`
	EngineType     *string                          `json:"engineType,omitempty"`
	Id             *string                          `json:"id,omitempty"`           // Tessell generated UUID for the entity
	OptionTypeId   *string                          `json:"optionTypeId,omitempty"` // Tessell option type UUID for the entity
	Metadata       *DatabaseOptionProfileMetadata   `json:"metadata,omitempty"`
	Name           *string                          `json:"name"` // Name of the entity
	MaturityStatus *string                          `json:"maturityStatus,omitempty"`
	Options        *[]DatabaseProfileOptionType     `json:"options,omitempty"` // Database Option Profile&#39;s associated options
	Owner          *string                          `json:"owner,omitempty"`
	TenantId       *string                          `json:"tenantId,omitempty"`
	Version        *string                          `json:"version,omitempty"` // Database Option Profile&#39;s version
}

type TessellDatabaseOptionProfileConsumptionListResponse struct {
	Response *[]TessellDatabaseOptionProfileConsumptionDTO `json:"response,omitempty"`
	Metadata *APIMetadata                                  `json:"metadata,omitempty"`
}
