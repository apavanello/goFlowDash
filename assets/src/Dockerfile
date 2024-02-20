FROM node:20
WORKDIR /app
# install vite globally
RUN npm install -g vite
# copy all filtes
COPY . .
# install all deps
RUN yarn install

# vite default port
EXPOSE 8080
CMD ["yarn", "run", "vite", "--host", "0.0.0.0", "--port", "8080"]