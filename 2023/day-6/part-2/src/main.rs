use std::iter::zip;

fn main() {
    let input = include_str!("input.txt");

    let result = process(input);

    println!("{}", result);
}

fn process(input: &str) -> usize {
    let mut times: Vec<usize> = vec![];
    let mut distances: Vec<usize> = vec![];

    let mut curr_num: usize = 0;
    for c in input.chars() {
        if c == ':' {
            if curr_num > 0 {
                    times.push(curr_num);
                curr_num = 0;
            }
            continue;
        }

        if c.is_digit(10) {
            curr_num *= 10;
            curr_num += c.to_digit(10).unwrap() as usize;
            continue;
        }
    }

    distances.push(curr_num);

    let mut total = 1;
    for (race_time, record_distance) in zip(times, distances) {
        let mut count = 0;

        for held_down in 1..race_time {
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

        let expected: usize = 71503;

        assert_eq!(result, expected)
    }
}
