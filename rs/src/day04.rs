fn get_numbers(s: String) -> Vec<i32> {
    s.split_whitespace()
        .map(|x| x.parse::<i32>().unwrap())
        .collect()
}

fn count_repeats(v1: Vec<i32>, v2: Vec<i32>) -> i32 {
    v1.iter()
        .fold(0, |acc, x| if v2.contains(&x) { acc + 1 } else { acc })
}

fn pow(a: i32, n: i32) -> i32 {
    match n {
        0 => 1,
        _ => a * pow(a, n - 1),
    }
}

fn line_score_part_one(line: String) -> i32 {
    let parts: Vec<String> = line.split(':').map(String::from).collect();
    let number_sets: Vec<Vec<i32>> = parts[1]
        .split('|')
        .map(String::from)
        .map(get_numbers)
        .collect();

    pow(
        2,
        count_repeats(number_sets[0].clone(), number_sets[1].clone()),
    )
}

pub fn part_one(filename: &str) -> i32 {
    let lines = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(|x| String::from(x))
        .collect::<Vec<String>>();

    lines
        .iter()
        .fold(0, |acc, l| acc + line_score_part_one(l.to_owned()))
}

pub fn part_two(filename: &str) -> i32 {
    let lines = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(|x| String::from(x))
        .collect::<Vec<String>>();
    let mut copies = vec![1; lines.len()];
    for i in 0..lines.len() {
        let parts: Vec<String> = lines[i].split(':').map(String::from).collect();
        let number_sets: Vec<Vec<i32>> = parts[1]
            .split('|')
            .map(String::from)
            .map(get_numbers)
            .collect();
        let repeats = count_repeats(number_sets[0].clone(), number_sets[1].clone());
        for j in 1..=repeats {
            copies[i + j as usize] += copies[i];
        }
    }
    copies.iter().sum()
}
