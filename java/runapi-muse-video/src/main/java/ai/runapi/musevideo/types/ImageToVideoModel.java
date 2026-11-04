package ai.runapi.musevideo.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for image to video operations. */
public final class ImageToVideoModel extends MusevideoValue {
  /** muse-video-v1 model slug. */
  public static final ImageToVideoModel MUSE_VIDEO_V1 = new ImageToVideoModel("muse-video-v1");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public ImageToVideoModel(String value) {
    super(value);
  }
}
