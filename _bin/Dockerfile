FROM busybox:1.31.1

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./snowflake-linux-amd64-$VERSION /bin/snowflake

ENV SNOWFLAKE_VERSION=$VERSION \
    SNOWFLAKE_PORT=8080 \
    SNOWFLAKE_NODE_ID=0 \
    SNOWFLAKE_TYPE=protobuf \
    SNOWFLAKE_INDENT=false \
    SNOWFLAKE_QUIET=false \
    GIN_MODE=release

CMD ["snowflake"]