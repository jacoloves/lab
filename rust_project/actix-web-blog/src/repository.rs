use crate::schema::*;
use diesel::{Insertable, Queryable};
use diesel::r2d2::{self, ConnectionManager};
use diesel::sqlite::SqliteConnection;
use serde::{Deserialize, Deserializer, Serialize};
use diesel::prelude::*;
use crate::error::ApiError;
use actix_web::web;

type DbPool = r2d2::Pool<ConnectionManager<SqliteConnection>>;

#[derive(Deserialize, Insertable)]
#[diesel(table_name = posts)]
pub struct NewPost {
    #[serde(deserialize_with = "max100")]
    title: String,
    body: String,
}

fn max100<'de, D>(de: D) -> Result<String, D::Error> where D: Deserializer<'de>,
{
    String::deserialize(de).and_then(|s| {
        if !s.is_empty() && s.len() <= 100 {
            Ok(s)
        } else {
            Err(serde::de::Error::custom("string length is 0 or too long.",))
        }
    })
}

#[derive(Serialize, Queryable)]
pub struct Post {
    id: i32,
    title: String,
    body: String,
    published: bool,
}

pub struct Repository {
    pool: DbPool,
}

impl Repository {
    pub fn new(database_url: &str) -> Self {
        let manager = ConnectionManager::<SqliteConnection,>::new(database_url);
        let pool = r2d2::Pool::builder().build(manager).expect("Failed to create a pool.");

        Self { pool }
    }

    pub async fn create_post(&self, new_post: NewPost,) -> Result<Post, ApiError> {
        let mut conn = self.pool.get()?;
        let post = web::block(move || {
            diesel::insert_into(posts::table).values(new_post).get_result(&mut conn)
        })
        .await??;

        Ok(post)
    }

    pub async fn list_posts(&self,) -> Result<Vec<Post>, ApiError> {
        let mut conn = self.pool.get()?;
        let res = web::block(move || {
            posts::table.load(&mut conn)
        })
        .await??;

        Ok(res)
    }

    pub async fn get_post(&self, id: i32,) -> Result<Post, ApiError> {
        let mut conn = self.pool.get()?;
        let res = web::block(move || {
            posts::table.find(id).first(&mut conn).optional()
        })
        .await??
        .ok_or(ApiError::NotFound)?;

        Ok(res)
    }
}

impl NewPost {
    pub fn validate(&self) -> bool {
        self.title.len() > 0 && self.title.len() <= 100
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_conn() {
        let databse_url = std::env::var("DATABASE_URL").unwrap();
        let repo = Repository::new(&databse_url);
        assert!(repo.pool.get().is_ok());
    }
}