#!/bin/bash
DOCKER_IMAGE_RAILS=`docker images | grep my/rails | awk '{ print $7}'`
if [ `echo $DOCKER_IMAGE_RAILS | grep MB` ];
then echo ${DOCKER_IMAGE_RAILS:0:-2} >> /tmp/docker_image_rails
else echo ${DOCKER_IMAGE_RAILS:0:-2} | awk '{print $1*1000}' >> /tmp/docker_image_rails
fi

DOCKER_IMAGE_MYSQL=`docker images | grep my/mysql | awk '{ print $7}'`
if [ `echo $DOCKER_IMAGE_MYSQL | grep MB` ];
then echo ${DOCKER_IMAGE_MYSQL:0:-2} >> /tmp/docker_image_mysql
else echo ${DOCKER_IMAGE_MYSQL:0:-2} | awk '{print $1*1000}' >> /tmp/docker_image_mysql
fi
