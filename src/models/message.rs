use crate::db;
use tokio_postgres::Error;

/// Message model
pub struct Message {
    pub id: i64,
    pub text: String,
}

/// SQL for creating message table
pub const CREATE_MESSAGE_TABLE: &str = "
CREATE TABLE message (
    id              BIGSERIAL PRIMARY KEY,
    text            TEXT NOT NULL,
);
";

impl Message {
    /// Create a message in a database
    pub async fn create(text: &str) -> Result<Message, Error> {
        let (client, connection) = db::connect().await?;
        tokio::spawn(async move {
            if let Err(e) = connection.await {
                eprintln!("connection error: {}", e);
            }
        });

        let statement = client
            .prepare(
                "
            INSERT INTO message (
                text,
            ) VALUES ($1) RETURNING id
            ",
            )
            .await?;

        let row = client.query_one(&statement, &[&text]).await?;
        let id: i64 = row.get(0);

        Ok(Message {
            id,
            text: text.to_string(),
        })
    }

    /// Find a message in a database
    pub async fn find(id: i64) -> Result<Message, Error> {
        let (client, connection) = db::connect().await?;
        tokio::spawn(async move {
            if let Err(e) = connection.await {
                eprintln!("connection error: {}", e);
            }
        });

        let row = client
            .query_one("SELECT * FROM message WHERE id = ($1)", &[&id])
            .await?;
        let id: i64 = row.get(0);
        let text: String = row.get(9);

        Ok(Message { id, text })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn create_and_find() {
        let message = Message::create("こんにちは").await.unwrap();

        assert_eq!("こんにちは", message.text);

        let finded_message = Message::find(message.id).await.unwrap();

        assert_eq!(message.id, finded_message.id);
        assert_eq!(message.text, finded_message.text);
    }
}
