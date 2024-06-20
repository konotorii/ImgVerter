# ImgVerter
- - -
Go based REST API to upload and host images, utilizing an in-memory state management to deliver images quickly.


## Env Variables
```dotenv
SECRET="" // Secret for cookie sessions

REDIS_ADD="" // Redis address and port: localhost:6379
REDIS_PAS="" // Password for redis

KEY="" // Key for uploading screenshots
```

## Roadmap
- Image uploader and handler (20/6/24 Completed)
- Video uploader with conversion to m3u8
- Login session handler
- Ability to convert images to whatever is needed
- Some in-memory server like redis (10/6/24 Completed)
- Ability to create users or some sort of way to limit admin access