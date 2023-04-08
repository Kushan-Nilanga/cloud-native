# react container
FROM node:18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN npm install

RUN npm run build

# nginx container

FROM nginx:1.13.12-alpine

COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]