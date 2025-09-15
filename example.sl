set min = fn(x, y) {
    if (x > y) {
        return y;
    }
    return x;
};

set a = 10;
print("set a = ", a);

set b = 5;
print("set b = ", b);

print("min(a, b) = ", min(a, b));

set ok = true;
print("set ok = ", ok);

print("!ok = ", !ok);
print("-a = ", -a);

print("a + b = ", a + b);
print("a - b = ", a - b);
print("a * b = ", a * b);
print("a / b = ", a / b);

set str = "slow";
print("str + '-lang' = ", str + "-lang");
