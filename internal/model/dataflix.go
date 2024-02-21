package model

type TessellAmDataflixDTO struct {
	AvailabilityMachineId *string                      `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId      *string                      `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName           *string                      `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType            *string                      `json:"engineType,omitempty"`
	CloudAvailability     *[]CloudRegionInfo           `json:"cloudAvailability,omitempty"` // The cloud and region information where the data is available for access
	Owner                 *string                      `json:"owner,omitempty"`             // Owner of the Availability Machine
	Tsm                   *bool                        `json:"tsm,omitempty"`               // Specify whether the associated DB Service is created using TSM compute type
	SharedWith            *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}

type TessellDataflixResponse struct {
	Metadata *APIMetadata            `json:"metadata,omitempty"`
	Response *[]TessellAmDataflixDTO `json:"response,omitempty"`
}
