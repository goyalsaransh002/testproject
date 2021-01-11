FROM ubuntu
USER root
ARG TARGETARCH=amd64
ENV PORT 8080
EXPOSE 8080
COPY hosts /etc/hosts
COPY policy-engine /
ENTRYPOINT ["/policy-engine"]
