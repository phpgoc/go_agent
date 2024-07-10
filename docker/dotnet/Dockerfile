FROM docker.io/debian:12
RUN apt update
RUN apt install -y  wget

RUN wget https://dot.net/v1/dotnet-install.sh
RUN chmod +x dotnet-install.sh
RUN ./dotnet-install.sh --channel 8.0
ENV PATH=$PATH:/root/.dotnet

RUN apt install -y libicu-dev
