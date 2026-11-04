# frozen_string_literal: true

Dir.chdir(__dir__) do

  Gem::Specification.new do |spec|
    spec.name = "runapi-muse-video"
    spec.version = "0.2.0"
    spec.metadata["runapi_slug"] = "muse-video"
    spec.authors = ["RunAPI"]
    spec.email = ["contact@runapi.ai"]
    spec.summary = "RunAPI Muse Video Ruby SDK"
    spec.description = "Create asynchronous Muse Video tasks from a prompt or one source image."
    spec.homepage = "https://runapi.ai/models/muse-video"
    spec.license = "Apache-2.0"
    spec.required_ruby_version = ">= 3.1.0"
    spec.metadata["homepage_uri"] = "https://runapi.ai/models/muse-video"
    spec.metadata["documentation_uri"] = "https://github.com/runapi-ai/muse-video-sdk/blob/main/ruby/README.md"
    spec.metadata["source_code_uri"] = "https://github.com/runapi-ai/muse-video-sdk"
    spec.metadata["bug_tracker_uri"] = "https://github.com/runapi-ai/muse-video-sdk/issues"
    spec.metadata["changelog_uri"] = "https://github.com/runapi-ai/muse-video-sdk/blob/main/CHANGELOG.md"

    spec.files = Dir.glob("lib/**/*") + %w[LICENSE README.md]
    spec.extra_rdoc_files = ["README.md"]
        spec.require_paths = ["lib"]
    spec.add_dependency "runapi-core", "~> 0.6.1"
  end
end
