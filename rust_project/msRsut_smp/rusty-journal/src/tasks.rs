use chrono::{DateTime, Utc};

#[derive(Debug)]
pub struct Task {
    pub text: String,
    pub created_at: DateTime<Utc>,
}

impl Task {
    pub fn new(text: String) -> Task {
        let created_at: DateTime<Utc> = Utc::now();
        Task { text, created_at }
    }
}