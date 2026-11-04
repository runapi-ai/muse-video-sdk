# RunAPI Muse Video Java SDK

The Muse Video Java SDK exposes text-to-video and single-image-to-video tasks.

```kotlin
implementation("ai.runapi:runapi-muse-video:0.2.0")
```

```java
import ai.runapi.musevideo.MuseVideoClient;
import ai.runapi.musevideo.types.ImageToVideoParams;
import ai.runapi.musevideo.types.MuseVideoModel;
import ai.runapi.musevideo.types.TextToVideoParams;

try (MuseVideoClient client = MuseVideoClient.builder().apiKey(System.getenv("RUNAPI_API_KEY")).build()) {
  var created = client.textToVideo().create(TextToVideoParams.builder()
      .model(MuseVideoModel.MUSE_VIDEO_V1)
      .prompt("A city at dawn")
      .build());

  var completed = client.imageToVideo().run(ImageToVideoParams.builder()
      .model(MuseVideoModel.MUSE_VIDEO_V1)
      .prompt("Animate this scene")
      .imageUrl("https://images.unsplash.com/photo-1470770841072-f978cf4d019e")
      .build());
}
```

Use `create()` and `get()` for asynchronous workflows, or `run()` to poll until completion. Generated file URLs are temporary; store needed results in durable storage during their retention window.

See the [Muse Video catalog](https://runapi.ai/models/muse-video) for current pricing and model details.
