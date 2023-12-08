use std::iter::zip;

fn main() {
    let input = include_str!("input.txt");

    let result = process(input);

    println!("{}", result);
}

pub fn process(input: &str) -> usize {
    let mut times: Vec<usize> = vec![];
    let mut distances: Vec<usize> = vec![];

    for (i, c) in input.split_whitespace().enumerate() {
        if i == 0 || i == 5 {
            continue;
        }

        let curr_num: usize = c.parse().unwrap();
        if i > 4 {
            distances.push(curr_num);
        } else {
            times.push(curr_num);
        }
    }

    let mut total = 1;
    for (race_time, record_distance) in zip(times, distances) {
        let mut count = 0;

        for held_down in 1..race_time - 1 {
            let time_remaining = race_time - held_down;
            let distance = held_down * time_remaining;

            if distance > record_distance {
                count += 1;
            } else if count > 0 {
                break;
            }
        }

        total *= count;
    }

    return total;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse() {
        let input = "Time:      7  15   30
Distance:  9  40  200
";

        let result = process(input);

        let expected: usize = 288;

        assert_eq!(result, expected)
    }
}
