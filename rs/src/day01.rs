use std::collections::HashMap;

pub fn line_score(l: &&str) -> i32 {
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

    let digits: Vec<char> = l.chars().filter(|c| digit_values.contains_key(c)).collect();
    digit_values[&digits[0]] * 10 + digit_values[&digits[digits.len() - 1]]
}

pub fn part_one(filename: &str) -> i32 {
    let f = std::fs::read_to_string(filename).expect("failed to read file");
    let lines: Vec<&str> = f.lines().collect();

    lines.iter().map(line_score).sum()
}
