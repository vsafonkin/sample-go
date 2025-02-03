FROM ubuntu:latest
COPY ./bin/sample /

EXPOSE 81

CMD [ "/sample" ]
