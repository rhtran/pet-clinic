# Pet Clinic REST API using Axum & Tokio

## Introduction
* Tokio is an asynchronous runtime for the Rust programming language. It provides the building blocks needed for writing network applications. It gives you the ability to write fast, reliable, and highly concurrent applications.
* Axum is a web framework built on top of Tokio and Tower. It is designed to be both ergonomic and performant. It is built on top of hyper, a fast and low-level HTTP implementation in Rust.

## Technologies Stacks
* [Rust](https://www.rust-lang.org/)
* [Tokio](https://tokio.rs/)
* [Axum](https://docs.rs/axum/latest/axum/)
* [Serde](https://serde.rs/)
* [Tower HTTP](https://github.com/tower-rs/tower-http)



## Create a new project

```bash
cargo new axum-petclinic-service --bin
cd axum-petclinic-service
```
Add tokio, serde, tower-http and axum dependencies to Cargo.toml

```toml
axum = { version = "0.7.5", features = ["macros"] }
tokio = { version = "1.37.0", features = ["full"] }
serde = { version = "1.0.199", features = ["derive"] }
tower-http = { version = "0.5.2" }

[dev-dependencies]
anyhow = { version = "1.0.82" }
httpc-test = { version = "0.1.9" }
```

## Simple health & info REST API

Create a module named api and add health and info modules to it.

health module
```rust
use axum::{Json, Router};
use axum::routing::get;
use serde::Serialize;


#[derive(Debug, Serialize)]
struct HealthCheck {
    status: String,
}

pub fn routes() -> Router {
    // let app_state = AppState { mc };
    Router::new().route("/health", get(check))
}

async fn check() -> Json<HealthCheck> {
    let health_check = HealthCheck {
        status: "UP".to_string(),
    };

    Json(health_check)
}

```

info module
```rust
use axum::{Json, Router};
use axum::routing::get;
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
struct InfoResponse {
    pub version: String,
    pub name: String,
    pub description: String,
    pub author: String,
    pub license: String,
}

pub fn routes() -> Router {
    // let app_state = AppState { mc };
    Router::new().route("/info", get(info))
}

async fn info() -> Json<InfoResponse> {
    let info = InfoResponse {
        version: "0.1.0".to_string(),
        name: "Rust API".to_string(),
        description: "Rust API using Axum".to_string(),
        author: "Ethan J. Tran".to_string(),
        license: "MIT".to_string(),
    };

    Json(info)
}
```

main.rs
```rust
use std::time::Duration;
use axum::{Router};
use tokio::net::TcpListener;
use tokio::signal;
use tower_http::timeout::TimeoutLayer;
use tower_http::trace::TraceLayer;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

mod api;
use crate::api::health;
use crate::api::info;

#[tokio::main]
async fn main() {
    tracing_subscriber::registry()
        .with(
            tracing_subscriber::EnvFilter::try_from_default_env()
                .unwrap_or_else(|_| "example_tls_graceful_shutdown=debug".into()),
        )
        .with(tracing_subscriber::fmt::layer())
        .init();


    let routes_all = Router::new()
        .merge(app_routes());

    fn app_routes() -> Router {
        Router::new()
            .merge(health::routes())
            .merge(info::routes())
            .layer((
                       TraceLayer::new_for_http(),
                       // Graceful shutdown will wait for outstanding requests to complete.
                       // Add a timeout so requests don't respond forever.
                       TimeoutLayer::new(Duration::from_secs(10))),)
    }

    let listener = TcpListener::bind("127.0.0.1:3003")
        .await
        .unwrap();
    println!("Listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, routes_all)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

/// Wait for a shutdown signal to gracefully shutdown the server.
async fn shutdown_signal() {
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
        let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
        let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }
}
```

### Run the application
Run the application using the following commands.
* 1st terminal:
```bash
make watch-src-change
```

### Test the application
Test the application using the following commands.
* 2nd terminal:
```bash
make watch-test-change
```




