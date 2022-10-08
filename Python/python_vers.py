from random import choice
from string import ascii_lowercase as al
from time import time
from os.path import exists
from os import remove

"""
Function: parses 5m lines at 49 char a piece, splitting them by the delimiter (,) and then writing to a list (end result, 10m lines)

In terms of ease, this one was by far easiest to create, it also seems pretty fast at 
reading files, usually seeing speeds of around 2 seconds (at the highest, 4 seconds).

You could probably make this even faster but I wanted to keep things very simple for when I was testing other languages.

"""


def random_string():
    """Generates a random string given the size (in this case 24)"""

    def gen():
        return "".join(choice(al) for i in range(24))  # joins a list of randomly selected chars from alphabet 24 times

    return f"{gen()},{gen()}"  # returns 2 of the gen strings as a delimited string


def generate_data():
    t1 = time()

    print("Generating test data, this might take a little while")
    lines = []
    while len(lines) != 5000000:
        lines.append(random_string())

    t2 = time()
    print(f"Done, duration: {round(t2-t1, 2)}\n")
    return lines


def write_file(lines):
    """Writes specified list (lines) to file (test_data-py.txt)"""
    t1 = time()

    print("Writing data to test_data-py.txt")
    with open("test_data-py.txt", "w", encoding="utf-8") as f:
        f.write("\n".join(lines))

    t2 = time()
    print(f"Done, duration: {round(t2-t1, 2)}\n")


def read_file():
    """Reads test_data-py.txt and tests how fast it can parse all the lines"""
    t1 = time()

    print("Reading and parsing data from test_data-py")
    with open("test_data-py.txt") as f:  # opens file as ITERATOR
        lines = []
        for data in f:  # runs through it in a loop, better on memory usage this way
            data = data.strip().split(",") # Parse delimited line
            lines.append(data[0])  # Add both strings to line list, should generate 10m lines
            lines.append(data[1])

    t2 = time()
    print(f"Done, duration: {round(t2-t1, 2)}\n")


def main():
    """Handles function calling"""
    start = time()

    data = generate_data()
    write_file(data)
    read_file()

    end = time()
    print(f"TOTAL DURATION: {round(end-start, 2)}")


if __name__ == "__main__":
    main()
