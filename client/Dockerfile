FROM node:14-slim as build
WORKDIR /app
COPY package*.json ./
RUN npm install --silent
EXPOSE 3000
COPY . ./
RUN npm run build

# production environment
FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]