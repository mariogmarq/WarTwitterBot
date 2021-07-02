FROM golang:latest

WORKDIR /usr/app/
COPY . .

ENV DB_URL /usr/app/the_database.db
#RUN touch ${DB_URL} //Create the DB in each execution, just for testing porposes

ENV IMAGES_DIR /usr/app/assets/images

RUN go test ./...
RUN rm ${DB_URL}
RUN go build -o civilbot

CMD [ "./civilbot" ]
