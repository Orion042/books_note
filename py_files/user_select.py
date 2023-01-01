
def select():
    user = -1

    while(True):
        print("================")
        print("1 : 書籍名検索")
        print("2 : 著者名検索")
        print("3 : タグ検索")
        print("4 : 書籍名全表示")
        print("5 : 追加")
        print("6 : 更新")
        print("9 : 終了")
        print("================")
        
        user = int(input("--> "))

        if ((user == 1) or (user == 2) or (user == 3) or (user == 4) or (user == 5) or (user == 6) or (user == 9)):
            break
    return user

if __name__ == "__main__":
    select()