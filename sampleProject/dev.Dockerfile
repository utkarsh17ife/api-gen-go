FROM node:11.7

RUN mkdir /app
WORKDIR /app

COPY ./package.json ./
RUN npm install

EXPOSE 3000
ENTRYPOINT ["npm", "run"]

