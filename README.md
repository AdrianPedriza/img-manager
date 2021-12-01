# Images manager
This repository contains the necessary code to deploy a server which exposes a rest api that handles two types of image operations: **loading a new image and getting an image given its name**.
## Project structure
```
├── cmd
│   └── server
│       └── main.go       // To run our rest server
├── pkg
│   └── model
│       └── image.go      // Model for our application
│   └── rest
│       └── handler.go    // Request handlers
│   └── storage
│       └── device.go     // Implementation when work on device storage
│       └── handler.go    // Implementation of an interface to work with different storage systems
│       └── memory.go     // Implementation when work on memory storage
```

## API
The api exposes two endpoints with which you can interact with it.

| Endpoint | HTTP Method | Possible http status codes | Note |
| ------------- | ------------- | --- | --- |
| `/images/<img-id>`  | `GET`  | `200`, `400`, `404`, `500`|img-id = image name (**extension included**). Ex: my_image.png
| `/images`  | `POST`  | `201`, `400`, `409`, `500` | Set Header **Content-Type: multipart/form-data**|

## How to consume API
The following are examples of how to consume the api using cURL.
### Get image by ID
- **cURL**
```cmd
curl --request GET 'http://localhost:8080/images/person.jpeg' --output <full-new-file-name>
```

### Upload new image
- **cURL**
```cmd
curl --request POST 'http://localhost:8080/images' \
--form 'image=@"<img-path>"'
```

## How to run the application locally
The application can be deployed in two ways: locally or using docker. But before deploying the application, it is necessary to configure the following environment variables:

| Environment variable | Possible value| Default value | Description |
| ------------- | ------------- | ------------- | ------------- |
| IMG_STORAGE_DEVICE  | memory / device  | device |type of storage used by the app |
| IMG_STORAGE_DIR  | path to folder  | generated in app |directory where the images are stored when the storage is 'device'|

### Locally
To execute the application locally, it is enough to generate the binary that contains our application. To do so, inside the cmd/server directory of our project, execute::
```cmd
go build -o cmd/server ./...
```
Once the binary has been generated and environment variables has been set, we can deploy our server locally by running:
```cmd
./cmd/server/server
```

### Using docker
Another alternative to deploy the server is through docker. For this it is necessary to have the tool installed. Instructions for this can be found here. Once docker is installed, follow the next steps:
### Build docker image (inside repository folder)
```cmd
docker build -t <image-name> .
```
### Run the docker image
```cmd
docker run -p 8080:8080 --env-file=./.env <image-name>
```
**NOTE:** Fill `.env` file (attached `.env.sample`) with the desired values.

## Observations
**NOTE:** Tests not implemented due to the limited business logic of the application itself.

