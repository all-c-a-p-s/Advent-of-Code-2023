mod day01;
mod day02;

fn main() {
    println!("day 1 part 1: {}", day01::part_one("./src/input01.txt"));
    println!("day 1 part 2: {}", day01::part_two("./src/input01.txt"));
    println!("day 2 part 1: {}", day02::part_one("./src/input02.txt"));
    println!("day 2 part 1: {}", day02::part_two("./src/input02.txt"));
}
