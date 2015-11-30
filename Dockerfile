FROM centos
RUN yum install -y java-1.8.0-openjdk-devel
COPY target/service.yaml .
COPY target/deposit-auth-ws-0.0.1-SNAPSHOT.jar .
EXPOSE 8080 8081
CMD java -jar deposit-auth-ws-0.0.1-SNAPSHOT.jar server service.yaml

