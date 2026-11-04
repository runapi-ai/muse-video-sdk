# frozen_string_literal: true

module RunApi
  module MuseVideo
    class Client < RunApi::Core::Client
      attr_reader :text_to_video, :image_to_video

      def initialize(api_key: nil, **options)
        super
        @text_to_video = Resources::TextToVideo.new(http)
        @image_to_video = Resources::ImageToVideo.new(http)
      end
    end
  end
end
