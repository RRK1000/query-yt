# query-yt
A demo project that polls the Youtube API and stores video information, built with Go & Docker.

# Running and Building 
```
$ docker compose up --build -d
```


# Component Diagram
![image](https://user-images.githubusercontent.com/38955457/217796433-9bc5a9a7-c19e-48a9-a736-a542c399859a.png)



# Supported APIs
### GET api/v1/videoinfo
Fetches upto `maxResults` objects of videos for the predefined query, starting from the `offset`
([Usage](https://github.com/RRK1000/query-yt/tree/docs/src/api-svc#get-apiv1videoinfo))
| Query Params | Type | Description |
| ----------- | ----------- | ----------- |
| maxResults | integer | Number of results in the response |
| offset | integer | offset value for pagination |


### GET api/v1/video
Fetches video information based on a full `title`, part of the `description` or both
([Usage](https://github.com/RRK1000/query-yt/tree/docs/src/api-svc#get-apiv1videoinfo))
| Form Data | Type | Description |
| ----------- | ----------- | ----------- |
| title | string | Video Title |
| description | string | part of the description |



# Development Practices 
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
