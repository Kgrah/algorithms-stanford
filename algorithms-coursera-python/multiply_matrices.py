import numpy as np

def main():
    a, b = [], []
    multiply_matrices(a, b)

def multiply_matrices(a, b):
    aRows, aCols = len(a), len(a[0])
    bRows, bCols = len(b), len(b[0])

    if aRows != bRows:
        return None
    
    result = [[0 for _ in range(bCols)] for _ in range(aRows)]
    
    for i in range(aRows):
        for j in range(bCols):
            for k in range(aCols):
                result[i][j] += a[i][k] * b[k][j]

if __name__ == "__main__":
    main()