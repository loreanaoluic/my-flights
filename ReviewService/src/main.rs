#[macro_use] extern crate rocket;

mod utils;

use rocket::{routes, Rocket, Build};
use utils::ReviewDTO;
use std::{thread};
use rocket::serde::json::Json;

#[get("/<id>")]
fn get_all_by_airline_id(id: i32) -> Json<String> {
    thread::spawn(move || {
        Json(utils::get_all_by_airline_id(id).unwrap())
    }).join().unwrap()
}

#[post("/", format = "json", data = "<review>")]
fn create_review(review: rocket::serde::json::Json<ReviewDTO>) -> Json<String> {
    thread::spawn(move || {
        Json(utils::create_review(review).unwrap())
    }).join().unwrap()
}

#[launch]
fn rocket() -> Rocket<Build> {
    thread::spawn(|| {
        utils::seed_db();
    }).join().expect("Thread panicked");

    rocket::build().mount("/api/reviews", routes![get_all_by_airline_id, create_review])
}