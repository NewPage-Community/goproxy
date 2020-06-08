FROM scratch
ADD build /app
CMD ["/app/goproxy"]