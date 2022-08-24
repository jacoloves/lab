price = 3950

for i500 in range(0, 11):
    for i100 in range(0, 4):
        for i50 in range(0, 11):
            total = i50 * 50 + i100 * 100 + i500 * 500
            if price == total:
                print("500円x{}+100円x{}+50円x{}={}".format(i500,i100,i50,total))
