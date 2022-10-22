FROM ruby:2.6.6
RUN apt-get update -qq && \
    apt-get install -y build-essential \
                       locales \
                       libmariadb-dev
RUN locale-gen ja_JP.UTF-8
ENV LANG ja_JP.UTF-8
ENV LC_CTYPE ja_JP.UTF-8
RUN localedef -f UTF-8 -i ja_JP ja_JP.utf8
RUN locale
WORKDIR /
COPY containers/ruby/Gemfile /
COPY web/api/db /db
RUN gem install bundler -v "1.17.3"
RUN bundle install