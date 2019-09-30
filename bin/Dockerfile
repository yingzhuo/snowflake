FROM busybox:1.28.4

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./snowflake-linux-amd64-$VERSION /bin/snowflake

CMD ["snowflake", "--host=0.0.0.0", "--port=8080", "--node-id=512", "--reponse-type=protobuf"]
