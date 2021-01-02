function generateTheString(n: number): string {
    return "a".repeat(n-1) + ((n%2) == 0 ? "b" : "a")
};
