FROM node:12.19.0-alpine3.12

ENV HOST 0.0.0.0

ARG UID=1000
ARG GID=1000
ARG USERNAME=rte
ARG GROUPNAME=rte

ARG APP_NAME
ENV APP_NAME=${APP_NAME}
ARG APP_ENV
ENV APP_ENV=${APP_ENV}
ARG APP_DEBUG
ENV APP_DEBUG=${APP_DEBUG}
ARG APP_URL
ENV APP_URL=${APP_URL}

RUN apk add --no-cache curl

RUN ( addgroup -g ${GID} ${GROUPNAME} || true ) \
    && ( adduser -D -u ${UID} -G ${GROUPNAME} ${USERNAME} || true )

WORKDIR /var/www/rte/web/
#ADD --chown=${UID}:${GID} ./application/frontend ./
ADD ./web/frontend ./
ADD ./containers/nuxt/startup.sh /tmp/startup.sh

RUN chown -R ${UID}:${GID} /var/www/

USER ${UID}:${GID}

#RUN yarn run build

EXPOSE ${PORT}
#CMD ["/bin/sh", "/tmp/startup.sh"]
