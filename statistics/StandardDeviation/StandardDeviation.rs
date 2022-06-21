use std::io;

fn stdev (n: Vec<f32>) -> f32 {
  let e_sum: f32 = n.iter().sum();
  let mean: f32 = e_sum / n.len() as f32;
  let mut variance: f32 = 0.0;

  for i in n.iter() {
    variance = variance + ( (i-mean) * (i-mean) )
  }

  variance = variance / (n.len() as f32);

  return variance.sqrt();

}

fn main() {
    let mut vec_size: String = String::new();
    io::stdin().read_line(&mut vec_size).unwrap();
    let mut _tc: i32 = vec_size.trim().parse().unwrap();

    let mut numbers: String = String::new();
    io::stdin().read_line(&mut numbers).unwrap();
    let numbers_arr: Vec<f32> = numbers
        .split(" ")
        .map(|n| n.trim().parse().unwrap())
        .collect();

    println!("{:.1}", stdev(numbers_arr));
}