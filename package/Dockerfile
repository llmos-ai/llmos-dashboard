FROM opensuse/leap:15.6

WORKDIR /var/lib/llmos-dashboard

# Add dumb-init
ENV INIT_VERSION 1.2.5
RUN ARCH=$(uname -m); \
    curl -sfL https://github.com/Yelp/dumb-init/releases/download/v${INIT_VERSION}/dumb-init_${INIT_VERSION}_${ARCH} -o dumb-init \
    && chmod +x dumb-init \
    && mv dumb-init /usr/bin/dumb-init

COPY ./ui/build ./ui/build
COPY static ./static
COPY data ./data
COPY package/entrypoint.sh llmos-dashboard /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh

VOLUME /var/lib/llmos-dashboard
ENTRYPOINT ["entrypoint.sh"]
