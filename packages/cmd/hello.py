import sys

def hello(name):
    return f"Hello, {name}!"

if __name__ == "__main__":
    name = sys.argv[1]
    print(hello(name))
