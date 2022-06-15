sealed trait Tree[+A]
case object E extends Tree[Nothing]
case class Node[A](r: A, lt: Tree[A], rt: Tree[A]) extends Tree[A]

object Tree {
  def traverseInOrder[A](t: Tree[A]): List[A] = {
    t match {
      case E             => Nil
      case Node(x, l, r) => traverseInOrder(l) ++ List(x) ++ traverseInOrder(r)
    }
  }

  def swapAt[A](tree: Tree[A], d: Int): Tree[A] = {
    (tree, d) match {
      case (E, _)             => E
      case (Node(v, l, r), 1) => Node(v, r, l)
      case (Node(v, l, r), d) => Node(v, swapAt(l, d - 1), swapAt(r, d - 1))
    }
  }
}

object BTree {

  def convertToInt(nodes: List[String]): List[List[Int]] = {
    nodes.map(_.split(" ").toList.map(_.toInt))
  }

  def buildTree(
      nodes: List[List[Int]],
      t: Tree[Int] = E,
      buff: List[Int] = List.empty
  ): Tree[Int] = {
    (nodes, t, buff) match {
      case (Nil, Node(v, l, r), _) => Node(v, l, r)
      case (nodes, E, Nil)         => buildTree(nodes, makeSingletonTree(1), List(1))
      case (n :: ns, t, b :: bs) =>
        buildTree(
          ns,
          appendT(t, makeSingletonTree(b), n),
          (bs ++ n).filter(_ != -1)
        )
    }
  }

  def appendT(tree: Tree[Int], bt: Tree[Int], e: List[Int]): Tree[Int] = {
    (tree, bt, e) match {
      case (_, E, _) => E
      case (E, _, _) => E
      case (Node(v, l, r), Node(v1, l1, r1), e)
          if (v == v1) && (l == l1) && (r == r1) =>
        joinT(Node(v, l, r), makeSingletonTree(e.head), makeSingletonTree(e(1)))
      case (Node(v, l, r), bt, e) =>
        Node(v, appendT(l, bt, e), appendT(r, bt, e))
    }
  }

  def joinT[A](t1: Tree[A], leafL: Tree[A], leafR: Tree[A]): Tree[A] = {
    (t1, leafL, leafR) match {
      case (E, _, _)               => E
      case (Node(a, _, _), le, ri) => Node(a, le, ri)
    }
  }

  def makeSingletonTree(node: Int): Tree[Int] = {
    node match {
      case -1 => E
      case _  => Node(node, E, E)
    }
  }

  def chainSwap[A](tree: Tree[A], h: Int, c: Int = 1): Tree[A] = {
    (tree) match {
      case E                               => E
      case t if t == Tree.swapAt(t, h * c) => t
      case t                               => chainSwap(Tree.swapAt(t, h * c), h, c + 1)
    }
  }

  def solve[A](tree: Tree[A], h: List[Int]): List[List[A]] = {
    (tree, h) match {
      case (E, _)   => Nil
      case (_, Nil) => Nil
      case (t, x :: xs) => {
        val temp = chainSwap(t, x)
        Tree.traverseInOrder(temp) :: solve(temp, xs)
      }
    }
  }

  def main(args: Array[String]): Unit = {
    val input = io.Source.stdin.getLines().toList.tail
    val inputAsIntegers = convertToInt(input)
    val nodes = inputAsIntegers.takeWhile(_.length > 1)
    val swapOp = inputAsIntegers.dropWhile(_.length > 1).flatten.tail
    val tree = buildTree(nodes, E, Nil)
    val solution = solve(tree, swapOp)
    solution.foreach(x => println(x.mkString(" ")))
  }
}
