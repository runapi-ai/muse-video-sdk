# frozen_string_literal: true

module RunApi
  module MuseVideo
    module Resources
      class TextToVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/muse_video/text_to_video"

        def initialize(http)
          @http = http
        end

        def create(options: nil, **params)
          request(:post, ENDPOINT, body: compact_params(params), options: options)
        end

        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end

        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end
      end
    end
  end
end
