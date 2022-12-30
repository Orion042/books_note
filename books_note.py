from ctypes import *
import py_files
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
    book_tags = []

    with open(file_list,encoding="utf-8") as f:
        index = 0
        for book_info in f:
            renamed_info = remove_comment(book_info)
            print(renamed_info)

            if index == 2:
                book_title = renamed_info.lstrip("book_title: ")
            elif index == 3:
                # book_author = renamed_info.lstrip("book_author: ") これだと著者名が消える
                book_author = renamed_info.lstrip("book_author:").strip(" ")
            elif index == 4:
                book_tags = renamed_info.lstrip("book_tags: ")

            index += 1

            if index >= 5:
                break

    return book_title, book_author, book_tags

def main():

    # goファイル読み込み
    lib = cdll.LoadLibrary("./main/go_sqlite.so")

    # 前処理(ファイル自動移動 and データベース保存)
    get_same_files = py_files.find_same_files.find_same_files()

    if (len(get_same_files) >= 1):
        for i in range(len(get_same_files)):
            book_title, book_author, book_tags = read_info(get_same_files[i])

            file_name = py_files.remove_path.removed_pass(get_same_files[i])

            lib.preprocessing_sql(book_title.encode(), file_name.encode(), book_author.encode(), book_tags.encode())
    else:
        lib.connect()


    # ユーザ選択開始
    user_select = py_files.user_select.select()

    if user_select == 9:
        print("終了")
        sys.exit()
    elif user_select == 1:
        # 書籍名検索
        book_name = input("検索したい書籍名を入力 : ")
        lib.read_db(book_name.encode())
    elif user_select == 2:
        # 著者名検索
        pass
    elif user_select == 3:
        # タグ検索
        pass
    elif user_select == 4:
        # 更新
        pass
    else:
        print("入力エラー")
        sys.exit()

if __name__ == "__main__":
    main()