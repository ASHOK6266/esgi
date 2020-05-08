from printEvenNumbers import check_even

print("Please enter a list of numbers to check it odd or even or E to (E)xit of the screen ")
while True:
    list1 = str(input("Enter list of numbers "))
    if list1 == "E":
        print("Bye")
        break
    result = check_even(list1)
    print("The numbers are : " + result)