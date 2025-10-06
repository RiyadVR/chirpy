# 🐦 Chirpy

Welcome to **Chirpy** — your friendly, Go-powered Twitter clone!  
Think of it as your cozy corner of the internet where you can post tiny thoughts (we call them *chirps*), manage your profile, and explore a full-featured REST API — all built with love and Go 💛.

---

## 📚 What’s This All About?

This project hatched from the **Boot.Dev HTTP Servers Course**, and it’s a complete backend for a mini social media app.  
It’s got everything you’d expect from a modern Go project:

- RESTful API design  
- JWT authentication  
- PostgreSQL database integration  
- Webhooks  
- Static file serving  
- Middleware magic 🪄  

Basically: it’s Go, but make it *social*.

---

## ✨ Features

Here’s what Chirpy can do (spoiler: a lot):

- 👤 **User Management:** Sign up, log in, and tweak your profile  
- 💬 **Chirps:** Post, read, and delete short messages (140 chars max, old-school style!)  
- 🔐 **Authentication:** Secure JWT-based login system  
- 🐥 **Chirpy Red:** Upgrade to a premium account via webhooks  
- 🚫 **Profanity Filter:** Keeps your chirps squeaky clean  
- 🧮 **Admin Dashboard:** Check metrics and manage the database  
- 🖼️ **Static File Serving:** For your frontend goodies  

---

## 🔧 Tech Stack

**Built With:**

- 🧠 **Language:** Go  
- 🗄️ **Database:** PostgreSQL  
- 🪪 **Auth:** JWT  
- 🛠️ **Router:** Go’s built-in `http.ServeMux`  
- 🧱 **SQL Builder:** SQLC  
- 🔄 **Migrations:** Goose  
- 🌿 **Env Management:** godotenv  
- 🔑 **Password Hashing:** bcrypt  
- 🆔 **UUIDs:** Google UUID  

Basically — it’s the Go developer’s dream stack. 😎

---

## 🚀 Getting Started

### 🧰 Prerequisites

- Go 1.21+  
- PostgreSQL  
- Goose (for migrations)  
- Git  

### 🪄 Installation Steps

1. **Clone this bird’s nest**
   ```bash
   git clone https://github.com/RiyadVR/chirpy.git
   cd chirpy

Chirpy Setup Guide
Installation

Grab the dependencies
go mod download


Install Goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest


Set up your .env file
DB_URL=postgres://username:password@localhost/chirpy?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key
POLKA_KEY=your-polka-api-key
PLATFORM=dev


Spin up your DB and run the app
go run .

Your Chirpy server will start at http://localhost:8080 🎉


📖 API Docs
All the nitty-gritty API details (endpoints, formats, tokens, etc.) can be found in the API Documentation.
You’ll find everything from:

User registration/login
Chirp creation and deletion
Token refresh/revoke
Webhook endpoints
Admin controls

Basically, everything a proper API should chirp about. 🐣
🔐 How Auth Works
We use JWT (JSON Web Tokens) for authentication:

Register or login → get access & refresh tokens
Send requests with your access token:Authorization: Bearer <your-token>


Refresh tokens when they expire
Log out → revoke refresh tokens

Easy peasy, secure and sleek 🔒
🎯 Cool Stuff Inside
🐤 Chirps

Max 140 characters (classic Twitter vibes)
Built-in profanity filter (sorry “kerfuffle” lovers)
Delete only your own chirps — no drama here

🟥 Chirpy Red

Upgrade to premium via Polka payments
Webhook integration keeps user status in sync
Because every bird deserves to shine ✨

🧑‍💻 Admin Zone

App metrics dashboard
DB reset (for dev mode)
Request monitoring and hit counts

🧪 Dev Notes
🗃️ Database Migrations
We use Goose to keep DB changes smooth.
Common commands:
# Run all migrations
goose -dir sql/schema postgres $DB_URL up

# Rollback last migration
goose -dir sql/schema postgres $DB_URL down

# Check status
goose -dir sql/schema postgres $DB_URL status

# Create new migration
goose -dir sql/schema create migration_name sql

🧑‍🔬 Dev Mode
Set PLATFORM=dev in .env to unlock:

/admin/reset endpoint
Extra debug logs

✅ Testing
go test ./...

📄 License
Open source and proud!Licensed under the MIT License — do what you want, just give credit 🫶
🙏 Thanks & Credits
Massive shoutout to:

🏫 Boot.Dev — for the awesome HTTP Servers course
🧑‍💻 Go community — for top-tier tools & libraries
🐘 PostgreSQL team — for a rock-solid database

Built with ❤️ and a lot of Go by following Boot.Dev’s course.
Happy chirping! 🐦💬