height_cm = float(input('cm?'))
weight = float(input('kg?'))

height = height_cm / 100.0
bmi = weight / (height ** 2)

bmi_list = [
        {"min": 0, "max":18.5, "label":"low weight"},
        {"min": 18.5, "max":25, "label":"normal weight"},
        {"min": 25, "max":30, "label":"heavy weight lv1"},
        {"min": 30, "max":35, "label":"heavy weight lv2"},
        {"min": 35, "max":40, "label":"heavy weight lv3"},
        {"min": 40, "max":99, "label":"heavy weight lv99"}
        ]

result = "unknown"
for range in bmi_list:
    if range["min"] <= bmi < range["max"]:
        result = range["label"]

print("BMI={:.1f},judge={}".format(bmi, result))
