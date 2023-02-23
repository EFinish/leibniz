FROM node:latest

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY fe/package*.json ./
RUN npm install

# Bundle app source
COPY fe/ .

# Build TypeScript
RUN npm run build