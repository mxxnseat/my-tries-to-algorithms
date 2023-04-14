import { BinaryTree } from "../ds/binary-tree";

/*  https://leetcode.com/problems/generate-parentheses/
    Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

    
    Example 1:

    Input: n = 3
    Output: ["((()))","(()())","(())()","()(())","()()()"]

    Example 2:

    Input: n = 1
    Output: ["()"]
*/
function generateParenthesis(n: number): string[] {
  const tree = new BinaryTree({ p: "(", lpc: 1, rpc: 0 });
  const stack = [tree];
  const result: string[] = [];
  let h = n * 2;
  while (h > 0) {
    let stackLength = stack.length;
    const tmpStack: any = [];
    while (stackLength > 0) {
      const root = stack.pop() as BinaryTree;

      if (root.value.rpc < n && root.value.lpc > root.value.rpc) {
        root.right = new BinaryTree({
          p: root.value.p + ")",
          rpc: root.value.rpc + 1,
          lpc: root.value.lpc,
        });
        tmpStack.push(root.right);
      }
      if (root.value.lpc < n) {
        root.left = new BinaryTree({
          p: root.value.p + "(",
          rpc: root.value.rpc,
          lpc: root.value.lpc + 1,
        });
        tmpStack.push(root.left);
      }
      stackLength--;
    }
    stack.push(...tmpStack);
    h--;
  }
  let current: BinaryTree | null = tree;
  while (current || stack.length) {
    while (current !== null) {
      stack.push(current);
      current = current.left;
    }
    current = stack[stack.length - 1];
    stack.pop();
    if (!current.left && !current.right) {
      result.push(current.value.p);
    }
    current = current.right;
  }
  return result;
}

console.log(generateParenthesis(3));
