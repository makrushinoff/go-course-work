FROM golang:1.22.2-bullseye
WORKDIR /application
COPY . .
RUN go mod download
RUN go build -o /godocker
 
EXPOSE 8080
 
CMD [ “/godocker” ]
