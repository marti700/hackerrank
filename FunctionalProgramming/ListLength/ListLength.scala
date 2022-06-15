object Solution {
  def f (arr: List[Int]) : Int = arr.foldLeft(0) {(acc, e) => acc +1}
}