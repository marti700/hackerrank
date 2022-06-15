object Solution {
  def f(arr: List[Int]) : Int = arr.foldLeft(0) {
    (acc, e) => {
      if (e % 2 != 0) acc + e
      else acc
    }
  }
}