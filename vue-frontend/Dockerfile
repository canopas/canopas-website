FROM node:20-alpine AS ui-build
WORKDIR /app
COPY . ./
RUN yarn install --frozen-lockfile && yarn build

FROM node:20-alpine AS server-build
WORKDIR /root/
COPY --from=ui-build /app/dist ./dist
COPY --from=ui-build /app/node_modules ./node_modules
COPY --from=ui-build /app/utils ./utils
COPY --from=ui-build /app/server.js /app/package.json /app/yarn.lock ./

EXPOSE 3080

CMD ["yarn", "serve"]
