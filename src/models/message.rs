use crate::db;
use chrono::prelude::*;
use serde::Serialize;
use tokio_postgres::Error;

/// Message model
#[derive(Serialize)]
pub struct Message {
    pub id: i64,
    pub user_id: i64,
    pub channel: String,
    pub timestamp: DateTime<Local>,
    pub text: String,
    pub hidden: bool,
}

/// SQL for creating message table
pub const CREATE_MESSAGE_TABLE: &str = "
CREATE TABLE message (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGSERIAL NOT NULL,
    channel         TEXT NOT NULL,
    timestamp       TIMESTAMP WITH TIME ZONE NOT NULL,
    text            TEXT NOT NULL,
    hidden          BOOLEAN NOT NULL DEFAULT false
);
";

impl Message {
    /// Create a message in a database
    pub async fn create(
        user_id: i64,
        channel: &str,
        timestamp: DateTime<Local>,
        text: &str,
    ) -> Result<Message, Error> {
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
                user_id,
                channel,
                timestamp,
                text
            ) VALUES ($1, $2, $3, $4) RETURNING id, timestamp, hidden
            ",
            )
            .await?;

        let row = client
            .query_one(&statement, &[&user_id, &channel, &timestamp, &text])
            .await?;
        let id: i64 = row.get(0);
        let timestamp: DateTime<Local> = row.get(1);
        let hidden: bool = row.get(2);

        Ok(Message {
            id,
            user_id,
            channel: channel.to_string(),
            timestamp,
            text: text.to_string(),
            hidden,
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
        let user_id: i64 = row.get(1);
        let channel: String = row.get(2);
        let timestamp: DateTime<Local> = row.get(3);
        let text: String = row.get(4);
        let hidden: bool = row.get(5);

        Ok(Message {
            id,
            user_id,
            channel,
            timestamp,
            text,
            hidden,
        })
    }

    /// Find messages in a database by date
    pub async fn find_many(
        channel: &str,
        count: Option<i64>,
        latest_datetime: Option<DateTime<Local>>,
        oldest_datetime: Option<DateTime<Local>>,
    ) -> Result<Vec<Message>, Error> {
        let (client, connection) = db::connect().await?;
        tokio::spawn(async move {
            if let Err(e) = connection.await {
                eprintln!("connection error: {}", e);
            }
        });

        let count = match count {
            Some(count) => count,
            None => 50,
        };

        let oldest_datetime = match oldest_datetime {
            Some(oldest_datetime) => oldest_datetime,
            None => Local.ymd(1970, 1, 1).and_hms(0, 0, 0),
        };

        let latest_datetime = match latest_datetime {
            Some(latest_datetime) => latest_datetime,
            None => Local::now(),
        };

        let rows = client
            .query(
                "SELECT * FROM message WHERE channel = ($1) AND timestamp BETWEEN ($2) AND ($3) AND hidden = false ORDER BY timestamp DESC LIMIT ($4)",
                &[&channel, &oldest_datetime, &latest_datetime, &count],
            )
            .await?;

        Ok(rows
            .iter()
            .map(|row| {
                let id: i64 = row.get(0);
                let user_id: i64 = row.get(1);
                let channel: String = row.get(2);
                let timestamp: DateTime<Local> = row.get(3);
                let text: String = row.get(4);
                let hidden: bool = row.get(5);

                Message {
                    id,
                    user_id,
                    channel,
                    timestamp,
                    text,
                    hidden,
                }
            })
            .collect())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn create_and_find() {
        let timestamp = Local::now();

        let message = Message::create(0, "general", timestamp, "こんにちは")
            .await
            .unwrap();

        assert_eq!(0, message.user_id);
        assert_eq!("general", message.channel);
        assert_eq!(timestamp, message.timestamp);
        assert_eq!("こんにちは", message.text);

        let finded_message = Message::find(message.id).await.unwrap();

        assert_eq!(message.id, finded_message.id);
        assert_eq!(message.user_id, finded_message.user_id);
        assert_eq!(message.channel, finded_message.channel);
        assert_eq!(message.timestamp, finded_message.timestamp);
        assert_eq!(message.text, finded_message.text);
        assert_eq!(message.hidden, finded_message.hidden);
    }
}
