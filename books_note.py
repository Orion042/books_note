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

def search_book_title(lib):
    book_title = input("検索したい書籍名を入力 : ")

    lib.search_bookTitle_db.restype = c_char_p

    result = lib.search_bookTitle_db(book_title.encode()).decode("utf-8")

    if result == "None":
        sys.exit()

    start_select = input(book_title + "を表示しますか[Y/N] : ")

    if ((start_select.lower() == "yes") or (start_select.lower() == "y")):
        py_files.start_web.start_html(result)
    else:
        sys.exit()

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
        search_book_title(lib)

    elif user_select == 2:
        # 著者名検索
        book_author = input("検索したい著者名を入力 : ")

        lib.search_author_db.restype = c_char_p

        result = lib.search_author_db(book_author.encode()).decode("utf-8")

        if result == "None":
            sys.exit()

        print("======================")

        search_book_title(lib)

    elif user_select == 3:
        # タグ検索
        lib.show_all_tags()

        book_tag = input("検索したいタグ名を入力(半角カンマで区切る) : ")

        lib.search_tags_db.restype = c_char_p

        result = lib.search_tags_db(book_tag.encode())

        if result == "None":
            sys.exit()
        else:
            search_book_title(lib)

    elif user_select == 4:
        # 書籍全表示
        lib.show_all_db()

    elif user_select == 5:
        # 追加
        file_name = input("htmlファイル名を入力 : ")

        file_result, md_file = py_files.find_files.find_htmlFile(file_name)

        if (file_result):
            book_title, book_author, book_tags = read_info(md_file[0])

            lib.insert(book_title.encode(), file_name.encode(), book_author.encode(), book_tags.encode())

        else:
            sys.exit()
        
    elif user_select == 6:
        # 更新
        pass
    else:
        print("入力エラー")
        sys.exit()

if __name__ == "__main__":
    main()