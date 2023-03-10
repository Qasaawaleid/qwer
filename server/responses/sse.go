package responses

import (
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
)

type WebhookStatusUpdateOutputs struct {
	ID               uuid.UUID                      `json:"id"`
	ImageUrl         string                         `json:"image_url"`
	UpscaledImageUrl string                         `json:"upscaled_image_url,omitempty"`
	GalleryStatus    generationoutput.GalleryStatus `json:"gallery_status,omitempty"`
}

type SSEStatusUpdateResponse struct {
	Status    CogTaskStatus                `json:"status"`
	Id        string                       `json:"id"`
	StreamId  string                       `json:"stream_id"`
	Error     string                       `json:"error,omitempty"`
	NSFWCount int32                        `json:"nsfw_count,omitempty"`
	Outputs   []WebhookStatusUpdateOutputs `json:"outputs,omitempty"`
}
