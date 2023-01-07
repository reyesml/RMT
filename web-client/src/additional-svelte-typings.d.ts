declare namespace svelteHTML {
  // enhance attributes
  interface HTMLAttributes<T> {
    'on:outclick'?: (event: any) => any;
  }
}