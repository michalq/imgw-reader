PLOTS = {}

def register(name):
    def wrapper(func):
        PLOTS[name] = func
        return func
    return wrapper
