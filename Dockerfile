FROM scratch

COPY . $GOPATH/src/github.com/UltimateImage/VanGogh
WORKDIR $GOPATH/src/github.com/UltimateImage/VanGogh

EXPOSE 8080
CMD ["./VanGogh"]