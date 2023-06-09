
# syntax=docker/dockerfile:1.2

# Stage 1: Build the static files
FROM node:16.15.0-alpine3.15 as frontend-builder
WORKDIR /builder
COPY /frontend/package.json /frontend/package-lock.json ./
RUN npm ci
COPY /frontend .
RUN npm run build

# Stage 2: Build the binary
FROM golang:1.18.3-alpine3.15 as binary-builder
ARG APP_NAME=http
RUN apk update && apk upgrade && \
  apk --update add git
WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /builder/build ./frontend/build/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-w -s -extldflags "-static"' -a \
  -o engine cmd/${APP_NAME}/main.go

# Stage 3: Run the binary
FROM gcr.io/distroless/static
ENV APP_PORT=5050
WORKDIR /app
COPY --from=binary-builder --chown=nonroot:nonroot /builder/engine .
EXPOSE $APP_PORT
ENTRYPOINT ["./engine"]
