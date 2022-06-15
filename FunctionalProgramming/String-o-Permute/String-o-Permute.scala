object Solution {
  def stringPermute(str: List[Char]) : List[Char]= {
    str match {
      case Nil => Nil
      case (x::y::rest) => y :: x :: stringPermute(rest)
    }

    //the below recursive solution do not pass the last test case because of a timeout apparently it is slow
    // to test it out change the str parameter to string and add second parameter result to the method signature
    // if (str == "") return result
    // stringPermute(str.tail.tail, result + str.tail.head.toString + str.head.toString)
  }

  def main (args: Array[String]) : Unit = {
    val input = io.Source.stdin.getLines().toList.tail
    input.foreach(x => println(stringPermute(x.toList).mkString("")))
  }
}