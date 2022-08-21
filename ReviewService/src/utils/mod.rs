use postgres::{Client, Error, NoTls};
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Review {
    id: i32,
    user_id: i32,
    username: String,
    airline_id: i32,
    message: String,
    rating: i32
}

#[derive(Serialize, Deserialize)]
pub struct ReviewDTO {
    user_id: i32,
    username: String,
    airline_id: i32,
    message: String,
    rating: i32
}

#[derive(Serialize)]
struct Response {
    message: String
}

pub fn seed_db() -> Result<(), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:loreana@localhost:5432/flights",
        NoTls,
    )?;

    client.batch_execute("DROP TABLE IF EXISTS reviews")?;

    client.batch_execute(
        "
        CREATE TABLE IF NOT EXISTS reviews (
            id              SERIAL PRIMARY KEY,
            user_id         INTEGER,
            username        VARCHAR NOT NULL,
            airline_id      INTEGER,
            message         VARCHAR NOT NULL,
            rating          INTEGER
            )
    ",
    )?;

    client.execute(
        "INSERT INTO reviews (user_id, airline_id, username, message, rating) VALUES ($1, $2, $3, $4, $5)",
        &[&2.to_owned(), &2.to_owned(), &"user", &"Excellent", &5.to_owned()],
    )?;

    client.execute(
        "INSERT INTO reviews (user_id, airline_id, username, message, rating) VALUES ($1, $2, $3, $4, $5)",
        &[&3.to_owned(), &2.to_owned(), &"user2", &"Good", &4.to_owned()],
    )?;

    client.close()?;

    Ok(())
}

pub fn get_all_by_airline_id(id: i32) -> Result<String, Error> {
    let mut client = Client::connect(
        "postgresql://postgres:loreana@localhost:5432/flights",
        NoTls,
    )?;

    let mut ret: Vec<Review> = vec![];
    for row in client.query("SELECT id, user_id, airline_id, username, message, rating FROM reviews WHERE airline_id = $1", &[&id])? {
        let id: i32 = row.get(0);
        let user_id: i32 = row.get(1);
        let airline_id: i32 = row.get(2);
        let username: &str = row.get(3);
        let message: &str = row.get(4);
        let rating: i32 = row.get(5);

        ret.push(Review { id: (id), user_id: (user_id), username: (username.to_string()), airline_id: (airline_id), message: (message.to_string()), rating: (rating) });
    }
    client.close()?;

    Ok(serde_json::to_string(&ret).unwrap())
}

pub fn create_review(review: rocket::serde::json::Json<ReviewDTO>) -> Result<String, Error> {
    let mut client = Client::connect(
        "postgresql://postgres:loreana@localhost:5432/flights",
        NoTls,
    )?;
    
    client.execute(
        "INSERT INTO reviews (user_id, airline_id, username, message, rating) VALUES ($1, $2, $3, $4, $5)",
        &[&review.user_id, &review.airline_id, &review.username, &review.message, &review.rating],
    )?;

    client.close()?;

    Ok(serde_json::to_string(&Response{message: "Made successfully.".to_string()}).unwrap())
}