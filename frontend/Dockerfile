FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

FROM httpd:alpine as production-stage
COPY httpd.conf /usr/local/apache2/conf/
COPY --from=build-stage /app/dist /usr/local/apache2/htdocs/
EXPOSE 80