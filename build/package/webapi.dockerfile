FROM jwhitaker/recipebook-base:latest

WORKDIR /app
ADD . /app

# RUN make webapi

RUN echo ${service}

CMD ["/app/bin/webapi"]