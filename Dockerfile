FROM alpine:3
ADD auth /bin/auth
EXPOSE 8000
ENV LOG_LEVEL=-1
CMD ["auth", "server"]