package model

type DatabaseProfileParameterType struct {
	DataType      *string `json:"dataType"`
	DefaultValue  *string `json:"defaultValue"`
	ApplyType     *string `json:"applyType,omitempty"`
	Name          *string `json:"name"`
	Value         *string `json:"value"`
	AllowedValues *string `json:"allowedValues,omitempty"`
	IsModified    *bool   `json:"isModified,omitempty"`
	IsFormulaType *bool   `json:"isFormulaType,omitempty"`
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
	FactoryParameterId *string                             `json:"factoryParameterId,omitempty"` // Tessell parameter type UUID for the entity
	Status             *string                             `json:"status,omitempty"`
	MaturityStatus     *string                             `json:"maturityStatus,omitempty"`
	Owner              *string                             `json:"owner,omitempty"`
	TenantId           *string                             `json:"tenantId,omitempty"`
	LoggedInUserRole   *string                             `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the db profile
	Parameters         *[]DatabaseProfileParameterType     `json:"parameters,omitempty"`       // Parameter Profile&#39;s associated parameters
	Metadata           *DatabaseParameterProfileMetadata   `json:"metadata,omitempty"`
	DriverInfo         *DatabaseParameterProfileDriverInfo `json:"driverInfo,omitempty"`
	UserId             *string                             `json:"userId,omitempty"` // Database Parameter Profile&#39;s user id
	SharedWith         *EntityAclSharingInfo               `json:"sharedWith,omitempty"`
	DBVersion          *string                             `json:"dbVersion,omitempty"`    // Database Parameter Profile&#39;s version
	DateCreated        *string                             `json:"dateCreated,omitempty"`  // Timestamp when the entity was created
	DateModified       *string                             `json:"dateModified,omitempty"` // Timestamp when the entity was last modified
}

type DatabaseParameterProfileMetadata struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type DatabaseParameterProfileDriverInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type APIErrorOps struct {
	Code       *string                   `json:"code,omitempty"`    // Status code for the error response
	Message    *string                   `json:"message,omitempty"` // Error message for API response
	Resolution *string                   `json:"resolution,omitempty"`
	Timestamp  *string                   `json:"timestamp,omitempty"`
	ContextId  *string                   `json:"contextId,omitempty"` // ContextId of API request
	SessionId  *string                   `json:"sessionId,omitempty"` // SessionId of API request
	UserView   *TessellExceptionUserView `json:"userView,omitempty"`
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
