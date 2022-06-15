object Solution {
  def f (arr: List[Int]) : List[Int] = {
    arr match {
      case (x :: xs) => x.abs :: (f(xs))
      case _ => Nil
    }
  }
}
