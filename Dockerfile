##
## We build everything here
##
FROM golang:1.14.4-alpine3.12 as build


##
## Add git, ca-certificates and timezone info
##
RUN apk add --no-cache git ca-certificates tzdata && update-ca-certificates



##
## Add a new user here since we can't add it in scratch
##
ENV USER=user \
    UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

##
## Build the go binary here. CGO_ENABLED=0 to disable clib requirement for image to work in scratch
##
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
ARG VERSION=dev
WORKDIR /go/src/github.com/darkcrux/webhook-manager
COPY . .
RUN go build -o /go/bin/webhook-manager -mod vendor -ldflags "-X main.version=${VERSION} -w -s"

##
## Final image uses scratch. We copy zoneinfo, ca-certs, user/group details, and the binary from the previous step
##
FROM scratch
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /go/bin/webhook-manager /webhook-manager

##
## This image contains the migration files and spec
##
ADD db /db
ADD openapi.yaml /openapi.yaml

##
## Set to the unprivileged user
##
USER user:user

##
## Set the binary as the entrypoint
##
ENTRYPOINT [ "/webhook-manager" ]
CMD [ "serve" ]
