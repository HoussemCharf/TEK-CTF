FROM ubuntu:16.04
RUN apt-get update && apt-get -y install python socat vim
COPY task7 .
COPY flag .
RUN cp /usr/bin/vi /usr/bin/fc
EXPOSE 8009
CMD socat TCP-LISTEN:8009,fork,reuseaddr EXEC:"./task7",pty,stderr