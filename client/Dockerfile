FROM node:14-slim as build
WORKDIR /app
COPY package*.json ./
RUN npm install --silent
EXPOSE 3000
COPY . ./
RUN npm run build

# production environment
FROM nginx:latest
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]