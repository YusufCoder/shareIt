FROM whatwedo/golang

#------------------  BASE SETTINGS -------------------#

#Maintainer
MAINTAINER Felix Imobersteg <felix.imobersteg@me.com>

#Update
RUN apt-get update


#----------------------- SHAREIT ---------------------#

#Compile
RUN apt-get install make
ADD . /usr/src
WORKDIR /usr/src
RUN make deps
WORKDIR /usr/src/app
RUN go build
RUN mv app shareit
RUN cp shareit /root
RUN cp -R static /root
WORKDIR /root
RUN rm -rf /usr/src/*

#Create data directory
RUN mkdir /data

#Alter upstart script
RUN echo -n './shareit /data $ADMIN_PW' >> /bin/upstart


#---------------- CONTAINER SETTINGS -----------------#

#Slimming down Docker containers
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

#Create volumes
VOLUME  ["/data"]

#Set upstart script
CMD /bin/upstart

#Expose Ports
EXPOSE 9000
