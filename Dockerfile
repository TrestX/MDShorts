FROM golang
WORKDIR /opt/MDShorts
COPY . .
ARG go_env
ENV GO_ENV=$go_env
RUN mv /opt/MDShorts/conf-override/application-qa.yml /opt/MDShorts/conf-override/application.yml
RUN go build
EXPOSE 6019
CMD go run main.go