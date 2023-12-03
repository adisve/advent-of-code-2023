from operator import mul
from functools import reduce

case_limits = {
    'red': 12, 
    'green': 13, 
    'blue': 14
}

def read_games_from_file(path) -> list:
    with open(path, 'r') as infile:
        return [line.strip() for line in infile.readlines()]

def parse_game(game_str):
    turns = [turn.strip() for turn in game_str.split(':')[-1].split(';')]
    return [process_turn(turn) for turn in turns]

def process_turn(turn_str):
    counts = {'red': 0, 'green': 0, 'blue': 0}
    for case in turn_str.split(','):
        amount, color = case.strip().split()
        counts[color] = max(counts[color], int(amount))
    return counts

def is_game_valid(turn_counts):
    for turn in turn_counts:
        for color, amount in turn.items():
            if amount > case_limits[color]:
                return False
    return True

def get_least_amount(turn_counts):
    mins = { 'red': 0, 'green': 0, 'blue': 0 }
    for turn in turn_counts:
        for color, amount in turn.items():
            if amount > mins[color]:
                mins[color] = amount
    return reduce(mul, mins.values(), 1)

def main():
    valid_games_sum = 0
    games = read_games_from_file('input.txt')
    for game_id, game_str in enumerate(games, start=1):
        turn_counts = parse_game(game_str)
        if is_game_valid(turn_counts):
            valid_games_sum += game_id

    print(f'Sum of valid games: {valid_games_sum}')

    print(f'Sum of the power of sets of least amount of cubes: {sum([get_least_amount(parse_game(game_str)) for game_str in games])}')


if __name__ == '__main__':
    main()
