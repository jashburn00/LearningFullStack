# Docker Notes

## Contents
- [Docker Notes](#docker-notes)
  - [Contents](#contents)
  - [Overview](#overview)
    - [getting started](#getting-started)
  - [dockerfiles](#dockerfiles)

## Overview
- Docker allows us to create, run, and orchestrate containerized environments
- Dockerfiles are files describing an environment, and they turn into Docker Images
- Docker Images are the files that Docker executes to spawn an instance of the described environment
- The concrete instances of the environment is called a Container
- This, along with other services like Docker Hub, Docker Compose, Kubernetes, etc. allow for quickly and easily running different services that will behave properly inside their containers on any machine.
### getting started
- start a container using an image:
> docker run --name newContainerName -e ENV_ATTRIBUTE=value -d -p 5431:5432 MyImage

  - the `-d` flag indicates run in detached mode
  - the `--name` flag is for naming the new container instance
  - the `-e` flag is for passing in environment variables (key-value pairs)
  - the `-p` flag is for passing in port specification in the format ToPort:FromPort
    - the "to port" is where you can access the container locally
    - the "from port" is the port specified in the docker image
- to now interact with the running container using a TTY, use the command:
> docker exec -it MyContainer

## dockerfiles
- the fuc
