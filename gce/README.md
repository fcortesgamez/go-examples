# Google Container Engine GoLang basic service example.

This example shows how to build and deploy a containerized Go web server
application using [Kubernetes](https://kubernetes.io).

This examples has been taken/adapted from the Google Cloud Platform
Container Engine tutorial.

Visit https://cloud.google.com/container-engine/docs/tutorials/hello-app
to follow the tutorial and deploy this application on [Google Container
Engine](https://cloud.google.com/container-engine), taking into account
that this project has been renamed from hello-app to gce-example.

This directory contains:

- `main.go` contains the HTTP server implementation. It responds to all HTTP
  requests with a  "Hello, world!" response.
- `Dockerfile` is used to build the Docker image for the application.
