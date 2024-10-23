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
- **martix-service**: The service that receives hacking paths and checks their validity.

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
curl -X POST "http://localhost:8080/Hack?matrix_id=1&path="00 10 11 21""
```
### Additional Methods for User Convenience
To verify the proper functioning of the hacker-service and to improve the readability of matrices for subsequent solutions in the matrix-service, the following methods have been added:
#### 1. GetSequenceSugar 
This endpoint replicates the functionality of `/GetSequence` but is specifically designed for end-users. It presents matrix data in a more readable format, making it easier to understand and solve.
    
```bash
curl -X GET "http://localhost:8080/GetSequenceSugar?id=1"
```
This endpoint replicates the functionality of `/GetSequence` but is specifically designed for end-users. It presents matrix data in a more readable format, making it easier to understand and solve.
    
#### 2. HelpHack 
This endpoint allows users to verify whether a specific matrix is solvable. If a solution exists, it provides a potential path to solve the matrix. It's particularly useful for users who want to see a sample solution.
```bash
curl -X POST "http://localhost:8081/HelpHack?matrix_id=3"
```

#### 3. GetReports
This endpoint offers compiled reports on the matrices that have been solved. It provides an overview of the completed matrices, which can be useful for tracking progress or understanding the types of challenges encountered.
```bash
curl -X GET "http://localhost:8081/GetReports"
```


### API Documentation
For convenient management of methods, we recommend visiting the Swagger interface after starting the services:
- Matrix Service: [Swagger Documentation](http://localhost:8080/swagger/index.html#/)
- Hacker Service: [Swagger Documentation](http://localhost:8081/swagger/index.html#/)

---

_Project created by Igor Balashko as a technical assignment on 23.10.2024._

