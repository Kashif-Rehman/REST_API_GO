# Profile Management API

This Go application provides a simple RESTful API for managing user profiles. It uses the Gorilla mux router for handling HTTP requests.

## Endpoints

- **Add a New Profile:**

  - Endpoint: `/profiles`
  - Method: `POST`
  - Description: Add a new user profile.
  - Request Body: JSON representing the new profile.

- **Get All Profiles:**

  - Endpoint: `/profiles`
  - Method: `GET`
  - Description: Retrieve all user profiles.

- **Get a Profile by ID:**

  - Endpoint: `/profiles/{id}`
  - Method: `GET`
  - Description: Retrieve a user profile by its ID.
  - URL Parameter: `id` - The ID of the profile.

- **Update a Profile by ID:**

  - Endpoint: `/profiles/{id}`
  - Method: `PUT`
  - Description: Update a user profile by its ID.
  - URL Parameter: `id` - The ID of the profile to be updated.
  - Request Body: JSON representing the updated profile.

- **Delete a Profile by ID:**
  - Endpoint: `/profiles/{id}`
  - Method: `DELETE`
  - Description: Delete a user profile by its ID.
  - URL Parameter: `id` - The ID of the profile to be deleted.

## How to Run

Follow these steps to run the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/kashif-rehman/rest_api_go.git
   ```
