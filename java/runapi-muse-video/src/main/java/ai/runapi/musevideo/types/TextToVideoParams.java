package ai.runapi.musevideo.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for text to video operations. */
public final class TextToVideoParams {
  private final String model;
  private final String prompt;
  private final String aspectRatio;
  private final String callbackUrl;

  private TextToVideoParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.aspectRatio = builder.aspectRatio;
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new TextToVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "muse-video/text-to-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", MusevideoParamUtils.wireValue(model));
    raw.put("prompt", MusevideoParamUtils.wireValue(prompt));
    raw.put("aspect_ratio", MusevideoParamUtils.wireValue(aspectRatio));
    raw.put("callback_url", MusevideoParamUtils.wireValue(callbackUrl));
    return MusevideoParamUtils.compact(raw);
  }



  /** Builder for {@link TextToVideoParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private String aspectRatio;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(TextToVideoModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = value;
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = value;
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = value;
      return this;
    }

    /** Builds immutable text to video parameters. */
    public TextToVideoParams build() {
      return new TextToVideoParams(this);
    }
  }
}
