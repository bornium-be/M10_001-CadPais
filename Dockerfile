FROM ubuntu:18.04

RUN mkdir -p /app/bornium

COPY teste /app/bornium/

CMD ["/app/bornium/"]

EXPOSE 5001

ENTRYPOINT ["./app/bornium/M10_001-CadPais"]