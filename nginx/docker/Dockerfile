FROM nginx:alpine as nginx

WORKDIR /etc/nginx

COPY index.html /app/html
COPY nginx.conf /etc/nginx/nginx.conf
COPY ssl /etc/nginx/ssl

EXPOSE 80
EXPOSE 8080
ENTRYPOINT [ "nginx" ]
CMD [ "-g", "daemon off;" ]
