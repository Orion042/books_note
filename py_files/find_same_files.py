import os
import glob
import shutil

def get_file_path(file_name):
    path = os.getcwd()
    md_path = path + "\\" + file_name + "\\"

    return md_path

def find_same_files():
    path = get_file_path("md_files")
    
    html_files = glob.glob(path + "*.html")

    md_file_list = []

    for i in range(len(html_files)):
        files = html_files[i].replace(path,"").replace(".html","") + ".md"

        md_file_list.extend(glob.glob(path + files))

    for i in range(len(md_file_list)):
        # shutil.move(html_files[i], get_file_path("notes"))
        pass

    return md_file_list

if __name__ == "__main__":
    find_same_files()