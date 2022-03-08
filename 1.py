from ctypes import cdll
from ctypes import cdll, c_char_p
lib = cdll.LoadLibrary('./libadd.so')
print("Loaded go generated SO library")
parsecode = lib.parsecode
parsecode.argtype = c_char_p
parsecode.restype = c_char_p
result = lib.parsecode("我是".encode("utf-8"))

print('python接受的数据',result.decode('utf-8'))




