FROM nginx:1.13

COPY fcgiapi entrypoint.sh /
COPY nginx-fcgiapi.conf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/nginx.conf

#CMD ["nginx", "-g", "daemon off;"]
ENTRYPOINT /entrypoint.sh
