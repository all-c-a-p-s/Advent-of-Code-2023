fn check_row_empty(data: &Vec<Vec<char>>, row: usize) -> bool {
    data[row]
        .iter()
        .filter(|x| **x == '#')
        .collect::<Vec<&char>>()
        .is_empty()
}

fn check_column_empty(data: &Vec<Vec<char>>, column: usize) -> bool {
    for r in data {
        if r[column] == '#' {
            return false;
        }
    }
    true
}

fn get_empty_rows(data: &Vec<Vec<char>>) -> Vec<usize> {
    let mut empty_rows: Vec<usize> = Vec::new();
    for r in 0..data.len() {
        if check_row_empty(data, r) {
            empty_rows.push(r);
        }
    }
    empty_rows
}

fn get_empty_cols(data: &Vec<Vec<char>>) -> Vec<usize> {
    let mut empty_cols: Vec<usize> = Vec::new();
    for c in 0..data[0].len() {
        if check_column_empty(data, c) {
            empty_cols.push(c);
        }
    }
    empty_cols
}

fn get_shortest_path(pair: ((usize, usize), (usize, usize))) -> i128 {
    let abs_x = (pair.0 .0 as i128 - pair.1 .0 as i128).abs();
    let abs_y = (pair.0 .1 as i128 - pair.1 .1 as i128).abs();
    abs_x + abs_y
}

fn shortest_path(
    pair: ((usize, usize), (usize, usize)),
    empty_rows: Vec<usize>,
    empty_cols: Vec<usize>,
    e: i128,
) -> i128 {
    let without_expansion = get_shortest_path(pair);
    let (mut empty_rows_between, mut empty_cols_between) = (0, 0);
    for r in empty_rows {
        if (r > pair.0 .0 && r < pair.1 .0) || (r > pair.1 .0 && r < pair.0 .0) {
            empty_rows_between += 1;
        }
    }
    for c in empty_cols {
        if (c > pair.0 .1 && c < pair.1 .1) || (c > pair.1 .1 && c < pair.0 .1) {
            empty_cols_between += 1;
        }
    }
    without_expansion + empty_rows_between * (e - 1) + empty_cols_between * (e - 1)
}

fn total(data: &Vec<Vec<char>>, e: i128) -> i128 {
    let empty_rows = get_empty_rows(&data);
    let empty_cols = get_empty_cols(&data);

    let mut galaxies: Vec<(usize, usize)> = Vec::new();
    for r in 0..data.len() {
        for c in 0..data[0].len() {
            if data[r][c] == '#' {
                galaxies.push((r, c));
            }
        }
    }

    let mut total = 0;
    for g in galaxies.clone() {
        let other_galaxies: Vec<(usize, usize)> =
            galaxies.iter().filter(|x| **x != g).map(|x| *x).collect();
        for o in other_galaxies {
            let p = (g, o);
            total += shortest_path(p, empty_rows.clone(), empty_cols.clone(), e);
        }
    }

    total / 2 //divide by 2 because each pair will be double-counted
}

pub fn part_one(filename: &str) -> i128 {
    let data = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(|s| s.chars().collect())
        .collect::<Vec<Vec<char>>>();

    total(&data, 2)
}

pub fn part_two(filename: &str) -> i128 {
    let data = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(|x| x.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    total(&data, 1_000_000)
}
