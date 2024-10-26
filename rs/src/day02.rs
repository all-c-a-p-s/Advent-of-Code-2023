struct Game {
    id: i32,
    rounds: Vec<Vec<i32>>,
}

fn parse_round(round: &str) -> Vec<i32> {
    let mut v = vec![0, 0, 0];
    let colours = round
        .split(',')
        .map(|x| x.split_whitespace().collect())
        .collect::<Vec<Vec<&str>>>();
    for c in colours {
        match c[1] {
            "red" => v[0] = c[0].to_string().parse().unwrap(),
            "green" => v[1] = c[0].to_string().parse().unwrap(),
            "blue" => v[2] = c[0].to_string().parse().unwrap(),
            x => panic!("invalid colour {}", x),
        }
    }
    v
}

fn highest_nth(v: Vec<Vec<i32>>, n: usize) -> i32 {
    let mut max = 0;
    for u in v {
        max = std::cmp::max(u[n], max)
    }
    max
}

impl Game {
    fn from_line(line: String) -> Self {
        let before_colon = &line[0..line.chars().position(|x| x == ':').unwrap()];
        let after_colon = &line[line.chars().position(|x| x == ':').unwrap() + 1..];
        let id: i32 = before_colon
            .to_string()
            .chars()
            .filter(|x| x.is_ascii_digit())
            .collect::<String>()
            .parse()
            .expect("failed to get id");
        let rounds: Vec<Vec<i32>> = after_colon.split(';').map(parse_round).collect();
        Self { id, rounds }
    }

    fn line_score_part_one(&self) -> i32 {
        let max_red = highest_nth(self.rounds.clone(), 0);
        let max_green = highest_nth(self.rounds.clone(), 1);
        let max_blue = highest_nth(self.rounds.clone(), 2);

        if max_red <= 12 && max_green <= 13 && max_blue <= 14 {
            self.id
        } else {
            0
        }
    }

    fn line_score_part_two(&self) -> i32 {
        let max_red = highest_nth(self.rounds.clone(), 0);
        let max_green = highest_nth(self.rounds.clone(), 1);
        let max_blue = highest_nth(self.rounds.clone(), 2);

        max_red * max_green * max_blue
    }
}

pub fn part_one(filename: &str) -> i32 {
    let ans: i32 = std::fs::read_to_string(filename)
        .expect("failed to read file")
        .lines()
        .map(|l| Game::from_line(l.to_string()))
        .map(|g| g.line_score_part_one())
        .sum();
    ans
}

pub fn part_two(filename: &str) -> i32 {
    let ans: i32 = std::fs::read_to_string(filename)
        .expect("failed to read file")
        .lines()
        .map(|l| Game::from_line(l.to_string()))
        .map(|g| g.line_score_part_two())
        .sum();
    ans
}
