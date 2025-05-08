# 🛠 Distributed Task Scheduler in Go

A simple, concurrent task scheduler built in Go using a worker pool architecture. Tasks are submitted via an HTTP API, placed in a queue, and processed by a pool of workers concurrently.

---

## 🚀 Features

- RESTful API to submit and track tasks
- In-memory task queue with thread-safe coordination
- Worker pool with configurable size
- Concurrent task execution with status updates
- Clear separation between API, coordinator, and worker logic

---

## 🧱 Architecture Overview

```text
                +-------------+
POST /task  --→ |   API Layer | ---+
GET /task/:id   +-------------+    |
                                   v
                          +----------------+
                          |  Coordinator   | (Queue + Task Map)
                          +----------------+
                                   |
                                   v
                          [Dispatcher Loop]
                                   |
                                   v
                        +-----------------------+
                        |     Worker Pool       |
                        |    (N Goroutines)     |
                        +-----------------------+
