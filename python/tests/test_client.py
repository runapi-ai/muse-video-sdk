from runapi.muse_video.resources.image_to_video import ImageToVideo
from runapi.muse_video.resources.text_to_video import TextToVideo


class StubHttp:
    def __init__(self):
        self.calls = []

    def request(self, method, path, **kwargs):
        self.calls.append((method, path, kwargs))
        return {"id": "task-1", "status": "processing"}


def test_text_to_video_create_and_get():
    http = StubHttp()
    resource = TextToVideo(http)
    resource.create(model="muse-video-v1", prompt="A city at dawn")
    assert http.calls[-1] == (
        "post",
        "/api/v1/muse_video/text_to_video",
        {"body": {"model": "muse-video-v1", "prompt": "A city at dawn"}, "options": None},
    )
    resource.get("task-1")
    assert http.calls[-1][1] == "/api/v1/muse_video/text_to_video/task-1"


def test_image_to_video_posts_source_image():
    http = StubHttp()
    resource = ImageToVideo(http)
    resource.create(model="muse-video-v1", prompt="Animate", image_url="https://cdn.runapi.ai/public/samples/image.jpg")
    assert http.calls[-1][1] == "/api/v1/muse_video/image_to_video"
    assert http.calls[-1][2]["body"]["image_url"].endswith("image.jpg")
