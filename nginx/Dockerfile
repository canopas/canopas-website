FROM nginx:alpine

RUN mkdir -p /etc/nginx/conf.d/

COPY nginx.conf /etc/nginx/

COPY conf.d/* /etc/nginx/conf.d/

EXPOSE 80