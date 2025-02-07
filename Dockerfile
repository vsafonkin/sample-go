FROM ubuntu:latest
COPY ./bin/sample /

EXPOSE 8083

CMD [ "/sample" ]
