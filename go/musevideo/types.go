package musevideo

import "github.com/runapi-ai/core-sdk/go/core"

// Model selects the Muse Video model.
type Model string

const ModelMuseVideoV1 Model = "muse-video-v1"

// TaskResponse contains the task status and completed media.
type TaskResponse struct {
	Usage  *core.TaskUsage `json:"usage,omitempty"`
	ID     string          `json:"id"`
	Status string          `json:"status"`
	Videos []Video         `json:"videos,omitempty"`
	Error  string          `json:"error,omitempty"`
}

func (r TaskResponse) GetID() string     { return r.ID }
func (r TaskResponse) GetStatus() string { return r.Status }
func (r TaskResponse) GetError() string  { return r.Error }

// Video contains a generated video URL.
type Video struct {
	URL string `json:"url"`
}

// TextToVideoParams describes a prompt-only video request.
type TextToVideoParams struct {
	Model       Model  `json:"model" help:"required; Muse Video model slug"`
	Prompt      string `json:"prompt" help:"required; video description, up to 20000 characters"`
	AspectRatio string `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio: 16:9, 9:16, 1:1, 4:3, or 3:4"`
	CallbackURL string `json:"callback_url,omitempty" help:"optional; URL for terminal task notifications"`
}

// ImageToVideoParams describes a source-image video request.
type ImageToVideoParams struct {
	Model       Model  `json:"model" help:"required; Muse Video model slug"`
	Prompt      string `json:"prompt" help:"required; video description, up to 20000 characters"`
	ImageURL    string `json:"image_url" help:"required; HTTP(S) URL of one source image"`
	AspectRatio string `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio: 16:9, 9:16, 1:1, 4:3, or 3:4"`
	CallbackURL string `json:"callback_url,omitempty" help:"optional; URL for terminal task notifications"`
}
