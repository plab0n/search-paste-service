# Paste Service with Semantic Search

[![Go](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15%2B-orange)](https://www.postgresql.org/)
[![Elasticsearch](https://img.shields.io/badge/Elasticsearch-8%2B-green)](https://www.elastic.co/)

This service provides an efficient and scalable solution for storing, indexing, and retrieving text and URL-based content through semantic search. Users can seamlessly submit text snippets or URLs, which the service then securely stores in a PostgreSQL database. Leveraging advanced NLP techniques, each stored entry is transformed into high-dimensional vector embeddings representing their semantic content.

Elasticsearch powers the semantic search capability, indexing these vector embeddings to allow fast and accurate retrieval of semantically similar pastes. Users can query the service using natural language or specific keywords, and the system efficiently returns relevant content ranked by semantic similarity.
This architecture integrates PostgreSQL for reliable storage and management of structured data, while Elasticsearch provides powerful, real-time semantic search functionality.

---

## Table of Contents
- [Features](#features)
- [API Documentation](#api-documentation)
- [How It Works](#how-it-works)
- [Setup](#setup)
- [Usage Examples](#usage-examples)

---

## Features
- **Store Text/URLs**: Submit plain text or URLs (content is scraped and stored).
- **Semantic Search**: Find pastes using vector similarity via Elasticsearch embeddings.
- **CRUD Operations**: Add, delete, update, and fetch pastes via REST APIs.
- **Scalable Storage**: PostgreSQL for metadata, Elasticsearch for vector embeddings.

---

## API Documentation
Explore endpoints interactively using Swagger UI at `http://<host>:<port>/swagger/index.html`.

### Key Endpoints
| Method | Endpoint                | Description                          | Parameters/Body                                   |
|--------|-------------------------|--------------------------------------|--------------------------------------------------|
| POST   | `/api/paste/add`        | Add text/URL                         | `title` (string), `paste` (string)               |
| DELETE | `/api/paste/delete/{id}`| Delete a paste by ID                 | `id` (integer, path)                             |
| PATCH  | `/api/paste/update`     | Update a paste                       | `id` (string), `paste` (string) (both in body)   |
| GET    | `/api/paste/{id}`       | Fetch a paste by ID                  | `id` (integer, path)                             |
| GET    | `/api/pastes`           | List all pastes                      | None                                             |

### Response Models
- **`IDResponse`**: Returns the ID of a newly created/updated paste.
  ```json
  { "id": 123 }

- **`GetBookResponse`**: Returns paste details
  ```json
  {
    "id": 123,
    "title": "Sample Title",
    "paste": "Text content or scraped URL data",
    "createdBy": "user123",
    "createdAt": "2023-09-01T12:34:56Z"
  }

# Paste Service API

A high-performance paste service built in Go that accepts plain text or URLs, stores content in PostgreSQL, and leverages semantic search by generating vector embeddings. For URL inputs, the service scrapes the web page content, generates embeddings, and indexes them in Elasticsearch to enable fast, relevant search results based on vector similarity.

## How It Works

### Add a Paste:
- **Text:** Stored directly in PostgreSQL.
- **URL:** Content is scraped and stored in PostgreSQL.

### Embedding Generation:
- Text/URL content is converted into a vector embedding (e.g., using a pre-trained NLP model like Sentence-BERT).

### Elasticsearch Indexing:
- Embeddings are indexed in Elasticsearch for fast similarity searches.

### Semantic Search:
- User input is vectorized, and Elasticsearch returns pastes with the closest embeddings.

## Setup

### Prerequisites
- **PostgreSQL** (v15+)
- **Elasticsearch** (v8+)
- **Go** (1.20+)
- **Ml service for embedding** (1.20+)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/plab0n/search-paste-service.git
   cd search-paste-service

1. **Configure Environment Variables (Create `.env`):**
    ```bash
    POSTGRES_URL="postgres://user:password@localhost:5432/pastedb"
    ELASTICSEARCH_URL="http://localhost:9200"
    HTTP_HOST="0.0.0.0:8080"

3. **Install Dependencies:**
    ```bash
    go mod download

4. **Start the service:**
   ```bash
    go run main.go

### Usage Examples

1. **Add a Text Paste:**
   ```bash
   curl -X POST http://localhost:8080/api/paste/add \
    -H "Content-Type: application/json" \
    -d '{
        "title": "My Notes",
        "paste": "Semantic search uses vector embeddings."
    }'

1. **Add a URL Paste:**
    ```bash
    curl -X POST http://localhost:8080/api/paste/add \
    -H "Content-Type: application/json" \
    -d '{
    "title": "Example Blog",
    "paste": "https://example.com/blog"
    }'

3. **Fetch a Paste:**
    ```bash
    curl http://localhost:8080/api/paste/123

4. **Delete a paste:**
   ```bash
    curl -X DELETE http://localhost:8080/api/paste/delete/123

# License: MIT
Contribute: PRs welcome!
