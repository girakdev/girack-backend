use crate::models::Message;
use actix_web::{get, post, web, HttpResponse, Responder};
use chrono::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
pub struct GetMessageRequest {
    // TODO: implement token
    // token: String,
    channel: String,
    count: Option<i64>,
    latest_datetime: Option<DateTime<Local>>,
    oldest_datetime: Option<DateTime<Local>>,
}

#[derive(Serialize)]
pub struct GetMessageResponse {
    ok: bool,
    messages: Vec<Message>,
}

#[get("/v1/message")]
pub async fn get_message_handler(request: web::Query<GetMessageRequest>) -> impl Responder {
    let messages = Message::find_many(
        &request.channel,
        request.count,
        request.latest_datetime,
        request.oldest_datetime,
    )
    .await
    .unwrap();

    HttpResponse::Ok().json(GetMessageResponse { ok: true, messages })
}

#[derive(Deserialize)]
pub struct PostMessageRequest {
    channel: String,
    text: String,
}

#[derive(Serialize)]
pub struct PostMessageResponse {
    ok: bool,
}

#[post("/v1/message")]
pub async fn post_message_handler(request: web::Json<PostMessageRequest>) -> impl Responder {
    Message::create(
        0, // TODO: implement user
        &request.channel,
        Local::now(),
        &request.text,
    )
    .await
    .unwrap();

    HttpResponse::Created().json(PostMessageResponse { ok: true })
}
