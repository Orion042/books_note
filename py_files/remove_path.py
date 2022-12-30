import os

def removed_pass(file_path):
    path = os.getcwd() + "\\md_files\\"

    file_name = file_path.replace(path,"").replace(".md","")

    return file_name

if __name__ == "__main__":
    removed_pass()