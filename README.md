# Social Network

A containerized full-stack social network built with a **Golang backend**, **Next.js frontend**, **SQLite**, and **WebSockets**. The application provides a Facebook-like experience with profiles, followers, posts, groups, notifications, and private or group chats.

> **Academic project — Zone01 Oujda**

## Table of contents

- [Overview](#overview)
- [Features](#features)
- [Technology stack](#technology-stack)
- [Project structure](#project-structure)
- [Getting started](#getting-started)
- [Running with Docker](#running-with-docker)
- [Configuration](#configuration)
- [Database and migrations](#database-and-migrations)
- [Authentication and authorization](#authentication-and-authorization)
- [Real-time communication](#real-time-communication)
- [Development](#development)
- [Team](#team)
- [Project objectives](#project-objectives)

## Overview

Social Network is a full-stack application designed to demonstrate the development and deployment of a modern social platform. Users can create accounts, manage their profiles, follow other users, publish posts, join groups, exchange messages, and receive real-time notifications.

The project is organized into separate backend and frontend services. The backend exposes the application API and manages authentication, business logic, persistence, file uploads, WebSockets, and migrations. The frontend provides the responsive user interface and communicates with the backend over HTTP and WebSockets.

## Features

### Authentication

- User registration and login
- Session-based authentication with cookies
- Secure password hashing with bcrypt
- Logout available throughout the application
- Optional avatar, nickname, and biography fields

### Profiles and followers

- Public and private profiles
- User information and activity
- Followers and following lists
- Follow and unfollow users
- Follow requests for private profiles
- Accept or decline pending follow requests
- Automatic following for public profiles

### Posts and comments

- Create posts and comments
- Attach JPEG, PNG, or GIF images
- Public posts visible to all users
- Follower-only posts
- Private posts shared with selected followers
- Group posts visible only to group members

### Groups and events

- Create groups with a title and description
- Browse available groups
- Invite users to join groups
- Request to join a group
- Accept or decline invitations and join requests
- Create posts and comments inside groups
- Create group events with a date, time, description, and attendance options
- Choose whether to attend an event
- Group chat for members

### Chat and notifications

- Private messaging between connected users
- Real-time delivery through WebSockets
- Emoji support
- Group chat rooms
- Notifications for follow requests, group invitations, join requests, and events
- Notifications remain distinct from private messages

## Technology stack

### Backend

- [Go](https://go.dev/)
- SQLite
- HTTP API
- WebSockets
- Sessions and cookies
- bcrypt password hashing
- Database migrations

### Frontend

- [Next.js](https://nextjs.org/)
- JavaScript
- HTML
- CSS
- Responsive client-side interface

### Infrastructure

- Docker
- Docker Compose
- Separate backend and frontend images

## Project structure

```text
.
├── backend/
│   ├── pkg/
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   └── sqlite/
│   │   │   └── sqlite/
│   │   └── ...
│   └── server.go
├── frontend/
│   ├── app/
│   ├── public/
│   ├── package.json
│   └── ...
├── docker-compose.yml
├── Dockerfile
└── README.md
```

The exact structure may evolve as the application grows, but migrations must remain in a dedicated SQLite migrations directory and must be applied when the backend starts.

## Getting started

### Prerequisites

Install the following tools before running the project locally:

- Go 1.20 or newer
- Node.js 18 or newer
- npm, pnpm, or yarn
- SQLite development dependencies required by the selected Go SQLite driver
- Docker and Docker Compose (recommended)

### Clone the repository

```bash
git clone https://github.com/yassinerahhaui/social-network.git
cd social-network
```

### Run the backend locally

```bash
cd backend
go mod download
go run .
```

The backend API will be available on the port configured by the application, commonly `http://localhost:8080`.

### Run the frontend locally

In a second terminal:

```bash
cd frontend
npm install
npm run dev
```

The frontend will be available at:

```text
http://localhost:3000
```

## Running with Docker

Docker is the recommended way to run the complete application because it provides isolated backend and frontend services with consistent dependencies.

From the project root:

```bash
docker compose up --build
```

To run the services in the background:

```bash
docker compose up --build -d
```

To stop the application:

```bash
docker compose down
```

To remove containers and persistent volumes as well:

```bash
docker compose down -v
```

The frontend and backend ports are defined in `docker-compose.yml`. Open the frontend URL shown by Docker after the services start.

## Configuration

Create environment files according to the configuration expected by each service. Typical settings include:

```env
# Backend
PORT=8080
DATABASE_PATH=./data/social-network.db
SESSION_SECRET=change-me
UPLOAD_DIR=./uploads

# Frontend
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_WS_URL=ws://localhost:8080
```

Never commit production secrets, session keys, or private credentials to the repository. Use a `.env.example` file to document required variables without exposing real values.

## Database and migrations

The project uses SQLite for persistence. Database migrations are stored in a dedicated directory and are applied when the backend starts.

Expected migration naming convention:

```text
backend/pkg/db/migrations/sqlite/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_create_posts_table.up.sql
├── 000002_create_posts_table.down.sql
└── ...
```

Migrations should cover the data model required for:

- Users and sessions
- Profiles and privacy settings
- Follow relationships and follow requests
- Posts, comments, and media
- Groups, memberships, invitations, and join requests
- Events and attendance responses
- Private and group messages
- Notifications

When changing the schema, add a new numbered migration instead of editing a migration that has already been applied in a shared environment.

## Authentication and authorization

The application uses sessions and cookies to keep users authenticated between requests. Passwords must never be stored in plain text; they are hashed before being persisted.

Authorization rules are enforced by the backend. Examples include:

- Private profile content is visible only to approved followers.
- Private posts are visible only to the selected followers.
- Group content is visible only to group members.
- Group administration actions are restricted to the group creator where required.
- Private chat is available when at least one user follows the other.

## Real-time communication

WebSockets are used for real-time features, especially:

- Private chat messages
- Group chat messages
- Live message delivery
- Real-time notifications where supported

The backend is responsible for authenticating WebSocket connections and ensuring that messages are delivered only to authorized recipients or group members.

## Development

### Backend checks

```bash
cd backend
gofmt -w .
go test ./...
go vet ./...
```

### Frontend checks

```bash
cd frontend
npm run lint
npm run build
```

### Recommended workflow

1. Create a focused branch for each feature or fix.
2. Add or update database migrations when the schema changes.
3. Validate authorization rules on the backend.
4. Test the API and WebSocket behavior.
5. Check responsive behavior in the frontend.
6. Run the backend and frontend checks before opening a pull request.

## Team

- Ibrahim El Harraq — `#ielharra`
- Yassine Rahhaoui — `#yrahhaou`
- Yasyn Nait Edderhm — `#ynaitedd`
- Omar Ait Benhammou — `#oaitbenh`
- Mohamed El-Fihry — `#melfihry`
- Oussama Benali — `#obenali`

## Project objectives

This project provides practical experience with:

- Full-stack application architecture
- Go backend development
- Next.js frontend development
- Authentication with sessions and cookies
- Password hashing and secure credential handling
- SQLite database design and SQL queries
- Database migrations
- Image upload and storage
- WebSocket communication
- Docker images and multi-container deployment
- Responsive frontend development
- Access control and privacy rules

## License

This project was created for educational purposes as part of the Zone01 Oujda curriculum.
