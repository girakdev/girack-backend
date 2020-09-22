use serde::Serialize;

// Channel Model
#[derive(Serialize)]
pub struct Channel {
    pub id: i64,
    pub name: String,
    pub is_channel: bool,
    pub created: i64,
    pub creator: i64,
    pub member: Vec<i64>,
    pub topic: String,
}
