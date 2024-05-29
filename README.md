# API Service

This is a simple API service built with Go.

## Installation

### 1. Install Go

Ensure you have Go installed on your system. You can download and install it from the official Go website: [https://golang.org/](https://golang.org/).

### 2. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/Nagarjunhabbu/Watchy.git
```
### 3.Install Docker compose

```bash
docker-compose up
```
Server will start on local host with port 8000.
Once the service is running, you can make HTTP requests to the API. Here's an example using curl:
### 4. Example API Usage
1. Create:
API end point to insert watch history of an user

```bash
curl --request POST \
  --url http://localhost:8000/v1/event \
  --header 'Content-Type: application/json' \
  --data '{
    "user_id": "user456",
    "event_title": "IPL Playoffs",
    "video_id": "v1965343",
    "action": "SKIP",
    "duration": 120
}
'
```
Response
```json
{
	"id": 4,
	"user_id": "user456",
	"event_title": "IPL Playoffs",
	"video_id": "v1902343",
	"action": "SKIP",
	"duration": 120
}
```



2. GET Events:
API end point to get the watch history data of particular user
```bash
curl --request GET \
  --url http://localhost:8000/v1/event/user123
```
Response
```json
[
	{
		"id": 1,
		"user_id": "user123",
		"event_title": "Movie Night",
		"video_id": "video456",
		"action": "watch",
		"duration": 120
	},
	{
		"id": 2,
		"user_id": "user123",
		"event_title": "IPL Final",
		"video_id": "video89",
		"action": "watch",
		"duration": 120
	},
	{
		"id": 3,
		"user_id": "user123",
		"event_title": "IPL Playoffs",
		"video_id": "v1902343",
		"action": "SKIP",
		"duration": 120
	}
]

```
