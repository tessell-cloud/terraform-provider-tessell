package model

type DatabaseProfileParameterType struct {
	DataType      *string `json:"dataType"`
	DefaultValue  *string `json:"defaultValue"`
	ApplyType     *string `json:"applyType,omitempty"`
	Name          *string `json:"name"`
	Description   *string `json:"description,omitempty"`
	Value         *string `json:"value"`
	AllowedValues *string `json:"allowedValues,omitempty"`
	IsModified    *bool   `json:"isModified,omitempty"`
	IsFormulaType *bool   `json:"isFormulaType,omitempty"`
	Source        *string `json:"source,omitempty"`
	TopParameter  *bool   `json:"topParameter,omitempty"` // Boolean variable indicating a parameter is a most modified / key parameter
	IsModifiable  *bool   `json:"isModifiable,omitempty"`
	UsageType     *string `json:"usageType,omitempty"`
}

type TerraformDBParameterProfile struct {
	Id                 *string                         `json:"id,omitempty"`          // Tessell generated UUID for the entity
	VersionId          *string                         `json:"versionId,omitempty"`   // Tessell generated UUID for the entity
	Name               *string                         `json:"name"`                  // Name of the entity
	Description        *string                         `json:"description,omitempty"` // Database Parameter Profile description
	Oob                *bool                           `json:"oob,omitempty"`
	EngineType         *string                         `json:"engineType,omitempty"`
	FactoryParameterId *string                         `json:"factoryParameterId,omitempty"` // Tessell parameter type UUID for the entity
	Status             *string                         `json:"status,omitempty"`
	MaturityStatus     *string                         `json:"maturityStatus,omitempty"`
	Owner              *string                         `json:"owner,omitempty"`
	TenantId           *string                         `json:"tenantId,omitempty"`
	LoggedInUserRole   *string                         `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the db profile
	Parameters         *[]DatabaseProfileParameterType `json:"parameters,omitempty"`       // Parameter Profile&#39;s associated parameters
	UserId             *string                         `json:"userId,omitempty"`           // Database Parameter Profile&#39;s user id
	SharedWith         *EntityAclSharingInfo           `json:"sharedWith,omitempty"`
	DBVersion          *string                         `json:"dbVersion,omitempty"`    // Database Parameter Profile&#39;s version
	DateCreated        *string                         `json:"dateCreated,omitempty"`  // Timestamp when the entity was created
	DateModified       *string                         `json:"dateModified,omitempty"` // Timestamp when the entity was last modified
}

type DatabaseParameterProfileResponse struct {
	Id                 *string                             `json:"id,omitempty"`          // Tessell generated UUID for the entity
	VersionId          *string                             `json:"versionId,omitempty"`   // Tessell generated UUID for the entity
	Name               *string                             `json:"name"`                  // Name of the entity
	Description        *string                             `json:"description,omitempty"` // Database Parameter Profile description
	Oob                *bool                               `json:"oob,omitempty"`
	EngineType         *string                             `json:"engineType,omitempty"`
	EngineInfo         *DatabaseParameterEngineInfo        `json:"engineInfo,omitempty"`
	FactoryParameterId *string                             `json:"factoryParameterId,omitempty"` // Tessell parameter type UUID for the entity
	Status             *string                             `json:"status,omitempty"`
	MaturityStatus     *string                             `json:"maturityStatus,omitempty"`
	Owner              *string                             `json:"owner,omitempty"`
	Parameters         *[]DatabaseProfileParameterType     `json:"parameters,omitempty"` // Parameter Profile&#39;s associated parameters
	Metadata           *DatabaseParameterProfileMetadata   `json:"metadata,omitempty"`
	DriverInfo         *DatabaseParameterProfileDriverInfo `json:"driverInfo,omitempty"`
	UserId             *string                             `json:"userId,omitempty"`       // Database Parameter Profile&#39;s user id
	DBVersion          *string                             `json:"dbVersion,omitempty"`    // Database Parameter Profile&#39;s version
	DateCreated        *string                             `json:"dateCreated,omitempty"`  // Timestamp when the entity was created
	DateModified       *string                             `json:"dateModified,omitempty"` // Timestamp when the entity was last modified
	InfraType          *string                             `json:"infraType,omitempty"`
	IsLegacy           *bool                               `json:"isLegacy,omitempty"` // Whether this Parameter Profile is Legacy or not
}

type DatabaseParameterEngineInfo struct {
	Edition *string                            `json:"edition,omitempty"`
	Oracle  *DatabaseParameterEngineInfoOracle `json:"oracle,omitempty"`
}

type DatabaseParameterEngineInfoOracle struct {
	MultiTenancy *string `json:"multiTenancy,omitempty"`
}

type DatabaseParameterProfileMetadata struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type DatabaseParameterProfileDriverInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type APIErrorOps struct {
	Code             *string                   `json:"code,omitempty"`    // Status code for the error response
	Message          *string                   `json:"message,omitempty"` // Error message for API response
	Resolution       *string                   `json:"resolution,omitempty"`
	Timestamp        *string                   `json:"timestamp,omitempty"`
	ContextId        *string                   `json:"contextId,omitempty"`        // ContextId of API request
	SessionId        *string                   `json:"sessionId,omitempty"`        // SessionId of API request
	TessellErrorCode *string                   `json:"tessellErrorCode,omitempty"` // Unique error code specific to Tessell
	UserView         *TessellExceptionUserView `json:"userView,omitempty"`
}

type TessellExceptionUserView struct {
	Message    *string               `json:"message,omitempty"`    // End-user representation of the message
	Resolution *string               `json:"resolution,omitempty"` // End-user representation of resolution
	ErrorCode  *TessellHttpErrorCode `json:"errorCode,omitempty"`
}

type TessellHttpErrorCode struct {
	HttpCode *int    `json:"httpCode,omitempty"` // HTTP code
	Code     *string `json:"code,omitempty"`     // Tessell&#39;s specific code with more context on error
}

type DatabaseParameterProfileListResponse struct {
	Response *[]DatabaseParameterProfileResponse `json:"response,omitempty"`
	Metadata *APIMetadata                        `json:"metadata,omitempty"`
}

// Request payload for creating a Parameter Profile
type DatabaseParameterProfileRequest struct {
	Name                     *string                            `json:"name"`
	Description              *string                            `json:"description,omitempty"`
	EngineType               *string                            `json:"engineType"`
	EngineInfo               *DatabaseParameterEngineInfo       `json:"engineInfo,omitempty"`
	Parameters               *[]DatabaseProfileParameterRequest `json:"parameters"`
	DBVersion                *string                            `json:"dbVersion"`
	InfraType                *string                            `json:"infraType,omitempty"`
	SourceParameterProfileId *string                            `json:"sourceParameterProfileId,omitempty"`
}

// Parameter request for create/update
type DatabaseProfileParameterRequest struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// Request payload for updating a Parameter Profile
type DatabaseParameterProfilePatchRequest struct {
	ParameterProfileInfo *DatabaseParameterProfilePatchInfo `json:"parameterProfileInfo"`
	PropagationPolicy    *UpdatePropagationPolicy           `json:"propagationPolicy"`
	ServiceInstances     *[]ServiceInstancesPatchRequest    `json:"serviceInstances,omitempty"`
}

type DatabaseParameterProfilePatchInfo struct {
	Description *string                            `json:"description,omitempty"`
	Parameters  *[]DatabaseProfileParameterRequest `json:"parameters"`
}

type UpdatePropagationPolicy struct {
	Strategy *string `json:"strategy"` // IMMEDIATELY, MAINTENANCE_WINDOW, CUSTOM_DATE_TIME, DO_NOT_APPLY
	Time     *string `json:"time,omitempty"`
}

// ServiceInstancesPatchRequest for targeting specific service instances during update
type ServiceInstancesPatchRequest struct {
	Service   *GovernanceServiceInfo            `json:"service,omitempty"`
	Instances *[]GovernancePatchInstanceRequest `json:"instances,omitempty"`
}

type GovernanceServiceInfo struct {
	Name *string `json:"name,omitempty"`
	Id   *string `json:"id,omitempty"`
}

type GovernancePatchInstanceRequest struct {
	Name *string `json:"name,omitempty"`
	Id   *string `json:"id,omitempty"`
}

// Response from GET /governance/parameter-profiles/{id}/usages
type ParameterProfileUsageResponse struct {
	ServiceInstances *[]ServiceInstances `json:"serviceInstances,omitempty"`
}

type ServiceInstances struct {
	Service   *GovernanceServiceInfo    `json:"service,omitempty"`
	Instances *[]GovernanceInstanceInfo `json:"instances,omitempty"`
}

type GovernanceInstanceInfo struct {
	Name   *string `json:"name,omitempty"`
	Id     *string `json:"id,omitempty"`
	Role   *string `json:"role,omitempty"`
	Status *string `json:"status,omitempty"`
}
