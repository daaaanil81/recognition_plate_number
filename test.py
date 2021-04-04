import cv2
import numpy as np
import pytesseract

image_file = "outimage.png"
img = cv2.imread(image_file)
gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
# Будет выведен весь текст с картинки
config1 = r'-c tessedit_char_whitelist=0123456789IAKBCEHMOPTX --oem 3 --psm 6'
config2 = r'-l eng --oem 3 --psm 6'

# print(pytesseract.image_to_string(img, config=config))

# Делаем нечто более крутое!!!

print(pytesseract.image_to_string(img, config=config1))
print(pytesseract.image_to_string(img, config=config2))
