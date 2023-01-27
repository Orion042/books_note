import os

def createDirectory():
    path = os.getcwd()
    
    directorys = ["md_files","notes"]

    for name in directorys:
        if not (os.path.isdir(path + "\\" + name)):
            print(name + " directory does not exist.")
            os.mkdir(path + "\\" + name)
            print(name + " created !!")

if __name__ == "__main__":
    createDirectory()