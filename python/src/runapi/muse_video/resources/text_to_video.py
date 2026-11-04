"""Text-to-video resource."""

from __future__ import annotations
from typing import Any, Optional
from runapi.core import Resource, RequestOptions


class TextToVideo(Resource):
    """Generate a video from a text prompt."""

    ENDPOINT = "/api/v1/muse_video/text_to_video"

    def create(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        return self._request("post", self.ENDPOINT, body=self._compact_params(params), options=options)

    def get(self, id: str, options: Optional[RequestOptions] = None) -> Any:
        return self._request("get", f"{self.ENDPOINT}/{id}", options=options)

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        task = self.create(options=options, **params)
        return self._poll_until_complete(lambda: self.get(task.id, options=options))
