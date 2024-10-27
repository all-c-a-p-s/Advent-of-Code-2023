mod day01;
mod day02;
mod day03;

fn main() {
    println!("day 1 part 1: {}", day01::part_one("./src/input01.txt"));
    println!("day 1 part 2: {}", day01::part_two("./src/input01.txt"));
    println!("day 2 part 1: {}", day02::part_one("./src/input02.txt"));
    println!("day 2 part 2: {}", day02::part_two("./src/input02.txt"));
    println!("day 3 part 1: {}", day03::part_one("./src/input03.txt"));
    println!("day 3 part 2: {}", day03::part_two("./src/input03.txt"));
}
