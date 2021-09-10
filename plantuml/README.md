# plantuml

Docker image containing the [PlantUML](https://github.com/plantuml/plantuml)

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/hadenlabs/plantuml?sort=semver)](https://hub.docker.com/r/hadenlabs/plantuml)

## Usage

You can either mount the diagrams as volume:

```bash
docker run --rm -v $(pwd):/data hadenlabs/plantuml my-diagram.puml
```
