
FROM nvidia/cuda:12.2.0-base-ubuntu22.04
WORKDIR /app
COPY . .
RUN apt update -y
RUN apt upgrade  -y
RUN apt install bash git wget build-essential python3 python3-pip cmake golang -y
CMD ["/bin/bash"]
