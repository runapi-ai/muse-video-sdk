plugins {
  `java-library`
  `maven-publish`
}

extra["runapiSlug"] = "muse-video"

description = "RunAPI Muse Video Java SDK for Muse Video workflows."

java {
  withSourcesJar()
  withJavadocJar()
}

dependencies {
  api("ai.runapi:runapi-core:0.12.0")

  testImplementation(platform("org.junit:junit-bom:5.10.3"))
  testImplementation("org.junit.jupiter:junit-jupiter")
}

publishing {
  publications {
    create<MavenPublication>("mavenJava") {
      from(components["java"])
      artifactId = "runapi-muse-video"
      pom {
        name = "RunAPI Muse Video Java SDK"
        description = "RunAPI Muse Video Java SDK for Muse Video workflows."
        url = "https://runapi.ai/models/muse-video"
        licenses {
          license {
            name = "Apache License, Version 2.0"
            url = "https://www.apache.org/licenses/LICENSE-2.0"
          }
        }
        developers {
          developer {
            id = "runapi"
            name = "RunAPI"
            email = "contact@runapi.ai"
          }
        }
        scm {
          url = "https://github.com/runapi-ai/muse-video-sdk"
          connection = "scm:git:https://github.com/runapi-ai/muse-video-sdk.git"
          developerConnection = "scm:git:ssh://git@github.com/runapi-ai/muse-video-sdk.git"
        }
      }
    }
  }
}
