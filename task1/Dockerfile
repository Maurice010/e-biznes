FROM ubuntu:24.04

RUN apt-get update && \
    apt-get install -y libssl-dev openssl make gcc wget tar zlib1g-dev

RUN cd /opt && wget https://www.python.org/ftp/python/3.10.16/Python-3.10.16.tgz && \
    tar xzvf Python-3.10.16.tgz && cd Python-3.10.16 && \
    cd /opt/Python-3.10.16 && ./configure && make && make install && \
    cd .. && rm -rf Python-3.10.16 Python-3.10.16.tgz && \
    python3 --version

RUN apt-get install -y openjdk-8-jdk && \
    java -version

RUN apt-get install -y curl zip unzip && \
    curl -s https://get.sdkman.io | bash && \
    bash -c "source /root/.sdkman/bin/sdkman-init.sh && sdk install kotlin"

RUN wget https://services.gradle.org/distributions/gradle-8.13-bin.zip && \
    unzip -d /opt/gradle gradle-8.13-bin.zip && \
    rm gradle-8.13-bin.zip

ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV GRADLE_HOME=/opt/gradle/gradle-8.13
ENV PATH="$JAVA_HOME/bin:$GRADLE_HOME/bin:/root/.sdkman/candidates/kotlin/current/bin:$PATH"

WORKDIR /app
COPY build.gradle .
COPY src src
RUN gradle build

# CMD ["bash", "-c", "kotlin -version && python3 --version && java -version && gradle --version"]
CMD ["gradle", "run"]