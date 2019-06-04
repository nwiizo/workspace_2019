FROM ruby:2.5

ENV LANG C.UTF-8
ENV WORKSPACE=/usr/local/src

# install bundler.
RUN apt-get update 
RUN apt-get install -y vim less 
RUN apt-get install -y build-essential libpq-dev 
RUN gem install bundler 

# add Developer Tool will delete it later
RUN apt-get install -y vim zsh htop tcpdump unzip traceroute curl man strace iftop 


# create user and group.
RUN groupadd -r --gid 1000 rails && \
    useradd -m -r --uid 1000 --gid 1000 rails

RUN apt-get update && apt-get install wget -y && wget -qO- https://get.docker.com | sh

# create directory.
RUN mkdir -p $WORKSPACE $BUNDLE_APP_CONFIG
RUN chown -R rails:rails $WORKSPACE
RUN chown -R rails:rails $BUNDLE_APP_CONFIG

USER rails
WORKDIR $WORKSPACE

# bundle install.
COPY --chown=rails:rails Gemfile $WORKSPACE/Gemfile
RUN bundle install && \
    bundle exec rails new . --force --database=mysql && \
    { \
        echo "gem 'therubyracer', platforms: :ruby"; \
    } >> $WORKSPACE/Gemfile && \
    bundle update
COPY --chown=rails:rails database.yml $WORKSPACE/config/database.yml

EXPOSE  3000
CMD ["rails", "server", "-b", "0.0.0.0"]
