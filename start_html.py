import webbrowser
import os

# url_input = input("--> ")

url_input = "test.html"

url = os.getcwd() + "\\notes\\" + url_input

firefox = webbrowser.Mozilla("C:\\Program Files\\Mozilla Firefox\\firefox.exe")

firefox.open(url)
