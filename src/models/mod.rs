pub mod message;

pub use message::Message;

pub mod sql {
    pub use super::message::CREATE_MESSAGE_TABLE;
}
