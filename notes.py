def notes():
    run = True
    while run:
        print("press q to exit\n")
        choice = str(input("enter word: "))

        print(choice + "\n")
        if choice == "q":
            run = False


if __name__ == "__main__":
    notes()
