def add(task):
    splitTask = task.split("+")
    partOne = splitTask[0]
    partTwo = splitTask[1]

    result = int(partTwo) + int(partOne)

    return result


def sub(task):
    splitTask = task.split("-")
    partOne = splitTask[0]
    partTwo = splitTask[1]

    result = int(partOne) - int(partTwo)

    return result


res = add("2+2")
print("Result:", res)

res = sub("4-2")
print("Result:", res)