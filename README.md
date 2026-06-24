# SecTrack — Security Vulnerability Tracker API

A backend REST API built in Go for tracking and managing security vulnerabilities.
Built with Gin, GORM, and PostgreSQL — designed to mirror real-world security platform workflows.

# Backend deployed link ->https://sectrack.onrender.com

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **Auth:** JWT (golang-jwt)
- **Deployment:** Render

## Features

- JWT-based authentication (register, login)
- Full CRUD for vulnerability management
- Severity and status filtering
- Role-based access control (admin, analyst, viewer)
- Automatic audit logging on every write action
- Production deployment with managed PostgreSQL

  ## Try it Out

Base URL: `https://sectrack.onrender.com`

No setup needed — use these requests directly in Postman or curl.

### 1. Register a user
```bash
curl -X POST https://sectrack.onrender.com/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123","role":"analyst"}'
```

### 2. Login and get token
```bash
curl -X POST https://sectrack.onrender.com/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123"}'
```

### 3. Create a vulnerability (use token from step 2)
```bash
curl -X POST https://sectrack.onrender.com/vulns \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"cve_id":"CVE-2024-1234","title":"SQL Injection in login form","severity":"critical","affected_system":"Auth Service"}'
```

### 4. List all vulnerabilities
```bash
curl -X GET https://sectrack.onrender.com/vulns \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 5. Filter by severity
```bash
curl -X GET "https://sectrack.onrender.com/vulns?severity=critical" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 6. Update a vulnerability
```bash
curl -X PATCH https://sectrack.onrender.com/vulns/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"status":"resolved"}'
```

## Roles & Permissions

| Action | Admin | Analyst | Viewer |
|---|---|---|---|
| Create vulnerability | ✅ | ✅ | ❌ |
| Update vulnerability | ✅ | ✅ | ❌ |
| Delete vulnerability | ✅ | ❌ | ❌ |
| View vulnerabilities | ✅ | ✅ | ✅ |
| View audit logs | ✅ | ❌ | ❌ |

## API Endpoints

### Auth
| Method | Endpoint | Description |
|---|---|---|
| POST | /auth/register | Register a new user |
| POST | /auth/login | Login and get JWT token |

### Vulnerabilities
| Method | Endpoint | Description |
|---|---|---|
| POST | /vulns | Create a vulnerability |
| GET | /vulns | List all vulnerabilities |
| GET | /vulns/:id | Get a vulnerability by ID |
| PATCH | /vulns/:id | Update a vulnerability |
| DELETE | /vulns/:id | Delete a vulnerability |

### Audit Logs
| Method | Endpoint | Description |
|---|---|---|
| GET | /audit-logs | Get all audit logs (admin only) |

## Filtering
GET /vulns?severity=critical

GET /vulns?status=open

GET /vulns?severity=high&status=resolved
