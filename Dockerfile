FROM node:20 as node-builder

WORKDIR /app

# copy all filtes
COPY assets/src/ .

RUN rm -rf dist
RUN rm -rf node_modules

# install vite globally
RUN npm install -g vite

# install all deps
RUN yarn install

RUN yarn vite build --outDir dist

FROM golang:latest AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go get -d -v ./...

COPY cmd cmd
COPY pkg pkg
COPY assets/web.go assets/web.go

#Copy the build from the node-builder
COPY --from=node-builder /app/dist /app/assets/dist

RUN ls -lha *

RUN go build -o main cmd/main.go

FROM ubuntu:latest

WORKDIR /app

COPY --from=go-builder /app/main /app/main
ENV MONGO_URI=mongodb://localhost:27017

EXPOSE 8082

CMD ["/app/main"]

