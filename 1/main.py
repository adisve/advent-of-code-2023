codebook = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9'
}

def find_numeric_part(string):
    for i in range(len(string)):
        if string[i].isnumeric():
            return string[i]

        for j in range(3, 6):
            if i + j <= len(string):
                word = string[i:i+j]
                if word in codebook:
                    return codebook[word]
    return ''

with open('input.txt', 'r') as infile:
    calibrations = [line.strip() for line in infile.readlines()]
    total_sum = 0

    for calibration in calibrations:
        forward_number = find_numeric_part(calibration)
        reverse_number = find_numeric_part(calibration[::-1])
        total_sum += int(forward_number + reverse_number)

    print(total_sum)