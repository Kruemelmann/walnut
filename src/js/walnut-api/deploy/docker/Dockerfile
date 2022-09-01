FROM node:18-alpine As build

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./

RUN yarn
COPY . .

RUN yarn build
RUN yarn install --frozen-lockfile --production && yarn cache clean

###################

FROM node:18-alpine As production

COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/dist ./dist

CMD [ "node", "dist/main.js" ]

