FROM node:latest

WORKDIR /usr/src/app

COPY package.json /usr/src/app/
COPY package-lock.json /usr/src/app/

RUN npm install -g yarn 

RUN yarn install

COPY . .

EXPOSE 3000

CMD [ "yarn", "start" ]