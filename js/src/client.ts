import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { TextToVideo } from './resources/text-to-video';
import { ImageToVideo } from './resources/image-to-video';

/** Muse Video text-to-video and image-to-video client. */
export class MuseVideoClient extends BaseClient {
  public readonly textToVideo: TextToVideo;
  public readonly imageToVideo: ImageToVideo;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.textToVideo = new TextToVideo(this.http);
    this.imageToVideo = new ImageToVideo(this.http);
  }
}
