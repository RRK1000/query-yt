# query-yt

# Supported APIs
### GET api/v1/videoinfo
Fetches upto `maxResults` objects of videos for the predefined query, starting from the `offset`

| Query Params | Type | Description |
| ----------- | ----------- | ----------- |
| maxResults | integer | Number of results in the response |
| offset | integer | offset value for pagination |


### GET api/v1/video
Fetches video information based on a full `title`, part of the `description` or both

| Form Data | Type | Description |
| ----------- | ----------- | ----------- |
| title | string | Video Title |
| description | string | part of the description |

# Running and Building 
`
    docker compose up --build -d
`

# Supported APIs
GET api/v1/videoinfo

GET api/v1/video

# Development Practices 
- Conventional Commits
