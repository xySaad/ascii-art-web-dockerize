## Stage 1 build
FROM golang:1.22.3 AS build

# Set working directory to /src
WORKDIR /src

# Copy nessecary files for build process
COPY app ./

# Set environemt variable for static go compling
ENV CGO_ENABLED=0

# Build the golang application into /app
RUN go build -o /app/ascii-art-web


## Stage 2
FROM alpine:3.20.3


# Set working directory to /app
WORKDIR /app

# Copy artifacts from stage 1
COPY --from=build /app/ .

# Copy assets from host
COPY assets static ./

# Expose port 8080
EXPOSE 8080


## Start Command
CMD [ "/app/ascii-art-web" ]




## Metadata
LABEL \
  z01.ascii-art-web.version="1.0.0" \
  z01.ascii-art-web.description="A Docker image for the ASCII Art Web application" \
  z01.ascii-art-web.maintainer="" \
  z01.ascii-art-web.team.members="srm, aayyada, zdiouri" \
  z01.ascii-art-web.build-date="2024-10-11" \
  z01.ascii-art-web.git-commit="latest" \
  z01.ascii-art-web.homepage="http://t5t.us.to/"