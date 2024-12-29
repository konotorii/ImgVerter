# ImgVerter
- - -
Go based REST API to upload and host images & files, utilizing an in-memory state management to deliver images quickly. Images are converted to WebP either on upload or while fetching, ensuring images are optimized to smaller sizes and delievered faster.

## Important Note
Currently all files uploaded through ImgVerter get assigned a unique UUID so be sure to save this ID to fetch the file. Any pre-existing files or uploaded differentl, for example SFTP, can be called with the original filename inside any nested folder.

## Env Variables
The following is just filled with default values, these can be customized to your own extent.
```dotenv
SECRET="" // Secret for cookie sessions
PORT=3000 // Port server runs on

REDIS_HOST="" // Redis address and port: localhost:6379
REDIS_PASS="" // Password for redis
REDIS_DB=0 // 0-16 are the available databases in Redis

DOMAIN="localhost" // This is very important to set for uploading files, this will be a part of the return link to access the file through ImgVerter

UPLOAD_KEY="" // Key for uploading images & files
PUBLIC_FOLDER="/public/" // If ImgVerter is run inside Docker then this should be an absolute path

// Upload Settings
UPLOAD_ALLOWED_FILE_TYPES=png,jpg,gif // Comma-seperated file types
UPLOAD_MAX_FILE_SIZE=5 // Values in MB only
```

## Available Routes
|Route|Method|Description|
|---|---|---|
|/i/:file| GET | Fetch files from folder ``/public/i/``
|/*/**| GET | Wildcard fetch for all files in ``/public/*/**``


## Roadmap
- Video uploader with conversion to m3u8
- Login session handler
- Serve image from URL as a optimization proxy
- Ability to create users or some sort of way to limit admin access
- Perceptual Hashing to determine if an image was already uploaded or not
- More configs
