import { describe, it, expect, vi } from 'vitest';
import { TextToVideo } from '../../src/resources/text-to-video';
import { ImageToVideo } from '../../src/resources/image-to-video';
import type { HttpClient } from '@runapi.ai/core';

describe('Muse Video resources', () => {
  const http: HttpClient = { request: vi.fn().mockResolvedValue({ id: 'task-1', status: 'processing' }) };

  it('posts text-to-video inputs and queries its task path', async () => {
    const resource = new TextToVideo(http);
    await resource.create({ model: 'muse-video-v1', prompt: 'A city at dawn', aspect_ratio: '16:9' });
    expect(http.request).toHaveBeenLastCalledWith('POST', '/api/v1/muse_video/text_to_video', {
      body: { model: 'muse-video-v1', prompt: 'A city at dawn', aspect_ratio: '16:9' },
    });
    await resource.get('task-1');
    expect(http.request).toHaveBeenLastCalledWith('GET', '/api/v1/muse_video/text_to_video/task-1', {});
  });

  it('posts an image source to image-to-video', async () => {
    const resource = new ImageToVideo(http);
    await resource.create({ model: 'muse-video-v1', prompt: 'Animate the scene', image_url: 'https://cdn.runapi.ai/public/samples/image.jpg' });
    expect(http.request).toHaveBeenLastCalledWith('POST', '/api/v1/muse_video/image_to_video', {
      body: { model: 'muse-video-v1', prompt: 'Animate the scene', image_url: 'https://cdn.runapi.ai/public/samples/image.jpg' },
    });
  });
});
