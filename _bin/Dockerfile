FROM alpine:3.12.0

LABEL maintainer="应卓 <yingzhor@gmail.com>"

COPY --chown=root:root ./snowflake-linux-amd64-* /bin/snowflake
COPY --chown=root:root ./docker-entrypoint.sh /bin/docker-entrypoint.sh
COPY --from=yingzhuo/gosu:1.12 /bin/gosu /bin/gosu

RUN addgroup -g 1001 -S snowflake && \
    adduser  -u 1001 -S snowflake -G snowflake --no-create-home && \
    chmod +x /bin/snowflake && \
    chmod +x /bin/gosu && \
    chmod +x /bin/docker-entrypoint.sh

ENV SNOWFLAKE_NODE_ID=0 \
    SNOWFLAKE_HTTP_PORT=8080 \
    SNOWFLAKE_INDENT=false \
    SNOWFLAKE_QUIET=false

ENTRYPOINT ["docker-entrypoint.sh"]