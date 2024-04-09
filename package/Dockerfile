FROM alpine:3.19

RUN apk -U --no-cache add dumb-init musl bash build-base && \
    mkdir -p /var/lib/llmos-dashboard/

WORKDIR /var/lib/llmos-dashboard

COPY ./ui/build ./ui/build
COPY static ./static
COPY package/entrypoint.sh llmos-dashboard /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh

VOLUME /var/lib/llmos-dashboard
ENTRYPOINT ["entrypoint.sh"]
