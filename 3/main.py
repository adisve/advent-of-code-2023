import functools
import operator


pattern = "[!@#$%^&*()~_+{}|:\"/<>?]"


def is_adjacent_to_symbol(symbol_matrix, start_x, end_x, y) -> tuple:
    rows = len(symbol_matrix)
    cols = len(symbol_matrix[0])

    for dx in range(start_x, end_x + 1):
        for dy in [-1, 1]:
            if 0 <= y + dy < rows and is_symbol(symbol_matrix[y + dy][dx]):
                return (True, (y + dy, dx))

    for dy in range(-1, 2):
        if 0 <= y + dy < rows:
            if start_x > 0 and is_symbol(symbol_matrix[y + dy][start_x - 1]):
                return (True, (y + dy, start_x - 1))
            if end_x < cols - 1 and is_symbol(symbol_matrix[y + dy][end_x + 1]):
                return (True, (y + dy, end_x + 1))

    return (False, None)


def is_symbol(char) -> bool:
    return not char.isnumeric() and not char == '.'


def calculate_gear_ratios(number_indices, schematic):
    adjacencies = {}
    for start_x, end_x, y in number_indices:
        adjacent, symbol_pos = is_adjacent_to_symbol(schematic, start_x,  end_x, y)
        if adjacent and schematic[symbol_pos[0]][symbol_pos[1]] == '*':
            if not adjacencies.get(symbol_pos):
                adjacencies[symbol_pos] = []
            adjacencies[symbol_pos].append(int(''.join(schematic[y][start_x:end_x+1])))
    
    return sum([functools.reduce(operator.mul, values) if len(values) > 1 else 0 for values in adjacencies.values()])


def calculate_sum_parts(number_indices, schematic):
    return sum([int(''.join(schematic[y][start_x:end_x+1])) if is_adjacent_to_symbol(schematic, start_x, end_x, y)[0] else 0 for start_x, end_x, y in number_indices])


def get_num_indices(schematic):
    number_indices = []
    for y in range(len(schematic)):
        start_index = None
        for x in range(len(schematic[y])):
            if schematic[y][x].isnumeric():
                if start_index is None:
                    start_index = x
                if x == len(schematic[y]) - 1 or not schematic[y][x + 1].isnumeric():
                    end_index = x
                    number_indices.append((start_index, end_index, y))
                    start_index = None
            else:
                start_index = None
    return number_indices


def main():
    with open('input.txt', 'r') as infile:
        
        schematic = [list(line) for line in [line.strip() for line in infile.readlines()]]

        number_indices = get_num_indices(schematic)

        sum_parts = calculate_sum_parts(number_indices, schematic)

        gear_ratios = calculate_gear_ratios(number_indices, schematic)

        print(f'Day 3 - Part 1: {sum_parts}')
        print(f'Day 3 - Part 2: {gear_ratios}')


if __name__ == '__main__':
    main()