version: '3.8'
services:
  mock-api-server:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:{{ . }}"
    command: ["./main"]