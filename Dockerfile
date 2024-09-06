#FROM alpine:3.19.0

FROM ubuntu:24.04

RUN mkdir -p /app/bornium

COPY .env /app/bornium/
COPY M10_001-CadPais /app/bornium/

CMD ["/app/bornium/"]

EXPOSE 5001

ENTRYPOINT ["./app/bornium/M10_001-CadPais"]