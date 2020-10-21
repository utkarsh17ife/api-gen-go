package template

func GetDockerFile() (string, string, string) {
	return `dev.Dockerfile`, `.`, `FROM node:11.7

RUN mkdir /app
WORKDIR /app

COPY ./package.json ./
RUN npm install

EXPOSE 3000
ENTRYPOINT ["npm", "run"]

`
}

func GetDockerCompose() (string, string, string) {
  return `docker-compose.yml`, `.`, `version: "3.5"
  
services:
  api:
    build:
      context: ./
      dockerfile: dev.Dockerfile
    command: dev
    ports:
      - "3000:3000"
    depends_on:
      - ganache
      - api-watch
      - mongo
    volumes:
      - ./build:/app/build:delegated
      - ./package.json:/app/package.json:delegated
    environment:
      WEB3_HTTP_PROVIDER: "http://ganache:8545"
      OFFCHAIN_DB_NAME: "ureakdb"
      OFFCHAIN_DB_URL: "mongodb://mongo:27017"
      PORT: 3000
    networks:
      - ureka

  api-watch:
    build:
      context: ./
      dockerfile: dev.Dockerfile
    command: watch
    volumes:
      - ./src:/app/src:delegated
      - ./build:/app/build:delegated
      - ./.babelrc:/app/.babelrc
    depends_on:
      - ganache
    logging:
      options:
        max-size: 10m

  mongo:
    image: mongo:latest
    networks:
      - ureka
    logging:
      options:
        max-size: 10m

  ganache:
    image: trufflesuite/ganache-cli:latest
    networks:
      - ureka
    command:
      ganache-cli -m 'candy maple cake sugar pudding cream honey rich smooth crumble sweet treat'
      --networkId '333'

networks:
  ureka:
    driver: bridge


`

}
