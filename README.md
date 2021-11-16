# GWI - Api

## Overview
Implements a simple API which a user can use in order to fetch data for some core features (Insights, Charts Audience information) as also as subscribe and configure a variety of dashboard-assets. 

## How to build locally the database
Navigate to project's root folder and execute the commands below:
```
make compose-up
make db-up
```

Now, the container that serves the `gwi` db is up, the schema is created and the objects are populated accordingly.

## How to run locally the API server
Navigate to project's root folder and execute the command below:
```
make server-serve
```

Now, the server is up and listening to port `8080`.

### Testing
While the API is up and running execute the command below:
```
make api-test
```

## Using the API
`Endpoint `                                     | `Description`
------------                                    | -------------
(POST)  /healt                                  | Checks API's health
(POST)  /auth/signup                            | User signup
(POST)  /auth/login                             | User login - provides a bearer token
(GET)   /dashboard/listassets                   | Lists all available assets
(POST)  /dashboard/userassets                   | Gets all user's assets
(PATCH) /dashboard/updateassetdescription       | Updates asset's description
(POST)  /dashboard/subscription                 | User subscribes / unsubscribes to an asset
(POST)  /insights/insights                      | Gets insights
(POST)  /charts/chartvisits                     | Gets chart with platform visits
(POST)  /charts/chartaudiencereach              | Gets chart with info about audience search trends
(POST)  /audience/audiencesocialmedia           | Gets info about audience behavior regarding social media
(POST)  /audience/audienceshopping              | Gets info about audience behavior regarding shopping

Also, a [Postman API collection json file](./gwi.postman_collection.json) is provided in project's root folder.

## Dependencies
- Masterminds Squirrel - SQL generator for Go
- Gorilla Mux - HTTP Router
- Dgrijalva JWT - JSON Web Token Library
- Go-MySQL-Driver - A MySQL-Driver for Go's database/sql package
- Golang Crypto - A repository with cryptography libraries

## Feature improvements 
tl;dr

### Storage
Considering that a Dashboard feature is provided, a cache service, in combination with the concurrency logic that is already implemented, could be applied in order to achieve faster retrieve time for multiple Assets.

### Documentation
It would be a good idea a Swagger documentation to be provided in order to explain the API endpoints extensively.

### Testing
More test cases for API integration testing as also as a unit testing package may be really helpful. 

### Dockerfile
A Dockerfile could speed up the build process.
