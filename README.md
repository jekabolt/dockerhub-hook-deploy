Dockerhub webhook deploy image to remote server 
========

Goals
-----

* Receive webhook from docker-hub and trigger ``docker pull`` and ``docker run``   

Environment variables
--------------
SERVER_PORT = default 8555  
DOCKER_AUTH_TOKEN = default kektoken  
REMOTE_SSH no default = root@11.111.111.11  


Using dockerhub-hook-deploy
--------------

``make run`` - to run locally   
``make dist`` - to build for linux/amd64    
``make scp`` - scp to server (only if REMOTE_SSH env exist)   