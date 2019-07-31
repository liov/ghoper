l1 = []
l2 = []
l3 = []

for i in range(1,100,3):
    l1.append(i%15)
    l2.append(i&15)
    l3.append(i%31)

print(l1)
print(l2)
print(l3)