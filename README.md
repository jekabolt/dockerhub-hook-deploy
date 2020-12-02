dockerhub webhook deploy image to remote server

ENV VARIABLES:
SERVER_PORT = default 8555
DOCKER_AUTH_TOKEN = default kektoken
REMOTE_SSH no default = root@11.111.111.11

``make run`` - to run locally 
``make dist`` - to build for linux/amd64 
``make scp`` - scp to server (only if REMOTE_SSH env exist)