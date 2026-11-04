# RunAPI Muse Video JavaScript SDK

Create asynchronous Muse Video tasks from a prompt or one source image.

<p align="center"><a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a></p>

<p align="center"><a href="https://github.com/runapi-ai/muse-video-sdk">Muse Video API SDK for RunAPI</a></p>

Muse Video API SDKs for JavaScript, Python, Ruby, Go, and Java on RunAPI.

[![npm](https://img.shields.io/npm/v/@runapi.ai/muse-video)](https://www.npmjs.com/package/@runapi.ai/muse-video)
[![PyPI](https://img.shields.io/pypi/v/runapi-muse-video)](https://pypi.org/project/runapi-muse-video/)
[![RubyGems](https://img.shields.io/gem/v/runapi-muse-video)](https://rubygems.org/gems/runapi-muse-video)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/muse-video-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/muse-video-sdk/go)
[![License](https://img.shields.io/github/license/runapi-ai/muse-video-sdk)](https://github.com/runapi-ai/muse-video-sdk/blob/main/LICENSE)

```typescript
import { MuseVideoClient } from "@runapi.ai/muse-video";

const client = new MuseVideoClient();
const task = await client.textToVideo.create({
  model: "muse-video-v1",
  prompt: "A city at dawn",
});

const completed = await client.imageToVideo.run({
  model: "muse-video-v1",
  prompt: "Animate this scene",
  image_url: "https://images.unsplash.com/photo-1470770841072-f978cf4d019e",
});
```

Use `create()` and `get()` for asynchronous workflows, or `run()` to poll until completion. Generated file URLs are temporary; store needed results in durable storage during their retention window.

See the [Muse Video catalog](https://runapi.ai/models/muse-video) for current model details and pricing.
