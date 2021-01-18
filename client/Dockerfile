FROM node:14-slim
WORKDIR /app
COPY package*.json ./
RUN npm install --silent
EXPOSE 3000
COPY . ./
CMD ["npm", "start"]