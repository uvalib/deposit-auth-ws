# define as appropriate if you do not want the defaults
# export DATABASE_URL=
# export DATABASE_USER=
# export DATABASE_PASSWORD=

java -jar target/deposit-auth-ws-0.0.1-SNAPSHOT.jar db migrate src/main/resources/service.yaml
