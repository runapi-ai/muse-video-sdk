package musevideo

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTP struct {
	method, path string
	body         any
}

func (s *stubHTTP) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method, s.path = method, path
	if opts != nil {
		s.body = opts.Body
	}
	return json.RawMessage(`{"id":"task-1","status":"processing"}`), nil
}

func TestTextToVideoCreateAndGet(t *testing.T) {
	stub := &stubHTTP{}
	client := NewClientWithHTTP(stub)
	_, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{Model: ModelMuseVideoV1, Prompt: "A city at dawn"})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != textPath {
		t.Fatalf("unexpected request %s %s", stub.method, stub.path)
	}
	if got := stub.body.(map[string]any)["prompt"]; got != "A city at dawn" {
		t.Fatalf("unexpected prompt: %v", got)
	}
	_, err = client.TextToVideo.Get(context.Background(), "task-1")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != textPath+"/task-1" {
		t.Fatalf("unexpected request %s %s", stub.method, stub.path)
	}
}

func TestImageToVideoCreate(t *testing.T) {
	stub := &stubHTTP{}
	client := NewClientWithHTTP(stub)
	_, err := client.ImageToVideo.Create(context.Background(), ImageToVideoParams{Model: ModelMuseVideoV1, Prompt: "Animate this scene", ImageURL: "https://cdn.runapi.ai/public/samples/image.jpg"})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != imagePath {
		t.Fatalf("unexpected request %s %s", stub.method, stub.path)
	}
	if got := stub.body.(map[string]any)["image_url"]; got != "https://cdn.runapi.ai/public/samples/image.jpg" {
		t.Fatalf("unexpected image URL: %v", got)
	}
}
