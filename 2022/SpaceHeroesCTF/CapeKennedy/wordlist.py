#!/usr/bin/python3
import string, random
 
#============ Word List Gen ============
#minimum=input('Please enter the minimum length of any give word to be generated: ')
#maximum=input('Please enter the maximum length of any give word to be generated: ')
#wmaximum=input('Please enter the max number of words to be generate in the dictionary: ')


alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYX0123456789'

word=''
builder=0

FILE = open("wordlistshort.txt","w")


#loop once for the trailers
for trail in range(0,len(alphabet)-1):

  #loop once for the mid
  for mid in range(0,len(alphabet)-1):
  
    #loop once for the third and sixth
    for third in range(0,len(alphabet)-1):
    
      #loop once for the second
      for second in range(0,len(alphabet)-1):
       
        #loop once for the third
        for first in range(0,len(alphabet)-1):
          #first and second char
          word+=alphabet[first]
        
          #second char
          word+=alphabet[second]

          #third
          word+=alphabet[third]

          #fourth and fifth
          word+=alphabet[mid]+alphabet[mid]

          #sixth
          word+=alphabet[third]

          #seventh and eigth
          word+=alphabet[trail]+alphabet[trail]
      
          #reduce the list to those that total 713 in ASCII
          for c in word:
            builder += ord(c)
            
          if builder == 713:
            FILE.write(word+'\n')
            print("Found a match:", word+'\n')
    
          #reset
          word=''
          builder=0

FILE.close()

print('DONE!')



#============    WLG END    ============
