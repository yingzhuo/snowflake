FROM registry.cn-shanghai.aliyuncs.com/yingzhuo/alpine:3

LABEL maintainer="应卓 <yingzhor@gmail.com>"

COPY --chown=root:root ./snowflake-linux-amd64-* /bin/snowflake
COPY --chown=root:root ./docker-entrypoint.sh /bin/docker-entrypoint.sh

RUN chmod +x /bin/snowflake && \
    chmod +x /bin/docker-entrypoint.sh

ENV SNOWFLAKE_NODE_ID=0 \
    SNOWFLAKE_HTTP_PORT=8080 \
    SNOWFLAKE_INDENT=false \
    SNOWFLAKE_QUIET=false

ENTRYPOINT ["docker-entrypoint.sh"]