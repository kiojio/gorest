# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Matrix  <matrix.kiojio@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk upgrade && \
	apk add --no-cache bash git vim

# ADD CURL
RUN apk add --update tzdata \
	bash wget curl git && \
	rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .   

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
# CMD if [ ${APP_ENV} = production ]; \
# 	then \
# 	api; \
# 	else \
# 	go get github.com/pilu/fresh && \
# 	fresh; \
# 	fi

CMD go get github.com/pilu/fresh && \
	  fresh; 

EXPOSE 9070