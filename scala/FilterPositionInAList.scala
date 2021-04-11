object FilterPositionInAList {
  def f(arr: List[Int]): List[Int] = {
    def go(xs: List[Int], s: Boolean): List[Int] = {
      xs match {
        case Nil => Nil
        case (x::xs) if s => x::go(xs, !s)
        case (x::xs) if !s => go(xs, !s)
      }
    }
    go(arr, false)
  }
}