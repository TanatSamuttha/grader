FROM node:20 AS builder

WORKDIR /app

COPY ./react_backup/package*.json ./
RUN npm install

COPY ./react_backup ./

RUN npm run build


FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]