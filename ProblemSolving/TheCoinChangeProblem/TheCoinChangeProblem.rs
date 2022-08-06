use std::{collections::HashMap, io};

fn get_ways(amt: i128, denominations: &Vec<i128>, dp: &mut HashMap<String, i128>) -> i128 {
    let key: String = amt.to_string() + &"," + &denominations.len().to_string();

    if dp.contains_key(&key) {
        return *dp.get(&key).unwrap();
    }
    if amt == 0 {
        return 1;
    }
    if amt < 0 || denominations.len() == 0 {
        return 0;
    }

    let ways: i128 = get_ways(amt - denominations[0], &denominations, dp)
        + get_ways(amt, &denominations[1..].to_vec(), dp);
    dp.insert(key, ways);
    return ways;
}

fn main() {
    let mut fli: String = String::new();
    io::stdin().read_line(&mut fli).unwrap();
    let elements: Vec<&str> = fli.split(' ').collect();
    // get the target
    let n: i128 = elements[0].trim().parse().unwrap();
    //get second line of input a.k.a the coins
    let mut co: String = String::new();
    io::stdin().read_line(&mut co).unwrap();

    let str_coins: Vec<&str> = co.split(" ").collect();
    let mut coins: Vec<i128> = Vec::new();

    // convert coins to a vector of integers
    for c in str_coins {
      let number: i128 = c.trim().parse().unwrap();
      coins.push(number);
    }

    println!("{}", get_ways(n, &coins, &mut HashMap::new()));
}
