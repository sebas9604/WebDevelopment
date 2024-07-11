#!/bin/python3

import math
import os
import random
import re
import sys



#
# Complete the 'twoSums' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. INTEGER_ARRAY A
#  2. INTEGER K
#

def twoSums(A, K):
    # Write your code here
    for x in A:
        for y in A:
            print(x,y)
            
            
            

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    A_count = int(input().strip())

    A = []

    for _ in range(A_count):
        A_item = int(input().strip())
        A.append(A_item)

    K = int(input().strip())

    result = twoSums(A, K)

    fptr.write(result + '\n')

    fptr.close()