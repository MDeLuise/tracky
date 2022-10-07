FROM ubuntu

RUN apt-get update
RUN apt install software-properties-common wget -y

RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get update
RUN apt-get install golang -y

RUN wget https://github.com/gobuffalo/cli/releases/download/v0.18.8/buffalo_0.18.8_Linux_x86_64.tar.gz
RUN tar -xvzf buffalo_0.18.8_Linux_x86_64.tar.gz
RUN mv buffalo /usr/local/bin/buffalo
RUN go install github.com/gobuffalo/buffalo-pop/v3@latest

RUN mkdir -p /src/dev
WORKDIR /src/dev

ADD . .
RUN sed -i database.yml -e "s/\(^  host\: \).*$/\1db/"
RUN buffalo build --static -o /bin/app

ENV ADDR=0.0.0.0
EXPOSE 3000

#CMD exec /bin/app
ENTRYPOINT [ "./deployment/run-server.sh" ]