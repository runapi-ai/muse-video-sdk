import type { AsyncTaskStatus, TaskResponse } from '@runapi.ai/core';

export type MuseVideoModel = 'muse-video-v1';
export type AspectRatio = '16:9' | '9:16' | '1:1' | '4:3' | '3:4';

export interface TextToVideoParams {
  model: MuseVideoModel;
  prompt: string;
  aspect_ratio?: AspectRatio;
  callback_url?: string;
}

export interface ImageToVideoParams extends TextToVideoParams {
  image_url: string;
}

export interface TaskCreateResponse { id: string; status?: AsyncTaskStatus; }
export interface MuseVideoResponse extends TaskResponse {
  id: string;
  status: AsyncTaskStatus;
  videos?: Array<{ url: string; [key: string]: unknown }>;
  error?: string;
}
export type CompletedMuseVideoResponse = MuseVideoResponse & { status: 'completed'; videos: Array<{ url: string; [key: string]: unknown }> };
