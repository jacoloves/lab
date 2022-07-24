height_cm = float(input('身長(cm)は?'))
weight = float(input('体重(kg)は?'))

height = height_cm / 100.0
bmi = weight / (height ** 2)
print("BMI={:.1f}".format(bmi))

if bmi < 18.5:
    print("low weight")
elif bmi < 25:
    print("weight")
elif bmi < 30:
    print("high weight lv.1")
elif bmi < 35:
    print("high weight lv.2")
elif bmi < 40:
    print("high weight lv.3")
else:
    print("high weight lv.4")
