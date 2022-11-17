FROM golang:1.18-alpine
RUN mkdir /study-event-go-asynq
WORKDIR /study-event-go-asynq
ADD bin/study-event-go-asynq bin/study-event-go-asynq
ADD conf conf
ARG BUILD_PORT
ENV PORT $BUILD_PORT
EXPOSE $BUILD_PORT
ENTRYPOINT ["bin/study-event-go-asynq"]