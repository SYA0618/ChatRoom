FROM nginx:1.27.0
WORKDIR /front
COPY ./front-end ./
COPY ./conf.d/default.conf /etc/nginx/conf.d/default.conf
RUN chown nginx:nginx ./*
EXPOSE 80