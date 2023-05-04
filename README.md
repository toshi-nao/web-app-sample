# web-app-sample
## Database: MongoDB

```sh
cd monbodb
```
```sh
docker build . -t mongo-test
````
Run MongoDB container
```sh
docker run -p 28001:27017 --name mongo-test -d mongo-test
```
Install mongosh
```sh
brew install mongosh
```

Test connection to MongoDB with mongosh
```sh
mongosh mongodb://root:'brHZ-!_rHAZF4xR2-EsRKx9e'@localhost:28001/admin
```
```mongo
admin> exit
```

## `/go-gin/`: Back-end sample by Go with Gin Framework

Install go
```sh
brew install go
```

```sh
cd go-gin
```

Run go
```sh
go run main.go
```

## `/vue-js/`: Front-end sample by Vue.js


Install npm
```sh
brew install npm
```
Install node
```sh
brew install node
```
Install vue-cli globally
```sh
npm install -g @vue/cli
```

```sh
cd vue-js
```

Install npm packages to `node_modules`
```sh
npm install
```

Compile and hot-reload for development
```sh
npm run serve
```


