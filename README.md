# BuckIT

## Tech Stack
SvelteKit: Front-end (client) \
Golang: Back-end (server) \
PostgreSQL (supabase): Database \

**Coming in the future** \
Redis: Caching, message broker \
Python: AI/ML models to personalize content for the users


## Features
- [x] Light/Dark Mode
- [ ] Registration
- [ ] Login 
- [ ] Logout
- [ ] User Profile
- [ ] Private Account   
- [ ] Follow/Unfollow Users
- [ ] Follow requests
- [ ] Friends
- [ ] Create BuckIT list (max 10 items in one list)
- [ ] User interactions to BuckIts
- [ ] Real-time chatting
- [ ] User session management
- [ ] User notifications (in real-time)
- [ ] Push notification
- [ ] Subscribe to a BuckIT list or an individual item from a list

## Requirements
```
Node: 22.4.1
Golang: 1.24.4
```

## Client
1. Step into the `client` directory
    ```sh
    cd client
    ```
2. Download dependencies
    ```sh
    npm install
    ```
3. Run the application
    ```sh
    npm run dev
    ```

## Server
1. Step into the `server` directory
    ```sh
    cd server
    ```
2. Download dependencies
    ```sh
    go mod download
    ```
3. Run the server
    ```sh
    go run .
    ```

    #### Folder Structure
    ```
    server/
        - api/
            - resource/
                - database/             ==> Database client for supabase
                - user/                 ==> User handlers, models, and anything else related to user
            - router/                   ==> Defines the routes
        - .env
        - .gitignore
        - generate-openapi-json.sh      ==> Generate openapi.json file
        - go.mod
        - go.sum
        - main.go
        - openapi.json                  ==> Generated documentation from application
    ```

## How to generate openapi.json file
1. Start the server
    ```sh
    # go to the server directory
    cd server

    # go run .
    ```
2. Open a new terminal
3. Go to the server directory
4. Run the generate script
    ```sh
    ./generate-openapi-json.sh
    ```
5. Now there should be an `openapi.json` file in `server/` if there wasn't one. Otherwise it either updated, if our api has changed since it was last generated, or remained the same.