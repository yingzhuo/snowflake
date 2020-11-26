FROM registry.cn-shanghai.aliyuncs.com/yingzhuo/alpine:3

LABEL maintainer="应卓 <yingzhor@gmail.com>"

HEALTHCHECK NONE

COPY --chown=root:root ./snowflake-linux-amd64-* /opt/snowflake
COPY --chown=root:root ./docker-entrypoint.sh /bin/entrypoint.sh

RUN chmod +x /opt/snowflake && \
    chmod +x /bin/entrypoint.sh

ENV SNOWFLAKE_NODE_ID=0 \
    SNOWFLAKE_HTTP_PORT=8080 \
    SNOWFLAKE_INDENT=false \
    SNOWFLAKE_QUIET=false

ENTRYPOINT ["entrypoint.sh"]