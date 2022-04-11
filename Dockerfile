FROM node:17 AS ui-build
WORKDIR /app
COPY vue-frontend/ ./
RUN npm install && npm run build

FROM node:17 AS server-build
WORKDIR /root/
COPY --from=ui-build /app/dist ./dist
COPY --from=ui-build /app/node_modules ./node_modules
COPY --from=ui-build /app/server.js /app/package*.json ./

EXPOSE 3080

CMD ["npm", "run", "serve"]
