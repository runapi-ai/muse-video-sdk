# RunAPI Muse Video Ruby SDK

Create asynchronous Muse Video tasks from a prompt or one source image.

```ruby
require "runapi/muse_video"

client = RunApi::MuseVideo::Client.new
task = client.text_to_video.create(
  model: "muse-video-v1",
  prompt: "A city at dawn"
)

completed = client.image_to_video.run(
  model: "muse-video-v1",
  prompt: "Animate this scene",
  image_url: "https://images.unsplash.com/photo-1470770841072-f978cf4d019e"
)
```

Use `create` and `get` for asynchronous workflows, or `run` to poll until completion. Generated file URLs are temporary; store needed results in durable storage during their retention window.

See the [Muse Video catalog](https://runapi.ai/models/muse-video) for current model details and pricing.
