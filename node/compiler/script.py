import sys
import os
import importlib.util

def load_source(module_name, file_path):
    spec = importlib.util.spec_from_file_location(module_name, file_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module

def execFn(fileName, funcName, args):
    filePath = os.path.join("../src", fileName )
    module = load_source(fileName, filePath)
    func = getattr(module, funcName)
    result = func(*args)
    return result

fileName = sys.argv[1]
funcName = sys.argv[2]

i = 3
args = []

while (len(sys.argv) > i):
    args.append(sys.argv[i])
    i = i + 1

result = execFn(fileName, funcName, args)
print(result)