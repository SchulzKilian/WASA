# PhotoShare Social Network

A modern social media platform for sharing photos and connecting with others. Built with Go backend and Vue.js frontend.

## 🌟 Features

- 📸 Photo sharing and discovery
- 👤 User profiles and following system
- ❤️ Like and interact with posts
- 💬 Comment on shared content
- 🔍 Explore feed with trending photos
- 📱 Responsive design for mobile and desktop

## 🔧 Tech Stack

### Backend
- Go with modular architecture
- RESTful API design
- SQLite database
- JWT authentication

### Frontend
- Vue.js with Bootstrap
- Feather icons
- Modern dashboard interface
- NPM build system

## 🚀 Getting Started

### Prerequisites
- Go 1.x
- Node.js (LTS version)
- NPM

### Backend Setup
# Clone repository
git clone [your-repo-url]

# Build and run backend
go build ./cmd/webapi/
./webapi

### Frontend Development
# Start NPM container
./open-npm.sh

# Inside container:
npm install
npm run dev

### Production Build
# Build frontend
./open-npm.sh
npm run build-embed
exit

# Build backend with UI
go build -tags webui ./cmd/webapi/

## 📁 Project Structure

├── cmd/
│   ├── webapi/         # Main web server
│   └── healthcheck/    # Health monitoring
├── service/
│   ├── api/           # API implementation
│   └── globaltime/    # Time utilities
├── webui/             # Vue.js frontend
└── doc/              # API documentation

## 🛠️ Development

### API Documentation
- OpenAPI specification in `doc/api.yaml`
- RESTful endpoints for user and photo operations
- Comprehensive authentication system

### Frontend Features
- Intuitive photo upload interface
- Interactive feed with infinite scroll
- User profile customization
- Real-time notifications
- Mobile-first responsive design

## 🔒 Security Features
- JWT-based authentication
- Secure file handling
- Input validation
- CORS protection

## 📊 Core Functionality
- Photo upload and management
- User following system
- Activity feed
- Like and comment system
- User profiles
- Content discovery

## 🧪 Testing
# Run backend tests
go test ./...

# Run frontend tests
npm run test

## 💻 Development Commands

### Backend
# Run development server
go run ./cmd/webapi/

# Run with specific config
go run ./cmd/webapi/ -config path/to/config.yaml

### Frontend
# Development mode
npm run dev

# Production build
npm run build-embed

# Preview production build
npm run preview

## 📝 Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 🐛 Known Issues
For Apple M1/M2 users: If encountering esbuild issues, run:
./open-npm.sh
npm install
exit

## 📄 License
See LICENSE file for details.
```
