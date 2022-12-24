from ctypes import *
import ctypes

def check():
    lib = cdll.LoadLibrary("./main/go_sqlite.so")
    lib.check_go()

def main():
    lib = cdll.LoadLibrary("./main/go_sqlite.so")
    lib.test_sql()

if __name__ == "__main__":
    main()