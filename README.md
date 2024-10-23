# Cyberpunk-hacking-minigame
## Description

This project is a hacking game inspired by the mini-game from Cyberpunk 2077. Users interact with a two-dimensional matrix and attempt to hack it using specified sequences of coordinates. The project consists of two microservices developed with Go, REST API, and MongoDB.

## Game Rules: Successful Hack and Matrix Navigation

These rules define the navigation process and the conditions for a successful hack, ensuring a strategic approach to the gameplay.

#### 1.  Successful Hack:

- A hack is considered successful if, by following a given path through the matrix cells, the resulting sequence of values matches one of the predefined sequences.

#### 2. Coordinate Change Rules:

- Start of the Path (Step 1): At the first step, the coordinates start with row = 0.
- Step 2: At the next step, the column coordinate remains the same as in the previous step.
- Step 3: Repeat the rule from step 2, but this time for the row coordinate.
- Step 4 and beyond: Continue alternating the coordinate change as in steps 2 and 3.

#### 3. Restriction on Cell Usage:

- During the entire path, it is prohibited to use the same cell more than once.

## Project Structure

- **hacker-service**: The service that generates the matrix and sequences for hacking.
- **martix-service**: The service that receives hacking attempts and checks their validity.

## Installation and Running

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/matrix-hacking-game.git
2. Build and run the services using Docker Compose:
   ```bash
   docker-compose up --build
3. The services will be available at the following addresses:
- martrix-service:  http://localhost:8080
- hacker-service:   http://localhost:8081

## Usage

### Matrix Generation and Sequences

To generate a new matrix and sequences, send a `GET` request to the `/GetSequence` endpoint with the desired ID:
```bash 
curl -X GET "http://localhost:8080/GetSequence?id=1" 
``` 
### Hacking Attempt
To attempt to solve the matrix manually, send a `POST` request to the `/Hack` endpoint with the matrix ID and a string representing the path of coordinates:
```bash
curl -X POST "http://localhost:8080/Hack?matrix_id=1&attempts="00 10 11 21""
```
### Automatic Hacking Attempt

If the matrix seems too challenging to solve manually, or if you want to view the solution for a specific matrix, you can use the second service for automatic solving. Send a `POST` request to the `/Hack` endpoint of the second service:

```bash
curl -X POST "http://localhost:8081/Hack?matrix_id=3"
```
This request will use the given matrix_id to retrieve the solution, if available, for the specified matrix. The service will try to solve the matrix using the provided ID and return the solution path if successful.

### API Documentation
For convenient management of methods, we recommend visiting the Swagger interface after starting the services:
- Matrix Service: [Swagger Documentation](http://localhost:8080/swagger/index.html#/)
- Hacker Service: [Swagger Documentation](http://localhost:8081/swagger/index.html#/)

---

_Project created by Igor Balashko as a technical assignment on 23.10.2024._

