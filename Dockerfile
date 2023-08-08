FROM golang:latest AS builder


# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
#RUN mkdir /user && \
#    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
#    echo 'nobody:x:65534:' > /user/group
# Copy the predefined netrc file into the location that git depends on
COPY ./.netrc /root/.netrc
RUN chmod 600 /root/.netrc

WORKDIR /usr/src/app/
ENV GOPRIVATE=github.com/bentkowskiw/*

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify
COPY . .
# RUN rm -rf cfg
RUN mkdir -p /usr/local/bin/app
RUN go test ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/local/bin/app ./...

# Final stage: the running container.
#FROM scratch AS final
# Import the user and group files from the first stage.
#COPY --from=builder /user/group /user/passwd /etc/
# Import the compiled executable from the first stage.
#COPY --from=builder /usr/local/bin/app/calendar-api /usr/local/bin/app/calendar-api
# Perform any further action as an unprivileged user.
#USER nobody:nobody
# Run the compiled binary.
ENTRYPOINT [ "/usr/local/bin/app/oauth" ] 