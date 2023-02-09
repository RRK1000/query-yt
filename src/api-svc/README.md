# api-svc

1. Paginated Video Info
2. Search Video

## GET api/v1/videoinfo
#### Example Request
```
curl --location --request GET 'http://localhost:8000/api/v1/videoinfo?maxResults=1&offset=3'
```

#### Example Response
```
[
    {
        "Title": "This New Simulation Football Game Looks GREAT!",
        "Description": "A new simulation football game is releasing soon, and it's looking really good! Maximum Football 2023 has a lot of people's ...",
        "PublishingDatetime": "2023-02-09T01:56:46Z",
        "Thumbnail": {
            "default": {
                "height": 90,
                "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/default.jpg",
                "width": 120
            },
            "high": {
                "height": 360,
                "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/hqdefault.jpg",
                "width": 480
            },
            "medium": {
                "height": 180,
                "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/mqdefault.jpg",
                "width": 320
            }
        }
    }
]
```

## GET api/v1/video
#### Example Request
```
curl --location --request GET 'http://localhost:8000/api/v1/video' \
--form 'title="This New Simulation Football Game Looks GREAT!"' \
--form 'description="Maximum Football 2023"'
```

#### Example Response
```
{
    "Title": "This New Simulation Football Game Looks GREAT!",
    "Description": "A new simulation football game is releasing soon, and it's looking really good! Maximum Football 2023 has a lot of people's ...",
    "PublishingDatetime": "2023-02-09T01:56:46Z",
    "Thumbnail": {
        "default": {
            "height": 90,
            "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/default.jpg",
            "width": 120
        },
        "high": {
            "height": 360,
            "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/hqdefault.jpg",
            "width": 480
        },
        "medium": {
            "height": 180,
            "url": "https://i.ytimg.com/vi/Xq9BE5JHZXU/mqdefault.jpg",
            "width": 320
        }
    }
}
```
