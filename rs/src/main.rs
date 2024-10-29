use std::io::{stdout, Write};
use std::time::Instant;

mod day01;
mod day02;
mod day03;
mod day04;
mod day11;

fn take_int() -> i32 {
    let mut input = String::new();
    std::io::stdin().read_line(&mut input).unwrap();
    input.trim().parse().unwrap()
}

fn main() {
    print!("Enter day to solve: ");
    let _ = stdout().flush();
    let day = take_int();

    print!("Enter part to solve: ");
    let _ = stdout().flush();
    let part = take_int();

    println!();
    let start = Instant::now();

    print!("Answer: ");
    match (day, part) {
        (1, 1) => println!("{}", day01::part_one("./src/input01.txt")),
        (1, 2) => println!("{}", day01::part_two("./src/input01.txt")),
        (2, 1) => println!("{}", day02::part_one("./src/input02.txt")),
        (2, 2) => println!("{}", day02::part_two("./src/input02.txt")),
        (3, 1) => println!("{}", day03::part_one("./src/input03.txt")),
        (3, 2) => println!("{}", day03::part_two("./src/input03.txt")),
        (4, 1) => println!("{}", day04::part_one("./src/input04.txt")),
        (4, 2) => println!("{}", day04::part_two("./src/input04.txt")),
        (11, 1) => println!("{}", day11::part_one("./src/input11.txt")),
        (11, 2) => println!("{}", day11::part_two("./src/input11.txt")),
        _ => panic!("sorry not implemented yet"),
    };

    let duration = start.elapsed();
    println!("Runtime: {:?}", duration);
}
