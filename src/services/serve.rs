use super::*;
use crate::util::get_env;
use actix_web::{middleware, App, HttpServer};

/// Run HTTP server
pub async fn run() -> std::io::Result<()> {
    HttpServer::new(|| {
        // Logger
        std::env::set_var("RUST_LOG", "actix_web=info");
        env_logger::init();

        App::new()
            .wrap(middleware::Logger::default())
            .service(get_message_handler)
            .service(post_message_handler)
    })
    .bind(format!("0.0.0.0:{}", get_env("BACKEND_PORT")))?
    .run()
    .await
}
