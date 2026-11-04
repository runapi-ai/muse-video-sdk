// Package musevideo provides Muse Video text-to-video and image-to-video operations.
package musevideo

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const textPath = "/api/v1/muse_video/text_to_video"
const imagePath = "/api/v1/muse_video/image_to_video"

// Client provides Muse Video generation resources.
type Client struct {
	base.Base
	TextToVideo  *TextToVideo
	ImageToVideo *ImageToVideo
}

func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{Base: base.New(httpClient), TextToVideo: &TextToVideo{httpClient}, ImageToVideo: &ImageToVideo{httpClient}}
}

type TextToVideo struct{ http core.HTTPClient }
type ImageToVideo struct{ http core.HTTPClient }

func (r *TextToVideo) Create(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textPath, core.CompactParams(params), requestOptions)
}

func (r *TextToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TaskResponse](ctx, r.http, core.ResourcePath(textPath, id), requestOptions)
}

func (r *TextToVideo) Run(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*TaskResponse, error) {
	_, polling := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*TaskResponse, error) { return r.Get(ctx, id, opts...) }, polling)
}

func (r *ImageToVideo) Create(ctx context.Context, params ImageToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, imagePath, core.CompactParams(params), requestOptions)
}

func (r *ImageToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TaskResponse](ctx, r.http, core.ResourcePath(imagePath, id), requestOptions)
}

func (r *ImageToVideo) Run(ctx context.Context, params ImageToVideoParams, opts ...option.RequestOption) (*TaskResponse, error) {
	_, polling := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*TaskResponse, error) { return r.Get(ctx, id, opts...) }, polling)
}
