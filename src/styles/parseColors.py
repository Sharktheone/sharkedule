#This was used to parse the colors from the mantine website (copy paste) and convert them to scss variables

file = open("colors.txt", "r")
scss = open("colors.scss", "w")

colors = file.readlines()

print(file.read())

print("colors: ")
print(colors)
print("\n\n")

for i in range(0, len(colors), 4):
    name = colors[i].strip()
    color = colors[i+2].strip()
    name = name.replace(" ", "-")
    scss.write("$" + name + "00: " + color + ";\n")