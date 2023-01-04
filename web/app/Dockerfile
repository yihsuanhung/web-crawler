# A quick implementation, not optimize for production !!!
# 快速實現版本，尚未進行優化

FROM node:16.14.2
WORKDIR /app
COPY package.json .
RUN yarn
COPY . .
EXPOSE 5173
CMD ["yarn", "dev"]
