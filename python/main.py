from fastapi import FastAPI
import uvicorn

app = FastAPI()


def quickSort(arr):
    if len(arr) <= 1:
        return arr

    stack = [(0, len(arr) - 1)]

    while stack:
        low, high = stack.pop()

        if low >= high:
            continue

        pivotIndex = partition(arr, low, high)

        stack.append((low, pivotIndex - 1))
        stack.append((pivotIndex + 1, high))


def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1

    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]

    arr[i + 1], arr[high] = arr[high], arr[i + 1]

    return i + 1


@app.post("/sort")
def sort(arr: list[int]):
    quickSort(arr)
    return arr
