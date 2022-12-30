import webbrowser
import os

def start_html(file_name):
    file_name = file_name + ".html"

    file_path = os.getcwd() + "\\notes\\" + file_name

    print(file_path)

    firefox = webbrowser.Mozilla("C:\\Program Files\\Mozilla Firefox\\firefox.exe")

    firefox.open(file_path)


if __name__ == "__main__":
    start_html()