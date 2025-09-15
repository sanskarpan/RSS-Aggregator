# RSS Feed Project

**RSS (Really Simple Syndication)** is a standardized web feed format that enables websites to share frequently updated content such as articles, blogs, or podcasts. It allows applications or users to subscribe and automatically access updates in a structured, computer-readable format.

🔗 **Sample RSS Feed:** [https://www.wagslane.dev/index.xml](https://www.wagslane.dev/index.xml)

---

## ⚙️ Tools & Technologies Used

1. **Routing:** [Chi Router](https://github.com/go-chi/chi) – Lightweight, idiomatic router for building Go HTTP services.
2. **Database:** **PostgreSQL** – Relational database for storing feed and metadata.
3. **Migrations:** [Goose](https://github.com/pressly/goose) – Database migrations for schema versioning and evolution.
4. **Code Generation:** [SQLC](https://github.com/sqlc-dev/sqlc) – Generates type-safe Go code directly from raw SQL queries.
5. **Containerization:** **Docker** – Used for running `sqlc generate` and isolating services.
6. **API Testing:** **Postman** – To test and validate API endpoints.

