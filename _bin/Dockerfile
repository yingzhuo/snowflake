FROM busybox:1.32.0

LABEL maintainer="应卓 <yingzhor@gmail.com>"

COPY ./snowflake-linux-amd64-* /bin/snowflake

ENV SNOWFLAKE_NODE_ID=0 \
    SNOWFLAKE_HTTP_PORT=8080 \
    SNOWFLAKE_INDENT=false \
    SNOWFLAKE_QUIET=false

EXPOSE 8080

ENTRYPOINT ["snowflake"]