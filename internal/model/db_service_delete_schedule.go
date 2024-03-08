package model

type DeletionSchedulePayload struct {
	DeleteAt       *string                       `json:"deleteAt"` // Time at which the DB Service should be deleted at
	DeletionConfig *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
}
