import type { HttpClient, PollingOptions, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type { CompletedMuseVideoResponse, ImageToVideoParams, MuseVideoResponse, TaskCreateResponse } from '../types';

const ENDPOINT = '/api/v1/muse_video/image_to_video';

export class ImageToVideo {
  constructor(private readonly http: HttpClient) {}

  async create(params: ImageToVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request('POST', ENDPOINT, { body: compactParams(params), ...options });
  }

  async get(id: string, options?: RequestOptions): Promise<MuseVideoResponse> {
    return this.http.request('GET', `${ENDPOINT}/${id}`, options ?? {});
  }

  async run(params: ImageToVideoParams, options?: RequestOptions & PollingOptions): Promise<CompletedMuseVideoResponse> {
    const { id } = await this.create(params, options);
    return await pollUntilComplete<MuseVideoResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    }) as CompletedMuseVideoResponse;
  }
}
