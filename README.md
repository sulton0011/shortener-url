# TASK 
**URL Shortener Service**

## Project Description:
    You should build a URL shortener service backend that allows users to shorten long URLs into shorter ones. The short URLs should redirect to the original long URLs when visited

## The service should have the following functionalities:
    - A user should be able to register and login to the service
    - A user should submit a long URL to the service and should get short URL
    - The service generates a unique short URL for the long URL and stores it in a database
    - The service redirects users who access the short URL to the original long URL
    - The service should support custom short URLs
    - The service should support expiration of short URLs after a certain amount of time or a certain number of clicks
    - The service should track the number of clicks for each short URL and display it to the user
    - User should be able to see list of URLs, and should be able to modify them
    - The short URLs generated should be as short as possible, while still being unique
    - Frequently accessed short URLs should be cached to improve performance
    - The user should be able to generate a QR code for a short URL. This feature is optional but it will be plus

## Technical Requirements:
    - Use Golang as the programming language
    - Use PostgreSQL database to store all data
    - Use Gin package to handle routing
    - Use Redis for caching
    - Use a hashing algorithm (e.g. SHA-256, MD5) to generate the short URLs
    - Write unit tests for the components, minimum unit test coverage should be at least 70%

## You should provide:
    - Link to github repository
    - A README file that includes instructions on how to run the program and any additional information about the service
    - Dockerfile and Docker Compose file for running backend
    - Swagger for API documentation



# Project structure: