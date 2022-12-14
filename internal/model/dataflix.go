package model

type TessellAmDataflixDTO struct {
	AvailabilityMachineId *string                      `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                      `json:"tessellServiceId,omitempty"`
	ServiceName           *string                      `json:"serviceName,omitempty"`
	EngineType            *string                      `json:"engineType,omitempty"`
	CloudAvailability     *[]CloudRegionInfo           `json:"cloudAvailability,omitempty"`
	Owner                 *string                      `json:"owner,omitempty"` // Availability Machine&#39;s owner
	SharedWith            *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}

type TessellDataflixResponse struct {
	Metadata *ApiMetadata            `json:"metadata,omitempty"`
	Response *[]TessellAmDataflixDTO `json:"response,omitempty"`
}
