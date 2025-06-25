
# InnoSync API Documentation

## SwaggerUI
Access http://localhost:8080/swagger-ui/index.html for more information about API

---
## 🧑‍💼 User API (not complete)
### `GET api/user/hello`
**Description:**  
Returns a simple greeting message to confirm the API is working and authorized.

**Responses:**
- `200 OK`: Returns a welcome message.
- `401 Unauthorized`: If the user is not authenticated.
---
## 🔐 Authentication API
### `POST api/auth/signup`
**Description:**  
Registers a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```
**Responses:**
- `200 Created`: User successfully registered.
- `400 Bad Request`: Invalid input or email already taken.
---
### `POST api/auth/login`
**Description:**  
Authenticates a user and returns access and refresh tokens.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```
**Responses:**
- `200 OK`: Returns JWT access and refresh tokens.
- `401 Unauthorized`: Invalid credentials.
---
### `POST api/auth/refresh`
**Description:**  
Refreshes the access token using a valid refresh token.

**Request Body:**
```json
{
  "refreshToken": "your_refresh_token"
}
```
**Responses:**
- `200 OK`: Returns new access token.
- `401 Unauthorized`: Invalid or expired refresh token.
---
### `POST api/auth/logout`
**Description:**  
Logs the user out by invalidating the refresh token.

**Request Body:**
```json
{
  "refreshToken": "your_refresh_token"
}
```
**Responses:**
- `200 OK`: Successfully logged out.
- `401 Unauthorized`: If token is missing or already invalidated.
---
## 👤 Profile API
### `POST api/profile`
**Description:**  
Creates a new user profile.

**Request Body Example:**
```json
{
"telegram":  "@dev_user",
"github":  "dev-user",
"bio":  "Backend developer specializing in Spring Boot.",
"position":  "Software Engineer",
"education":  "BACHELOR",
"expertise":  "Java, Microservices",
"expertise_level":  "MID",
"resume":  "https://linkedin.com/in/dev-user/resume.pdf"
}
```
**Responses:**
- `200 Created`: Profile successfully created.
- `400 Bad Request`: Invalid data or profile already exists.
---
### `PUT api/profile`
**Description:**  
Updates the existing user profile.

**Request Body Example:**
```json
{
"telegram":  "@dev_user",
"github":  "dev-user",
"bio":  "Backend developer specializing in Spring Boot.",
"position":  "Software Engineer",
"education":  "BACHELOR",
"expertise":  "Java, Microservices",
"expertise_level":  "MID",
"resume":  "https://linkedin.com/in/dev-user/resume.pdf"
}
```
**Responses:**
- `200 OK`: Profile updated successfully.
- `404 Not Found`: Profile does not exist.
- `400 Bad Request`: Invalid input.
---
### `GET api/profile/me`
**Description:**  
Fetches the personal profile of the authenticated user.

**Responses:**
- `200 OK`: Returns the user’s profile.
- `401 Unauthorized`: If the user is not logged in.
- `404 Not Found`: Profile not found.
