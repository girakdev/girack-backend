use actix_web::{get, HttpResponse, Responder};

#[get("/")]
pub async fn get_hello_handler() -> impl Responder {
    HttpResponse::Ok().body("Hello, owrld!")
}
