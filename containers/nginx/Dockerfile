FROM nginx:1.19.3-alpine

ARG PHP_UPSTREAM_CONTAINER
ARG UID=1000
ARG GID=1000
ARG USERNAME=rte
ARG GROUPNAME=rte

COPY ./containers/nginx/nginx.conf /etc/nginx/

RUN apk update \
    && apk upgrade \
    && apk --update add logrotate \
    && apk add --no-cache openssl \
    && apk add --no-cache bash

RUN apk add --no-cache curl

# Create 'messages' file used from 'logrotate'
RUN touch /var/log/messages

# Copy 'logrotate' config file
COPY ./containers/nginx/logrotate/nginx /etc/logrotate.d/

# Set upstream conf and remove the default conf
RUN rm /etc/nginx/conf.d/default.conf

#RUN echo "mindori:$(openssl passwd -apr1 7e7QOQg?hAXS9]50)" >> /etc/nginx/.htpasswd

ADD --chown=nginx:nginx ./containers/nginx/sites/ /etc/nginx/sites-available
ADD --chown=nginx:nginx  ./containers/nginx/ssl/ /etc/nginx/ssl
ADD --chown=nginx:nginx  ./containers/nginx/startup.sh /opt/startup.sh

RUN sed -i 's/\r//g' /opt/startup.sh
CMD ["/bin/bash", "/opt/startup.sh"]

RUN ( addgroup -g ${GID} ${GROUPNAME} || true ) \
    && ( adduser -D -u ${UID} -G ${GROUPNAME} ${USERNAME} || true )

RUN touch /run/nginx.pid && \
  chown -R nginx:nginx /run/nginx.pid

USER nginx

EXPOSE 8080 8443

WORKDIR /var/www/rte/web
