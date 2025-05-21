FROM alpine:latest
COPY . .
EXPOSE 80
CMD ["/main"]
