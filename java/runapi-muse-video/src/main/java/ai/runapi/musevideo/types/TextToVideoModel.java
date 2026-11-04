package ai.runapi.musevideo.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to video operations. */
public final class TextToVideoModel extends MusevideoValue {
  /** muse-video-v1 model slug. */
  public static final TextToVideoModel MUSE_VIDEO_V1 = new TextToVideoModel("muse-video-v1");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToVideoModel(String value) {
    super(value);
  }
}
