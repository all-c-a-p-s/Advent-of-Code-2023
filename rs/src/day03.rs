use std::collections::HashSet;

fn is_symbol(c: char) -> bool {
    !(c.is_ascii_digit() || c == '.')
}

fn coords_to_check(xs: (usize, usize), y: usize, rows: usize, cols: usize) -> Vec<(usize, usize)> {
    //where rows and cols are the max indices, not the number of rows/columns
    let rows_to_check: Vec<usize> = match y {
        0 => vec![0, 1],
        r if r == rows => vec![rows - 1, rows],
        _ => vec![y - 1, y, y + 1],
    };
    let columns_to_check: Vec<usize> = match xs {
        (0, c) if c == cols => (0..=cols).collect(),
        (0, x2) => (0..=(x2 + 1)).collect(),
        (x1, c) if c == cols => ((x1 - 1)..=cols).collect(),
        (x1, x2) => ((x1 - 1)..=(x2 + 1)).collect(),
    };
    let mut coords_to_check = Vec::new();
    for r in rows_to_check {
        for c in &columns_to_check {
            coords_to_check.push((r, *c))
        }
    }
    coords_to_check
}

pub fn part_one(filename: &str) -> i32 {
    let lines: Vec<Vec<char>> = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .map(|x| x.chars().collect())
        .collect();
    let mut total = 0;
    for l in 0..lines.len() {
        let mut start = 0usize;
        let mut current_num: String = String::new();
        for i in 0..lines[l].len() {
            if lines[l][i].is_ascii_digit() {
                if current_num.is_empty() {
                    start = i;
                }

                current_num = current_num + lines[l][i].to_string().as_str();

                let is_end = match i {
                    x if x == lines[l].len() - 1 => true,
                    y if !lines[l][y + 1].is_ascii_digit() => true,
                    _ => false,
                };
                if is_end {
                    let n: i32 = current_num.parse().unwrap();
                    let to_check =
                        coords_to_check((start, i), l, lines.len() - 1, lines[0].len() - 1);
                    'a: for (r, c) in to_check {
                        if is_symbol(lines[r][c]) {
                            total += n;
                            break 'a;
                        }
                    }
                    current_num = String::new();
                }
            }
        }
    }
    total
}

fn get_number_indices(row: &Vec<char>, rown: usize, idx: usize) -> (usize, usize, usize) {
    //return indices rather than number to handle case when the same number is found twice
    //this way we can check if the indices are the same so not too account for both
    let (mut start, mut end) = (0, row.len() - 1);
    for i in (0..idx).rev() {
        if !row[i].is_ascii_digit() {
            start = i + 1;
            break;
        }
    }
    for i in idx..row.len() {
        if !row[i].is_ascii_digit() {
            end = i - 1;
            break;
        }
    }
    (rown, start, end)
}

pub fn part_two(filename: &str) -> i32 {
    let lines: Vec<Vec<char>> = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .map(|x| x.chars().collect())
        .collect();
    const GEAR: char = '*';
    let mut total = 0;
    for l in 0..lines.len() {
        for i in 0..lines[l].len() {
            if lines[l][i] == GEAR {
                let to_check = coords_to_check((i, i), l, lines.len(), lines[0].len());
                let mut number_indices = vec![];
                for (r, c) in to_check {
                    if lines[r][c].is_ascii_digit() {
                        number_indices.push(get_number_indices(&lines[r], r, c));
                    }
                }
                let unique_numbers: HashSet<(usize, usize, usize)> =
                    number_indices.into_iter().collect();
                if unique_numbers.len() == 2 {
                    let v = unique_numbers
                        .into_iter()
                        .collect::<Vec<(usize, usize, usize)>>();
                    total += v.iter().fold(1, |acc, n| {
                        acc * lines[n.0][n.1..=n.2]
                            .iter()
                            .collect::<String>()
                            .parse::<i32>()
                            .unwrap()
                    });
                }
            }
        }
    }
    total
}
