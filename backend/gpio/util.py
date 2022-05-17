from gpiozero import LED


def isPi():
    try:
        testpin = LED(17)  
    except Exception as e:
        print("No pi detected..") 
        return False
    return True