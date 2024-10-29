use std::cmp::{max, min};

#[derive(Clone, Copy, PartialEq)]
struct Range {
    min: i128,
    max: i128,
}

struct Map {
    source: Range,
    destination: Range,
}

fn input_seeds(start: usize, end: usize, lines: &[String]) -> Vec<i128> {
    let mut seeds = String::new();
    for i in start..=end {
        seeds.push_str(&lines[i]);
    }

    let nums_str: Vec<&str> = seeds[7..].split_whitespace().collect();
    nums_str
        .iter()
        .map(|&num| num.parse::<i128>().unwrap())
        .collect()
}

fn read_map(start: usize, end: usize, lines: &[String]) -> Vec<Map> {
    let mut maps = Vec::new();
    let mut num_lists = Vec::new();

    for i in start..=end {
        let str = &lines[i];
        let ns: Vec<&str> = str.split_whitespace().collect();
        let nums: Vec<i128> = ns.iter().map(|&n| n.parse::<i128>().unwrap()).collect();
        num_lists.push(nums);
    }

    for list in num_lists {
        let destination = Range {
            min: list[0],
            max: list[0] + list[2] - 1,
        };
        let source = Range {
            min: list[1],
            max: list[1] + list[2] - 1,
        };
        maps.push(Map {
            source,
            destination,
        });
    }

    maps
}

pub fn part_one(filename: &str) -> i128 {
    let lines: Vec<String> = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut maps: Vec<Vec<Map>> = Vec::new();
    let mut first_map;
    let mut seeds = Vec::new();

    for (i, line) in lines.iter().enumerate() {
        if line.is_empty() {
            continue;
        }
        if line == "seed-to-soil map:" {
            first_map = i;
            seeds = input_seeds(0, first_map - 2, &lines);
        }
        if line.ends_with("map:") {
            let mut end = i + 1;
            while end < lines.len() && !lines[end].is_empty() {
                end += 1;
            }
            maps.push(read_map(i + 1, end - 1, &lines));
        }
    }

    let mut location = Vec::new();

    for &seed in &seeds {
        let mut key = seed;
        for map_layer in &maps {
            for m in map_layer {
                let dif = m.destination.min - m.source.min;
                if key >= m.source.min && key <= m.source.max {
                    key += dif;
                    break;
                }
            }
        }
        location.push(key);
    }

    *location.iter().min().unwrap()
}

fn find_intersection(r1: Range, r2: Range) -> Range {
    Range {
        min: max(r1.min, r2.min),
        max: min(r1.max, r2.max),
    }
}

fn subtract_range(r1: Range, r2: Range) -> Vec<Range> {
    if r1 == find_intersection(r1, r2) {
        return vec![];
    }

    let mut subtracted = Vec::new();
    let intersection = find_intersection(r1, r2);

    if intersection.min < intersection.max {
        if r1.min < r2.min {
            subtracted.push(Range {
                min: r1.min,
                max: r2.min - 1,
            });
        }
        if r1.max > r2.max {
            subtracted.push(Range {
                min: r2.max + 1,
                max: r1.max,
            });
        }
    } else {
        subtracted.push(r1);
    }

    subtracted
}

pub fn part_two(filename: &str) -> i128 {
    let lines: Vec<String> = std::fs::read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut maps: Vec<Vec<Map>> = Vec::new();
    let mut first_map;
    let mut seeds = Vec::new();

    for (i, line) in lines.iter().enumerate() {
        if line.is_empty() {
            continue;
        }
        if line == "seed-to-soil map:" {
            first_map = i;
            seeds = input_seeds(0, first_map - 2, &lines);
        }
        if line.ends_with("map:") {
            let mut end = i + 1;
            while end < lines.len() && !lines[end].is_empty() {
                end += 1;
            }
            maps.push(read_map(i + 1, end - 1, &lines));
        }
    }

    let mut location = Vec::new();
    let mut seed_ranges = Vec::new();

    for i in (0..seeds.len() - 1).step_by(2) {
        seed_ranges.push(Range {
            min: seeds[i],
            max: seeds[i] + seeds[i + 1] - 1,
        });
    }

    let mut location_ranges = Vec::new();

    for sr in seed_ranges {
        let mut input_ranges = vec![sr];

        for map_layer in &maps {
            let mut output_ranges = Vec::new();
            let mut unmapped = input_ranges.clone();

            for m in map_layer {
                let dif = m.destination.min - m.source.min;
                let mut i = 0;
                while i < unmapped.len() {
                    let current = unmapped[i];
                    let intersection = find_intersection(current, m.source);

                    if intersection.min <= intersection.max && intersection.max > 0 {
                        let subtracted = subtract_range(current, m.source);
                        unmapped.remove(i);
                        unmapped.extend(subtracted);
                        output_ranges.push(Range {
                            min: intersection.min + dif,
                            max: intersection.max + dif,
                        });
                    } else {
                        i += 1;
                    }

                    if unmapped.is_empty() {
                        break;
                    }
                }
            }

            input_ranges = output_ranges;
        }

        location_ranges.extend(input_ranges);
    }

    for r in location_ranges {
        location.push(r.min);
    }

    *location.iter().min().unwrap()
}
