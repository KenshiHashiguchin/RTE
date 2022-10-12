FROM mysql:5.7

#####################################
# Set Timezone
#####################################

ARG TZ=UTC
ENV TZ ${TZ}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_AUTHENTICATION_PLUGIN=${MYSQL_AUTHENTICATION_PLUGIN}

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && chown -R mysql:root /var/lib/mysql/

COPY my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]

EXPOSE 3306
