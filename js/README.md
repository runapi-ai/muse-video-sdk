# Muse Video JavaScript SDK for RunAPI

Use `@runapi.ai/muse-video` to create typed Muse Video tasks from a prompt or a single image.

```bash
npm install @runapi.ai/muse-video
```

```typescript
import { MuseVideoClient } from '@runapi.ai/muse-video';

const client = new MuseVideoClient();
const task = await client.textToVideo.create({ model: 'muse-video-v1', prompt: 'A city at dawn' });
const result = await client.textToVideo.get(task.id);
```

Use `create()` and `get()` for asynchronous workflows, or `run()` to poll until completion.

- [Model reference](https://runapi.ai/models/muse-video)
- [API reference](https://runapi.ai/docs/api/muse-video/text-to-video)
- [Pricing and rate limits](https://runapi.ai/models/muse-video)

Licensed under the Apache License, Version 2.0.
