FROM maven:3.8.2-jdk-11-slim as plantuml

WORKDIR /app

COPY pom.xml .

RUN mvn dependency:copy-dependencies

FROM openjdk:11-slim

ENV PERSIST_DEPS \
    graphviz \
    fonts-dejavu

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
    ${PERSIST_DEPS} \
    && apt-get clean \
    && apt-get purge -y \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=plantuml /app/target/dependency/plantuml-*.jar /app/plantuml.jar

WORKDIR /data

ENTRYPOINT [ "java", "-Djava.awt.headless=true", "-jar", "/app/plantuml.jar" ]

CMD [ "-help" ]
