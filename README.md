# Coding Challenge

This repository is my approach to a 48 hour coding challenge including Golang, Docker and Firebase authentication.

## Requirements

- Create a backend REST service that receives a list of desired micro-courses in JSON format. The payload data is not organized in any specific order.
- Develop a service that can create a study schedule that lists courses in an order that respects the constraints (see the sorting constraints below).
- Use Docker Compose to provide the solution of an SQL database.
- Develop unit and end-to-end tests.
- API endpoints should be protected with authentication using Firebase Authentication.

### Data input
**JSON payload example**:

```
{
   "userId":"30ecc27b-9df7-4dd3-b52f-d001e79bd035",
   "courses":[
      {
         "desiredCourse":"PortfolioConstruction",
         "requiredCourse":"PortfolioTheories"
      },
      {
         "desiredCourse":"InvestmentManagement",
         "requiredCourse":"Investment"
      },
      {
         "desiredCourse":"Investment",
         "requiredCourse":"Finance"
      },
      {
         "desiredCourse":"PortfolioTheories",
         "requiredCourse":"Investment"
      },
      {
         "desiredCourse":"InvestmentStyle",
         "requiredCourse":"InvestmentManagement"
      }
   ]
}
```

**The service must return the courses sorted as follows**:
0. Finance
1. Investment
2. InvestmentManagement
3. PortfolioTheories
4. InvestmentStyle
5. PortfolioConstruction

## Ecosystem in which it was tested

- Docker version `20.10.14`
- Docker Compose version `1.29.2`, build `5becea4c`
- MacBook Pro 13-inch 2020, `macOS Monterrey Version 12.3.1`
- Postman version `9.15.13`
- Firebase project

## Assumptions

A few assumptions were made when developing this challenge, specifically regarding the sorting algorithm:

- It was assumed that no course could have more than one required course/correlative, meaning that the data could be modeled using a tree data structure instead of a graph.
- It was assumed that the end user sends correct data, with correctness being tied to this list of assumptions.
- Since multiple courses could be at the same level (ie the exact same required courses/amount of required courses), we assumed that the first one to appear as a `desired course` (aka a `child node`), is returned first, since that is the same criteria followed in the example provided.
- It was assumed that, just as in the example provided, the data could always be modeled by a tree data structure. We assume that there is always a root node (ie a required course at the top that has no correlatives), and that there is only one root node.

## How to run app locally

- First and foremost, create a Firebase project, download the service account key JSON file and insert it at the project's root under the name `serviceAccountKey.json`
- Insert your project's API key in the `Dockerfile` (`FIREBASE_API_KEY`)
- Standing in the projects root folder, the following command was used: `docker-compose build && docker-compose up`

#### Step-by-step requests to try authentication as well as the sorting algorithm

- Initially, if one tries to consume `POST /courses/sort`, it will return an error since an authentication token is needed.
- If one tries to log (`POST user/login`) in with any user, it should return an error since you must sing up before logging in.
- By consuming the endpoint to sign up (`POST user/signup`) with the corresponding body (see examples in the Postman collection) containing a valid `email` (can be any random email so long as it can be parsed into an email address), a non-empty `password` and the `returnSecureToken` field set to `true`, that email will be registered and the bearer token will be returned, which is valid for an hour.
    - If one wanted to test the token creation manually using Firebase's API, there is an example in the Postman collection.
- From this point on, consuming `POST user/login` with the registered `email` and `password` will return a renewed bearer token.
- With this bearer token, now one can consume `POST /courses/sort`.

## How to run unit tests

- With Goland, you can use the GUI and right-click the project to run the tests
- From the command line, you can use `go test -v ./...` or `go test -v ./... | grep FAIL` to filter for failed tests

## How to test locally with Postman

Postman collection containing 4 requests:

- Request to sort the collection
- Request to sign up with a random email and password
- Request to log in with a previously signed-up email and password
- Request to Firebase's API to get the bearer token
    - **Note:** While with Firebase's Admin SDK for other languages one can easily call the sign-in methods to get the required token, Go's SDK doesn't provide this, so without a frontend I had to manually consume this API.

```
{
	"info": {
		"_postman_id": "6a764292-35ca-488a-afb7-dddf6c3a88db",
		"name": "golangchallenge",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Sort",
			"request": {
				"auth": {
					"type": "bearer",
					"bearer": [
						{
							"key": "token",
							"value": "BEARER_TOKEN",
							"type": "string"
						}
					]
				},
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\"userId\": \"30ecc27b-9df7-4dd3-b52f-d001e79bd035\",\"courses\": [{\n            \"desiredCourse\": \"PortfolioConstruction\",\n            \"requiredCourse\": \"PortfolioTheories\"\n        },\n        {\n            \"desiredCourse\": \"InvestmentManagement\",\n            \"requiredCourse\": \"Investment\"\n        },\n        {\n            \"desiredCourse\": \"Investment\",\n            \"requiredCourse\": \"Finance\"\n        },\n        {\n            \"desiredCourse\": \"PortfolioTheories\",\n            \"requiredCourse\": \"Investment\"\n        },\n        {\n            \"desiredCourse\": \"InvestmentStyle\",\n            \"requiredCourse\": \"InvestmentManagement\"\n        }\n    ]\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:4000/courses/sort",
					"host": [
						"localhost"
					],
					"port": "4000",
					"path": [
						"courses",
						"sort"
					]
				}
			},
			"response": []
		},
		{
			"name": "Login",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"email\": \"EMAIL\",\n    \"password\": \"PASSWORD\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:4000/user/login",
					"host": [
						"localhost"
					],
					"port": "4000",
					"path": [
						"user",
						"login"
					]
				}
			},
			"response": []
		},
		{
			"name": "Sign up",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"email\": \"EMAIL\",\n    \"password\": \"PASSWORD\",\n    \"returnSecureToken\": true\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:4000/user/signup",
					"host": [
						"localhost"
					],
					"port": "4000",
					"path": [
						"user",
						"signup"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create token",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"email\": \"EMAIL\",\n    \"password\": \"PASSWORD\",\n    \"returnSecureToken\": true\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=API_KEY",
					"protocol": "https",
					"host": [
						"www",
						"googleapis",
						"com"
					],
					"path": [
						"identitytoolkit",
						"v3",
						"relyingparty",
						"verifyPassword"
					],
					"query": [
						{
							"key": "key",
							"value": "API_KEY"
						}
					]
				}
			},
			"response": []
		}
	]
}
```

## Next steps

- The dimension that I would like to improve upon the most is testing, most of all E2E tests, which I could not get
  around to due to some Firebase authentication using solely a Golang backend taking more time than expected, as well as
  an unexpected issue with the protocol used to connect to MySQL
- The hexagonal implementation could be improved regarding, for example, how the logger is injected into the project.
  With more time, I'd implement a more robust logger layer to stand between the inbound traffic and the actual calls to
  the controllers. By intercepting evert call, we could log the requests (query parameters, body, origin, responses and
  errors) from a single place
- The formatting of the logs could be improved to simplify parsing them and processing them programmatically
- The `web pkg` could be improved into a robust response and handler to avoid duplicated code