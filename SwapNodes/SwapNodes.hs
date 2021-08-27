-- Swap  Nodes
--
-- Tree data structure
data Tree a = E | Node a (Tree a) (Tree a) deriving (Show, Read, Eq)

-- In order traversal algorithm
traverseInOrder :: Tree a -> [a]
traverseInOrder E = []
traverseInOrder (Node a l r) = traverseInOrder l ++ [a] ++ traverseInOrder r

-- Swaps the left subtree with the right one at a given depth
-- if the specified deth is greather than the actual depth of the tree
-- the same tree will be returned
swapNodeAt :: Tree a -> Integer -> Tree a
swapNodeAt E _ = E
swapNodeAt (Node a l r) 1 = Node a r l
swapNodeAt (Node a l r) d = Node a (swapNodeAt l (d-1)) (swapNodeAt r (d-1))

-- Creates a tree which node is an integer and the left and right subtrees are empty nodes (E)
makeSingletonTree :: Integer -> Tree Integer
makeSingletonTree (-1) = E
makeSingletonTree x = Node x E E

-- given tree trees join them together
-- the first tree in the list of arguments will be the root the second the left sub tree and the thrid the right one
joinT :: Tree a -> Tree a -> Tree a -> Tree a
joinT E _ _ = E
joinT (Node a left right) leftLeaf rightLeaf = Node a leftLeaf  rightLeaf

-- makes n consecutive node swaps. The number of swaps is the same as the lenght of the second function argument
swapNodesAtHeights :: (Eq a) => Tree a -> [Integer] -> Tree a
swapNodesAtHeights tree [] = tree
swapNodesAtHeights tree (x:xs)
  | tree == swapNodeAt tree x = tree
  | otherwise = swapNodesAtHeights (swapNodeAt tree x) xs


-- IO TREATEMENT
-- builds a tree from a list of integers (in the format specified by the problem statement)
-- the first parameter is the tree to be constructed (the first time this function must be called with E as this parameter)
-- the Second argument is the list of Nodes of the tree
-- the third argument is .... well I call it the buffer
-- This functions searches for the last added node (the head of the buffer) and then adds the head of the list of nodes
-- (the second function argument) to it. Note that the null nodes are not included in the buffer.
buildTree :: Tree Integer -> [[Integer]] -> [Integer] -> Tree Integer
buildTree E xs [] = buildTree (makeSingletonTree 1) xs [1]
buildTree (Node a left right) [] _ = Node a left right
buildTree (Node a left right) (x:xs) buff =
  buildTree (appendT (Node a left right) (makeSingletonTree $ head buff) x) xs (tail buff ++ ignoreNull [head x, x !! 1])
  where ignoreNull = filter (/= -1) -- ignores the null trees (null trees are represented by -1 in the inpunt)

-- Appends a tree to a Node
-- The first function argument is the tree we want to appent the new node to (this is an entire tree)
-- The second argument is the node to which we want to append the node to (this is just a node which is part of the three, the one in the first arg)
-- the third argument is the nodes we want to add (in the format of the input specified in the problem statement ([int, int]))
-- this function search recursively the tree  (the one in the first arg) until it finds the node to which we want to append the new nodes to (the second arg)
-- and then converts the thrid argument numbers to singleton trees (the third argument should be a two element integer array, maybe i should have used a tuple)
-- and applied to the found node.
appendT :: Tree Integer -> Tree Integer -> [Integer] -> Tree Integer
appendT E _ _= E
appendT (Node a left right) (Node b left1 right1) strN
  | (a == b) && (left == left1) && (right == right1) = joinT (Node a left right) (makeSingletonTree (head strN)) (makeSingletonTree (strN !! 1))
  | otherwise = Node a (appendT left (Node b left1 right1) strN) (appendT right (Node b left1 right1) strN)

-- converts a list of strings (in the format specified in the problem statement) to a list of integers of integers
convertToIntList :: [String] -> [[Integer]]
convertToIntList = map (map (read :: String -> Integer) . words)

-- reads n consecutive lines of stdin
readLines :: Int -> IO [String]
readLines 0 = return []
readLines x = do
  line <- getLine
  rest <- readLines (x-1)
  return (line : rest)

-- applies builds the problem solution
solve :: Tree Integer -> [Integer] -> [[Integer]]
solve t [] = []
solve t (x:xs) =  traverseInOrder swapT : solve swapT xs
  where swapT = swapNodesAtHeights t (heights x) -- swapT is an alias for swapNodesAtHeights
        heights x = [x*j | j <- [1..] ] -- get the heights on wich the swaps will happend

-- Formtas the output to match hackerrank expected output
formatOutput :: [Integer] -> String
formatOutput = foldl (\acc x -> if acc == "" then  acc ++ show x else acc ++ " " ++ show x) ""



main :: IO ()
main = do
  numberOfNodes <-  readLn :: IO Int

  --Handle the tree nodes part of the input
  numberOfTreeNodes <- readLines numberOfNodes
  let treeNodes = convertToIntList numberOfTreeNodes
  numberOfSwapOperations <- readLn :: IO Int

  -- Handle the swap operation part of the input
  swapOpDepth <- readLines numberOfSwapOperations
  let swapOperations = convertToIntList swapOpDepth


  -- Builds the tree
  let tree = buildTree E treeNodes []
  -- Applies the solution
  let x =  solve tree (concat swapOperations)
  -- print to stdout
  --convert to string
  let r = map (unlines .map show) x
  -- mapM_ (putStrLn . formatOutput) x
  print tree