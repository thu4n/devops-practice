FROM node:18.20.4 AS builder
WORKDIR /app
COPY package.json .
RUN npm install --slient
COPY . .
RUN npm run build

FROM nginx:1.27.1-alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]