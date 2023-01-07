declare namespace svelteHTML {
  interface HTMLAttributes<T> {
    'on:outclick'?: (event: any) => any;
  }
}