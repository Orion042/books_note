from ctypes import *
import ctypes
from py_files import find_same_files
import os
import sys

def check():
    lib = cdll.LoadLibrary("./main/go_sqlite.so")
    lib.check_go()

def remove_comment(comment):
    removed_comment = comment.lstrip("<!-- ").rstrip(" -->\n")

    return removed_comment

def read_info(file_list):
    book_title = ""
    book_author = ""
    book_tag = []

    with open(file_list) as f:
        index = 0
        for book_info in f:
            renamed_info = remove_comment(book_info)
            print(renamed_info)

            if index == 2:
                book_title = renamed_info.lstrip("book_title: ")
            elif index == 3:
                book_author = renamed_info.lstrip("book_author: ")
            elif index == 4:
                book_tag = renamed_info.lstrip("book_tags: ")
            index += 1

            if index >= 5:
                break
    
    print(book_title)
    print(book_author)
    print(book_tag)

    return book_title, book_author, book_tag

def main():
    get_same_files = find_same_files.find_same_files()

    for i in range(len(get_same_files)):
        book_title, book_author, book_tag = read_info(get_same_files[i])

    sys.exit()

    lib = cdll.LoadLibrary("./main/go_sqlite.so")
    lib.test_sql()

if __name__ == "__main__":
    main()