FROM debian:stable-slim
# --platform=linux/amd64 

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely

CMD ["notely"]
