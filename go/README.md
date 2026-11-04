# RunAPI Muse Video Go SDK

Create asynchronous Muse Video tasks from a prompt or one source image.

```go
import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/runapi-ai/core-sdk/go/option"
	"github.com/runapi-ai/muse-video-sdk/go/musevideo"
)

client, err := musevideo.NewClient(option.WithAPIKey(os.Getenv("RUNAPI_API_KEY")))
if err != nil {
	log.Fatal(err)
}

task, err := client.TextToVideo.Create(context.Background(), musevideo.TextToVideoParams{
	Model:  musevideo.ModelMuseVideoV1,
	Prompt: "A city at dawn",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(task.ID)
```

Use `Create` and `Get` for asynchronous workflows, or `Run` to poll until completion. Generated file URLs are temporary; store needed results in durable storage during their retention window.

See the [Muse Video catalog](https://runapi.ai/models/muse-video) for current model details and pricing.
