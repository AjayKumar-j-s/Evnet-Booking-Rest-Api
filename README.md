Event Booking REST API

A Go-powered RESTful API for creating, managing, and booking events. 

📌 Features

- User registration and authentication
- Create, update, and delete events (only by event creators)
- Register for and cancel event participation
- View a list of available events

🧩 Endpoints

| Method | Endpoint                    | Description                        | Authentication |
|--------|-----------------------------|------------------------------------|----------------|
| GET    | `/events`                   | Get a list of available events     | ❌ No          |
| GET    | `/events/<id>`              | Get details of a single event      | ❌ No          |
| POST   | `/events`                   | Create a new bookable event        | ✅ Yes         |
| PUT    | `/events/<id>`              | Update an event                    | ✅ Creator Only|
| DELETE | `/events/<id>`              | Delete an event                    | ✅ Creator Only|
| POST   | `/signup`                   | Create a new user                  | ❌ No          |
| POST   | `/login`                    | Authenticate user                  | ❌ No          |
| POST   | `/events/<id>/register`     | Register user for an event         | ✅ Yes         |
| DELETE | `/events/<id>/register`     | Cancel event registration          | ✅ Yes         |

🔐 Authentication

- JWT-based authentication is used.
- After login, include the token in the `Authorization` header for all protected routes.

Example:
Authorization: Bearer <your_token_here>
 📦 Setup & Installation

🛠 Tech Stack
Go (Golang)
gin
SQLite
JWT for authentication
