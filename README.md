# Grader

A distributed competitive programming online judge. This online judge is designed using a microservices architecture.

## Feature
- Multiple worker-based grading system
- OAuth-based authentication
- Docker-based sandboxed code execution
- Real-time grading process via Websocket
- Multiple test cases evaluation
- Runtime and memory limit error detection

## Services
### Web-service
- Serve frontend static files
- Provide reverse-proxy

### Auth-service
- User login
- Generate JSON web token

### Problem-service
- Manage problems in database
- Keep problems meta data in `Supabase`
- Keep problem testcases and pdf files in **Storage-service**

### User-service
- Manage users data
- Serve user profile

### Storage-service
- Keep file
- Need secret key to access

### Grade-service
- In-memory worker queue using Go channels
- Goroutine-based worker pool
- Create isolated Docker container
- Compile source code
- Execute testcases
- Return grading result

### Guild-service
- Manage guild system

## Hosting
### Singleton hosting
Copy this repository to machine then run
```
docker compose up --build
```
### Distributed hosting
Host each services on specific type of hosting services
- Web-service on container service
- Auth-service on container service
- User-service on container service
- Guild-service on container service
- Storage-service on virtual machine service

## Workflow
### User open the website
1. **Web-service** serve static files.
2. Get user profile from **user-service**. If success show the profile otherwise show login button.

### Login
1. Get token from Google-Firebase.
2. Send token to **auth-service**.
3. Set JSON web token in user cookie.

### Create new problem
1. Send problem meta data to **problem-service**.
2. Generate problem id and set to cookie.
3. Send testcases and pdf files to **problem-service**.
4. Check if meta data and files problem id matched.
5. Keep meta data in `Supabase` and keep files in **storage-service**.

### Submission
1. Send problem id, language and code to **grade-service**.
2. Push submission request to job queue.
3. Worker get a job from request queue.
4. Create new container.
5. Compile code to executable.
6. Get resource limitaion from problem meta data in `Supabase`.
7. Set container resource.
8. Get testcases from **storage-service**.
9. Establish websocket connection with user.
10. Execute testcases and stream results in real time.
11. Save result to `Supabase`.

## Future plan
- Frontend
- Multi-languages support
- Guild system
- Contest system

## Tech stack
### Frontend
- React

### Backend
- Go
- Fiber

### Infrastructure
- Docker
- Nginx

### Authentication
- Firebase

### Database
- Supabase
