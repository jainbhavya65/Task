FROM alpine:3.9 as base
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

FROM scratch
WORKDIR /app
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /usr/share/zoneinfo/ /usr/share/zoneinfo/
COPY main .
CMD [ "./main" ]