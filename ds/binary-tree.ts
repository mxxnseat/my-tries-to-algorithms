export class BinaryTree<T = any> {
  constructor(
    public value: T,
    public left: BinaryTree<T> | null = null,
    public right: BinaryTree<T> | null = null
  ) {}
}
