use std::collections::HashMap;

fn digit_values() -> HashMap<char, i32> {
    let mut digit_values: HashMap<char, i32> = HashMap::new();
    digit_values.insert('0', 0);
    digit_values.insert('1', 1);
    digit_values.insert('2', 2);
    digit_values.insert('3', 3);
    digit_values.insert('4', 4);
    digit_values.insert('5', 5);
    digit_values.insert('6', 6);
    digit_values.insert('7', 7);
    digit_values.insert('8', 8);
    digit_values.insert('9', 9);
    digit_values
}

fn line_score(l: &&str) -> i32 {
    let digits: Vec<char> = l
        .chars()
        .filter(|c| digit_values().contains_key(c))
        .collect();
    digit_values()[&digits[0]] * 10 + digit_values()[&digits[digits.len() - 1]]
}

fn first_digit(line: &String) -> i32 {
    match line {
        s if s.starts_with("one") || s.starts_with('1') => 1,
        s if s.starts_with("two") || s.starts_with('2') => 2,
        s if s.starts_with("three") || s.starts_with('3') => 3,
        s if s.starts_with("four") || s.starts_with('4') => 4,
        s if s.starts_with("five") || s.starts_with('5') => 5,
        s if s.starts_with("six") || s.starts_with('6') => 6,
        s if s.starts_with("seven") || s.starts_with('7') => 7,
        s if s.starts_with("eight") || s.starts_with('8') => 8,
        s if s.starts_with("nine") || s.starts_with('9') => 9,
        s if s.starts_with("zero") || s.starts_with('0') => 0,

        _ => first_digit(&line[1..].to_string()),
    }
}

fn last_digit(line: &String) -> i32 {
    match line {
        s if s.ends_with("one") || s.ends_with('1') => 1,
        s if s.ends_with("two") || s.ends_with('2') => 2,
        s if s.ends_with("three") || s.ends_with('3') => 3,
        s if s.ends_with("four") || s.ends_with('4') => 4,
        s if s.ends_with("five") || s.ends_with('5') => 5,
        s if s.ends_with("six") || s.ends_with('6') => 6,
        s if s.ends_with("seven") || s.ends_with('7') => 7,
        s if s.ends_with("eight") || s.ends_with('8') => 8,
        s if s.ends_with("nine") || s.ends_with('9') => 9,
        s if s.ends_with("zero") || s.ends_with('0') => 0,

        _ => last_digit(&line[..line.len() - 1].to_string()),
    }
}

fn line_score_part_two(line: &String) -> i32 {
    10 * first_digit(line) + last_digit(line)
}

pub fn part_one(filename: &str) -> i32 {
    let f = std::fs::read_to_string(filename).expect("failed to read file");
    let lines: Vec<&str> = f.lines().collect();

    lines.iter().map(line_score).sum()
}

pub fn part_two(filename: &str) -> i32 {
    let f = std::fs::read_to_string(filename).expect("failed to read file");
    let lines: Vec<&str> = f.lines().collect();

    lines
        .iter()
        .map(|x| line_score_part_two(&x.to_string()))
        .sum()
}
