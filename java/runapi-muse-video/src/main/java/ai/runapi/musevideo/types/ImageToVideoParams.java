package ai.runapi.musevideo.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for image to video operations. */
public final class ImageToVideoParams {
  private final String model;
  private final String prompt;
  private final String imageUrl;
  private final String aspectRatio;
  private final String callbackUrl;

  private ImageToVideoParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.imageUrl = builder.imageUrl;
    this.aspectRatio = builder.aspectRatio;
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new ImageToVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "muse-video/image-to-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", MusevideoParamUtils.wireValue(model));
    raw.put("prompt", MusevideoParamUtils.wireValue(prompt));
    raw.put("image_url", MusevideoParamUtils.wireValue(imageUrl));
    raw.put("aspect_ratio", MusevideoParamUtils.wireValue(aspectRatio));
    raw.put("callback_url", MusevideoParamUtils.wireValue(callbackUrl));
    return MusevideoParamUtils.compact(raw);
  }



  /** Builder for {@link ImageToVideoParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private String imageUrl;
    private String aspectRatio;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(ImageToVideoModel value) {
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

    /** Sets the image URL. */
    public Builder imageUrl(String value) {
      this.imageUrl = value;
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

    /** Builds immutable image to video parameters. */
    public ImageToVideoParams build() {
      return new ImageToVideoParams(this);
    }
  }
}
