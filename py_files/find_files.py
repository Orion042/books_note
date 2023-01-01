import os
import glob

def find_htmlFile(file_name):
    path = os.getcwd()
    html_path = path + "\\notes\\"
    md_path = path + "\\md_files\\"

    html_files = glob.glob(html_path + file_name + ".html")

    if len(html_files) == 1:
        md_file = glob.glob(md_path + file_name + ".md")
        if len(md_file) == 1:
            return True, md_file
        else:
            print("mdファイルがありません")
            return False, False
    elif len(html_files) > 1:
        print("ファイルが重複しています")
        return False, False
    else:
        print("ファイルがありません")
        return False, False


if __name__ == "__main__":
    find_htmlFile()