"""Muse Video client."""

from __future__ import annotations

from typing import Any, Optional
from runapi.core import ProviderClient
from .resources import ImageToVideo, TextToVideo


class MuseVideoClient(ProviderClient):
    """Text-to-video and image-to-video API client."""

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        super().__init__(api_key, **options)
        self.text_to_video = TextToVideo(self._http)
        self.image_to_video = ImageToVideo(self._http)
