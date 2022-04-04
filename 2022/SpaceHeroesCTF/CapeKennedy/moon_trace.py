import sys

def main():
  if len(sys.argv) != 2:   
    print("Invalid args")
    return

  password = sys.argv[1]
  builder = 0
  
  for c in password:
    builder += ord(c)

  #add this to track the size
  print("length = ", len(password)) 

  #add this to track the variable   
  print("builder = ", str(builder))
  
  print("[6]= ", password[6], "[7] = ", password[7])
  
  if builder == 713 and len(password) == 8 and (ord(password[2]) == ord(password[5])):
    if (ord(password[3]) == ord(password[4])) and ((ord(password[6])) == ord(password[7])):
        print("correct")
    else:
        print("incorrect")
  else:
    print("incorrect")

if __name__ == "__main__":
  main()
