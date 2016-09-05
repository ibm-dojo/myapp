FROM ubuntu
ADD myapp /
EXPOSE 80
ENTRYPOINT /myapp
